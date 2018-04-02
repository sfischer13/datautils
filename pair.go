package datautils

import "sort"

// Pair TODO
type Pair struct {
	Key   string
	Value int64
}

// Pairs TODO
type Pairs []Pair

// ReverseKeys TODO
func (ps Pairs) ReverseKeys() {
	sort.SliceStable(ps, func(i, j int) bool {
		return ps[i].Key > ps[j].Key
	})

}

// ReverseValues TODO
func (ps Pairs) ReverseValues() {
	sort.SliceStable(ps, func(i, j int) bool {
		return ps[i].Value > ps[j].Value
	})
}

// SortKeys TODO
func (ps Pairs) SortKeys() {
	sort.SliceStable(ps, func(i, j int) bool {
		return ps[i].Key < ps[j].Key
	})
}

// SortValues TODO
func (ps Pairs) SortValues() {
	sort.SliceStable(ps, func(i, j int) bool {
		return ps[i].Value < ps[j].Value
	})
}
