package datautils

import "sort"

type Pair struct {
	Key   string
	Value int64
}

type Pairs []Pair

func (ps Pairs) ReverseKeys() {
	sort.SliceStable(ps, func(i, j int) bool {
		return ps[i].Key > ps[j].Key
	})

}

func (ps Pairs) ReverseValues() {
	sort.SliceStable(ps, func(i, j int) bool {
		return ps[i].Value > ps[j].Value
	})
}

func (ps Pairs) SortKeys() {
	sort.SliceStable(ps, func(i, j int) bool {
		return ps[i].Key < ps[j].Key
	})
}

func (ps Pairs) SortValues() {
	sort.SliceStable(ps, func(i, j int) bool {
		return ps[i].Value < ps[j].Value
	})
}
