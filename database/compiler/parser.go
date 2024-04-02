package compiler

import (
	"database/compiler/lib"
	"fmt"
	"strconv"
	"time"
)

type LRParser struct {
	input       []Token
	handlerName []string
	stack       lib.Stack
	Lrt         LRTable
	handlers    SetOfFunc
}

func (c *Compiler) NewParser(handler SetOfFunc) *LRParser {
	lrp := &LRParser{
		stack:       lib.Stack{Item: []lib.Stackable{&SState{state: 0}}},
		Lrt:         *c.LRTable,
		handlerName: c.LRTable.HandlerName,
		handlers:    handler,
	}
	return lrp
}

func (p *LRParser) Parse(input []Token) error {
	p.input = input
	start := time.Now()
	for {
		if len(p.input) == 0 {
			return fmt.Errorf("empty")
		}

		var state int
		if st, ok := p.stack.Top().(*SState); ok {
			state = st.Value().(int)
		}

		token := p.input[0].Type
		action, ok := p.Lrt.States[state].LrAction[token]
		for _, i := range p.stack.Item {
			fmt.Print(i.Value(), " ")
		}
		fmt.Println()
		if !ok {
			return fmt.Errorf("cannot find action")
		}
		switch action[0] {
		case 'a':
			loadDuration := time.Since(start)
			fmt.Println("Execution Time for Parse & Execute:", loadDuration)
			return nil
		case 's':
			if err := p.shift(action); err != nil {
				return err
			}
		case 'r':
			ruleIdx, err := strconv.Atoi(action[1:])
			if err != nil {
				return fmt.Errorf("cannot convert string to int")
			}
			terminalValue, isEpsilon, err := p.reduce(ruleIdx)
			if err != nil {
				return err
			}
			if !isEpsilon {
				if terminalValue == nil {
					terminalValue = [][]string{{"test"}}
				}
				rule := p.Lrt.Grammar.GetRuleByIndex(ruleIdx)
				nonTerminal := rule.Pattern[0]
				gotoState := p.Lrt.States[p.stack.Top().Value().(int)].LrGoto[nonTerminal]
				p.goTo(terminalValue, gotoState)
			} else {
				for _, next := range p.Lrt.States[p.stack.Top().Value().(int)].LrGoto {
					terminalValue = [][]string{{"nil"}}
					p.goTo(terminalValue, next)
				}
			}
		default:
			return fmt.Errorf("format invalid")
		}
	}
}

func (p *LRParser) shift(action string) error {
	nextState, err := strconv.Atoi(action[1:])
	if err != nil {
		return fmt.Errorf("cannot convert string to int")
	}
	p.stack.Push(&Terminal{id: p.input[0].Value})
	p.stack.Push(&SState{state: nextState})
	p.input = p.input[1:]
	return nil
}

func (p *LRParser) reduce(ruleIdx int) ([][]string, bool, error) {
	rule := p.Lrt.Grammar.GetRuleByIndex(ruleIdx)
	if !(len(rule.Development) == 1 && rule.Development[0] == EPSILON) {
		var values [][]string
		for i := 0; i < len(rule.Development)*2; i++ {
			if t, ok := p.stack.Top().(*Terminal); ok {
				values = append(values, []string{t.Value().(string)})
			}
			if t, ok := p.stack.Top().(*Nontermial); ok {
				values = append(values, t.Value().([][]string)...)
			}
			p.stack.Pop()
		}
		t, err := p.handlers.executeHandler(p.handlerName[ruleIdx], values)
		if err != nil {
			return nil, true, err
		}
		return t, false, nil
	}
	return nil, true, nil
}

func (p *LRParser) goTo(terminalValue [][]string, nextState int) {
	p.stack.Push(&Nontermial{list: terminalValue})
	p.stack.Push(&SState{state: nextState})
}
