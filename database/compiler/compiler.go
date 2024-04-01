package compiler

import "strings"

type Compiler struct {
	LRTable  *LRTable
}

func InitCompiler(grammer string) *Compiler {
	comp := Compiler{
		LRTable: initLRtable(grammer),
	}
	return &comp
}

func initLRtable(text string) *LRTable {
	grammer := ""
	var handler []string
	lines := strings.Split(text, "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
		if len(lines[i]) == 0 {
			continue
		}
		if strings.Contains(lines[i], "##") {
			parts := strings.Split(lines[i], "##")
			grammer += parts[0] + "\n"
			handler = append(handler, parts[1])
		} else {
			grammer += lines[i] + "\n"
			handler = append(handler, "empty")
		}
	}
	g := NewGrammar(grammer)
	lrClosureTable := newLRClosureTable(*g)
	lrt := BuildState(lrClosureTable)
	lrt.HandlerName = handler
	return lrt
}
