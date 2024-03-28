package compiler

import (
	"database/compiler/lib"
	"strconv"
)

type LRTable struct {
	Grammar     *Grammar
	States      []*State
	HandlerName []string
}

type State struct {
	Index    int
	LrAction map[string]string
	LrGoto   map[string]int
}

func BuildState(closureTable *LRClosureTable) *LRTable {
	lrTable := &LRTable{
		Grammar: closureTable.Grammar,
		States:  []*State{},
	}

	for _, kernel := range closureTable.Kernels {
		state := &State{
			Index:    len(lrTable.States),
			LrAction: make(map[string]string),
			LrGoto:   make(map[string]int),
		}
		lrTable.States = append(lrTable.States, state)

		for _, key := range kernel.Keys {
			nextStateIndex := kernel.Gotos[key]
			if lib.IsElement(key, closureTable.Grammar.Terminals) {
				text := "s" + strconv.Itoa(nextStateIndex)
				state.LrAction[key] = text
			} else {
				state.LrGoto[key] = nextStateIndex
			}
		}

		for _, item := range kernel.Closure {
			if item.DotIndex == len(item.Rule.Development) || item.Rule.Development[0] == EPSILON {
				for _, lookAhead := range item.LookAheads {
					text := "r" + strconv.Itoa(item.Rule.Index)
					if text == "r0" {
						text = "accept"
					}
					state.LrAction[lookAhead] = text
				}
			}
		}
	}

	return lrTable
}
