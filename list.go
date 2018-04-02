package datautils

// List TODO
type List struct {
	Elements []string
}

// NewList TODO
func NewList() List {
	list := List{}
	return list
}

// Add TODO
func (l *List) Add(element string) {
	l.Elements = append(l.Elements, element)
}

// Clear TODO
func (l *List) Clear() {
	l.Elements = []string{}
}

// IsEmpty TODO
func (l *List) IsEmpty() bool {
	return len(l.Elements) == 0
}
