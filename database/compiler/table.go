package compiler

type LRClosureTable struct {
	Grammar *Grammar
	Kernels []*Kernel
}

func newLRClosureTable(grammar Grammar) *LRClosureTable {
	lrClosureTable := &LRClosureTable{
		Grammar: &grammar,
		Kernels: []*Kernel{NewKernel(0, []*Item{NewItem(grammar.Rules[0], 0)}, grammar)},
	}
	lrClosureTable.initialize()
	return lrClosureTable
}

func (lrct *LRClosureTable) initialize() {
	// fmt.Println("================================================================")
	for i := 0; i < len(lrct.Kernels); {
		kernel := lrct.Kernels[i]
		// fmt.Println("////////////////////////////////////////////////////////////")
		// fmt.Println(len(lrct.Kernels))
		// for _, k := range kernel.Closure {
		// 	fmt.Println(k)
		// }
		// fmt.Println("////////////////////////////////////////////////////////////")
		// for _, k := range kernel.Closure {
		// fmt.Println("Kernel Dot Index", k.DotIndex)
		// fmt.Println("Kernel Pattern", k.Rule.Pattern)
		// fmt.Println("Kernel Development", k.Rule.Development)
		// fmt.Println("Kernel LookAheads", k.LookAheads)
		// fmt.Println("-----")
		// }
		// fmt.Println("Kernel Gotos", kernel.Gotos)
		// fmt.Println("Kernel Index", kernel.Index)
		// fmt.Println("--Update Closure")
		lrct.updateClosure(kernel)

		if lrct.addGotos(kernel) {
			i = 0
		} else {
			i++
		}
	}
}

func (lrct *LRClosureTable) updateClosure(kernel *Kernel) {
	// fmt.Println("closure Len", len(kernel.Closure))
	for _, item := range kernel.Closure {
		// fmt.Println("For item in Closure", item)
		// fmt.Println("--imte.NewItemsFromSymbolAfterDot")
		newItemsFromSymbolAfterDot := item.NewItemsFromSymbolAfterDot()
		// fmt.Println("added", len(kernel.Closure))
		// fmt.Println("new", newItemsFromSymbolAfterDot)
		for _, newItem := range newItemsFromSymbolAfterDot {
			newItem.AddUniqueTo(&kernel.Closure)
		}
		// fmt.Println("!!!!updateClosure", len(kernel.Closure))
		// for _, clo := range kernel.Closure {
		// 	fmt.Println(clo)
		// }
	}
}

func (lrct *LRClosureTable) addGotos(kernel *Kernel) bool {
	lookAheadsPropagated := false
	newKernels := make(map[string][]*Item)
	// fmt.Println("AddGotos", len(kernel.Closure))
	for _, item := range kernel.Closure {
		newItem := item.NewItemAfterShift()
		// fmt.Println("GOTOS")
		// fmt.Println("Key11", kernel.Keys)
		if newItem != nil {
			symbolAfterDot := item.Rule.Development[item.DotIndex]
			kernel.addUniqueToKeys(symbolAfterDot)
			newKernels[symbolAfterDot] = append(newKernels[symbolAfterDot], newItem)
		}
		// fmt.Println("After Add Key")
		// fmt.Println("After Key", kernel.Keys)
		// fmt.Println("New Kernel", newKernels)
		// fmt.Println("/GOTOS")
	}

	for _, key := range kernel.Keys {
		newKernel := NewKernel(len(lrct.Kernels), newKernels[key], *lrct.Grammar)
		targetKernelIndex := lrct.indexOfUsingEquals(newKernel, lrct.Kernels)
		// fmt.Println("newKernel", len(lrct.Kernels))
		// fmt.Println("targetKernelIndex", targetKernelIndex)
		if targetKernelIndex < 0 {
			lrct.Kernels = append(lrct.Kernels, newKernel)
			targetKernelIndex = newKernel.Index
		} else {
			for _, newItem := range newKernel.Items {
				lookAheadsPropagated = newItem.AddUniqueTo(&lrct.Kernels[targetKernelIndex].Items)
			}
		}

		kernel.Gotos[key] = targetKernelIndex
	}

	return lookAheadsPropagated
}

func (lrct *LRClosureTable) indexOfUsingEquals(target *Kernel, kernels []*Kernel) int {
	for i, k := range kernels {
		if target.Equals(k) {
			return i
		}
	}
	return -1
}
