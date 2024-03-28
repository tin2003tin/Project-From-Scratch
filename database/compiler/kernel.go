package compiler

type Kernel struct {
	Index   int
	Items   []*Item
	Closure []*Item
	Gotos   map[string]int
	Keys    []string
}

func NewKernel(index int, items []*Item, grammar Grammar) *Kernel {
	closure := make([]*Item, len(items))
	copy(closure, items)

	return &Kernel{
		Index:   index,
		Items:   items,
		Closure: closure,
		Gotos:   make(map[string]int),
		Keys:    []string{},
	}
}

func (k *Kernel) Equals(that *Kernel) bool {
	if k.Items[0].DotIndex == that.Items[0].DotIndex && k.Items[0].Rule.Pattern[0] == that.Items[0].Rule.Pattern[0] && len(k.Items[0].Rule.Development) == len(that.Items[0].Rule.Development) {
		for i := range k.Items[0].Rule.Development {
			if k.Items[0].Rule.Development[i] != that.Items[0].Rule.Development[i] {
				return false
			}
			return true
		}
	}
	return false
}

func (k *Kernel) addUniqueToKeys(key string) {
	for _, k := range k.Keys {
		if k == key {
			return
		}
	}
	k.Keys = append(k.Keys, key)
}
