package compiler

import (
	"database/compiler/lib"
	"fmt"
	"strings"
)

type Rule struct {
	Grammar     *Grammar
	Index       int
	Nonterminal string
	Pattern     []string
	Development []string
}

func NewRule(grammar *Grammar, text string) *Rule {
	r := &Rule{
		Grammar: grammar,
		Index:   len(grammar.Rules),
	}

	split := strings.Split(text, "->")
	r.Nonterminal = strings.TrimSpace(split[0])
	r.Pattern = lib.TrimElements(strings.Split(r.Nonterminal, " "))
	r.Development = lib.TrimElements(strings.Split(strings.TrimSpace(split[1]), " "))

	return r
}

func (r *Rule) String() string {
	return fmt.Sprintf("%s -> %s", r.Nonterminal, strings.Join(r.Development, " "))
}

func (r *Rule) Equals(that *Rule) bool {
	if r.Nonterminal != that.Nonterminal {
		return false
	}

	if len(r.Development) != len(that.Development) {
		return false
	}

	for i := range r.Development {
		if r.Development[i] != that.Development[i] {
			return false
		}
	}

	return true
}
