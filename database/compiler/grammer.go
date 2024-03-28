package compiler

import (
	"database/compiler/lib"
	"strings"
)

type Grammar struct {
	Alphabet     []string
	Nonterminals []string
	Terminals    []string
	Rules        []Rule
	Firsts       map[string][]string
	Follows      map[string][]string
	Axiom        string
}

func NewGrammar(text string) *Grammar {
	g := &Grammar{
		Alphabet:     []string{},
		Nonterminals: []string{},
		Terminals:    []string{},
		Rules:        []Rule{},
		Firsts:       make(map[string][]string),
		Follows:      make(map[string][]string),
	}

	g.initializeRulesAndAlphabetAndNonterminals(text)
	g.initializeAlphabetAndTerminals()
	g.initializeFirsts()
	g.initializeFollows()
	return g
}

func (g *Grammar) GetRulesForNonterminal(nonterminal string) []*Rule {
	var result []*Rule
	for _, rule := range g.Rules {
		// fmt.Println("Compare Nonterminal", rule.Nonterminal, nonterminal, rule.Nonterminal == nonterminal)
		if nonterminal == rule.Nonterminal {
			result = append(result, &rule)
		}
	}
	return result
}

func (g *Grammar) GetSequenceFirsts(sequence []string) []string {
	var result []string
	epsilonInSymbolFirsts := true
	for _, symbol := range sequence {
		epsilonInSymbolFirsts = false
		if lib.IsElement(symbol, g.Terminals) {
			lib.AddUnique(symbol, &result)
			break
		}
		for _, first := range g.Firsts[symbol] {
			epsilonInSymbolFirsts = epsilonInSymbolFirsts || first == EPSILON
			lib.AddUnique(first, &result)
		}
		epsilonInSymbolFirsts = epsilonInSymbolFirsts || len(g.Firsts[symbol]) == 0
		if !epsilonInSymbolFirsts {
			break
		}
	}
	if epsilonInSymbolFirsts {
		lib.AddUnique(EPSILON, &result)
	}
	return result
}

func (g *Grammar) GetRuleByIndex(idx int) *Rule {
	if 0 < idx && idx < len(g.Rules) {
		return &g.Rules[idx]
	}
	return nil
}

func (g *Grammar) initializeRulesAndAlphabetAndNonterminals(text string) {
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			rule := NewRule(g, line)
			g.Rules = append(g.Rules, *rule)

			if g.Axiom == "" {
				g.Axiom = rule.Nonterminal
			}

			lib.AddUnique(rule.Nonterminal, &g.Alphabet)
			lib.AddUnique(rule.Nonterminal, &g.Nonterminals)
		}
	}
}

func (g *Grammar) initializeAlphabetAndTerminals() {
	for _, rule := range g.Rules {
		for _, symbol := range rule.Development {
			if symbol != EPSILON && !lib.IsElement(symbol, g.Nonterminals) {
				lib.AddUnique(symbol, &g.Alphabet)
				lib.AddUnique(symbol, &g.Terminals)
			}
		}
	}
}

func (g *Grammar) initializeFirsts() {
	var notDone bool

	for {
		notDone = false
		for _, rule := range g.Rules {
			if len(rule.Development) == 1 && rule.Development[0] == EPSILON {
				notDone = lib.AddUniqueMap(EPSILON, g.Firsts, rule.Nonterminal)
			} else {
				notDone = g.collectDevelopmentFirsts(rule.Development, rule.Nonterminal) || notDone
			}
		}

		if !notDone {
			break
		}
	}
}

func (g *Grammar) collectDevelopmentFirsts(development []string, key string) bool {
	var result, epsilonInSymbolFirsts bool

	for _, symbol := range development {
		epsilonInSymbolFirsts = false

		if lib.IsElement(symbol, g.Terminals) {
			result = lib.AddUniqueMap(symbol, g.Firsts, key) || result
			break
		}

		for _, first := range g.Firsts[symbol] {
			epsilonInSymbolFirsts = epsilonInSymbolFirsts || first == EPSILON

			result = lib.AddUniqueMap(first, g.Firsts, key) || result
		}

		if !epsilonInSymbolFirsts {
			break
		}
	}

	if epsilonInSymbolFirsts {
		result = lib.AddUniqueMap(EPSILON, g.Firsts, key) || result
	}
	return result
}

func (g *Grammar) initializeFollows() {
	var notDone bool

	for !notDone {
		notDone = false

		for i, rule := range g.Rules {
			if i == 0 {
				// nonterminalFollows := lib.GetOrCreateArray(g.Follows, rule.Nonterminal)
				notDone = lib.AddUniqueMap("$", g.Follows, rule.Nonterminal) || notDone
			}

			for j, symbol := range rule.Development {
				if lib.IsElement(symbol, g.Nonterminals) {
					// symbolFollows := lib.GetOrCreateArray(g.Follows, symbol)
					afterSymbolFirsts := g.GetSequenceFirsts(rule.Development[j+1:])

					for _, first := range afterSymbolFirsts {
						if first == EPSILON {
							nonterminalFollows := g.Follows[rule.Nonterminal]

							for _, ntFollow := range nonterminalFollows {
								notDone = lib.AddUniqueMap(ntFollow, g.Follows, symbol) || notDone
							}
						} else {
							notDone = lib.AddUniqueMap(first, g.Follows, symbol) || notDone
						}
					}
				}
			}
		}
	}
}
