package compiler

import (
	"database/compiler/lib"
	"fmt"
	"strconv"
)

type LRParser struct {
	input   []Token
	handler []string
	stack   lib.Stack
	Lrt     LRTable
}

func (c *Compiler) NewParser(input []Token, lrt *LRTable) *LRParser {
	if len(input) == 0 {
		panic("Input tokens are empty")
	}
	lrp := &LRParser{
		input:   input,
		stack:   lib.Stack{Item: []lib.Stackable{&SState{state: 0}}},
		Lrt:     *lrt,
		handler: lrt.HandlerName,
	}
	c.LRParser = lrp
	return lrp
}

func (c *Compiler) Parse(handler SetOfFunc) error {
	for {
		p := c.LRParser
		if len(p.input) == 0 {
			return fmt.Errorf("empty")
		}

		var state int
		if st, ok := p.stack.Top().(*SState); ok {
			state = st.Value().(int)
		}

		token := p.input[0].Type
		action, ok := p.Lrt.States[state].LrAction[token]
		if !ok {
			return fmt.Errorf("cannot find action")
		}
		switch action[0] {
		case 'a':
			return nil
		case 's':
			if err := c.shift(action); err != nil {
				return err
			}
		case 'r':
			ruleIdx, err := strconv.Atoi(action[1:])
			if err != nil {
				return fmt.Errorf("cannot convert string to int")
			}
			terminalValue, isEpsilon, err := c.reduce(ruleIdx)
			if err != nil {
				return err
			}
			if !isEpsilon {
				if terminalValue == nil {
					terminalValue = [][]string{{"test"}}
				}
				rule := c.LRParser.Lrt.Grammar.GetRuleByIndex(ruleIdx)
				nonTerminal := rule.Pattern[0]
				gotoState := p.Lrt.States[p.stack.Top().Value().(int)].LrGoto[nonTerminal]
				p.goTo(terminalValue, gotoState)
			} else {
				for _, next := range p.Lrt.States[p.stack.Top().Value().(int)].LrGoto {
					p.goTo(nil, next)
				}
			}
		default:
			return fmt.Errorf("format invalid")
		}
	}
}

func (c *Compiler) shift(action string) error {
	nextState, err := strconv.Atoi(action[1:])
	if err != nil {
		return fmt.Errorf("cannot convert string to int")
	}
	c.LRParser.stack.Push(&Terminal{id: c.LRParser.input[0].Value})
	c.LRParser.stack.Push(&SState{state: nextState})
	c.LRParser.input = c.LRParser.input[1:]
	return nil
}

func (c *Compiler) reduce(ruleIdx int) ([][]string, bool, error) {
	rule := c.LRParser.Lrt.Grammar.GetRuleByIndex(ruleIdx)
	if !(len(rule.Development) == 1 && rule.Development[0] == EPSILON) {
		var values [][]string
		for i := 0; i < len(rule.Development)*2; i++ {
			if t, ok := c.LRParser.stack.Top().(*Terminal); ok {
				values = append(values, []string{t.Value().(string)})
			}
			if t, ok := c.LRParser.stack.Top().(*Nontermial); ok {
				values = append(values, t.Value().([]string))
			}
			c.LRParser.stack.Pop()
		}
		t, err := c.Handlers.executeHandler(c.LRParser.handler[ruleIdx], values)
		if err != nil {
			return nil, true, err
		}
		return t, true, nil
	}
	return nil, false, nil
}

func (p *LRParser) goTo(terminalValue [][]string, nextState int) {
	p.stack.Push(&Nontermial{list: terminalValue})
	p.stack.Push(&SState{state: nextState})
}
