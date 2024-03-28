package compiler

import (
	"database/compiler/lib"
)

type Item struct {
	Rule       Rule
	DotIndex   int
	LookAheads []string
}

func NewItem(rule Rule, dotIndex int) *Item {
	bli := Item{Rule: rule,
		DotIndex: dotIndex,
	}
	if rule.Index == 0 {
		bli.LookAheads = []string{"$"}
	} else {
		bli.LookAheads = make([]string, 0)
	}
	return &bli
}

func (bli *Item) Equals(other interface{}) bool {
	if otherBLi, ok := other.(*Item); ok {
		if len(otherBLi.Rule.Development) != len(bli.Rule.Development) {
			return false
		}
		for i := range otherBLi.Rule.Development {
			if bli.Rule.Development[i] != otherBLi.Rule.Development[i] {
				return false
			}
		}
		return bli.DotIndex == otherBLi.DotIndex && bli.Rule.Nonterminal == otherBLi.Rule.Nonterminal
	}
	return false
}

func (bli *Item) NewItemsFromSymbolAfterDot() []*Item {
	var result []*Item
	// fmt.Println("item Rule Development", bli.Rule.Development)
	if bli.DotIndex >= len(bli.Rule.Development) {
		return result
	}
	// fmt.Println("item Rule Development at DotInex", bli.Rule.Development[bli.DotIndex])
	// fmt.Println("--GetRulesForNonterminal at DotIndex")
	nonterminalRules := bli.Rule.Grammar.GetRulesForNonterminal(bli.Rule.Development[bli.DotIndex])
	for _, nonterminalRule := range nonterminalRules {
		newNonterminalRules := nonterminalRule.Grammar.GetRulesForNonterminal(nonterminalRule.Development[0])
		for _, newRule := range newNonterminalRules {
			for _, existingItem := range result {
				if existingItem.DotIndex != 0 && existingItem.Rule.Pattern[0] != newRule.Pattern[0] {
					if len(existingItem.Rule.Development) == len(newRule.Development) {
						mismatch := false
						for i := range existingItem.Rule.Development {
							if existingItem.Rule.Development[i] != newRule.Development[i] {
								mismatch = true
								break
							}
						}
						if !mismatch {
							nonterminalRules = append(nonterminalRules, newRule)
							break
						}
					} else {
						nonterminalRules = append(nonterminalRules, newRule)
						break
					}
				}
			}
		}
	}
	// fmt.Println("All Rule that is Nontermial", nonterminalRules)
	// fmt.Println("--AddUniqueUsingEquals")
	for _, nonterminalRules := range nonterminalRules {
		bli.AddUniqueUsingEquals(*NewItem(*nonterminalRules, 0), &result)
	}
	// fmt.Println("--AddUniqueUsingEquals Completed result is", result)
	newLookAheads := []string{}
	epsilonPresent := false
	firstsAfterSymbolAfterDot := bli.Rule.Grammar.GetSequenceFirsts(bli.Rule.Development[bli.DotIndex+1:])
	// fmt.Println("firstsAfterSymbolAfterDot", firstsAfterSymbolAfterDot)
	for _, first := range firstsAfterSymbolAfterDot {
		if first == EPSILON {
			epsilonPresent = true
		} else {
			lib.AddUnique(first, &newLookAheads)
		}
	}
	// fmt.Println("newLookAheads", newLookAheads)
	if epsilonPresent {
		for _, lh := range bli.LookAheads {
			lib.AddUnique(lh, &newLookAheads)
		}
	}

	for _, res := range result {
		res.LookAheads = append([]string{}, newLookAheads...)
	}
	// fmt.Println("Final NewItemsFromSymbolAfterDot")
	// for _, res := range result {
	// 	fmt.Println(res)
	// }
	return result
}

func (bli *Item) NewItemAfterShift() *Item {
	var result *Item // Declare result as a pointer to Item

	if bli.DotIndex < len(bli.Rule.Development) && bli.Rule.Development[bli.DotIndex] != EPSILON {
		result = &Item{Rule: bli.Rule, DotIndex: bli.DotIndex + 1} // Initialize result if conditions are met
	}

	if result != nil {
		result.LookAheads = append([]string{}, bli.LookAheads...) // Copy bli.LookAheads to result.LookAheads
	}

	return result // Return the result
}

func (bli *Item) AddUniqueUsingEquals(element Item, array *[]*Item) bool {
	// fmt.Println("AddUniqueUsingEquals input ", element, array)
	// fmt.Println("Check that element is array")
	if !bli.isElementUsingEquals(element, *array) {
		*array = append(*array, &element)
		return true
	}
	return false
}

func (bli *Item) AddUniqueTo(items *[]*Item) bool {
	// fmt.Printf("Address of bli: %p\n", bli)
	// fmt.Printf("Address of items[0]: %p\n", (*items)[0])
	// fmt.Println(bli.Equals((*items)[0]))
	// fmt.Println("compare", bli, (*items)[0], bli.Equals((*items)[0]))
	for _, item := range *items {
		if bli.Equals(item) {
			for _, lh := range bli.LookAheads {
				if lib.AddUnique(lh, &item.LookAheads) {
					return true // Return true if any lookahead was added to item.LookAheads
				}
			}
			return false // Return false if no lookahead was added to item.LookAheads
		}
	}

	(*items) = append((*items), bli)
	return true
}

func (bli *Item) isElementUsingEquals(element Item, array []*Item) bool {
	// fmt.Println("Start len of array", len(array))
	for _, arr := range array {
		if element.Equals(arr) {
			return true
		}
	}
	return false
}
