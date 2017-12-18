package datautils

type List struct {
	Elements []string
}

func NewList() List {
	list := List{}
	return list
}

func (l *List) Add(element string) {
	l.Elements = append(l.Elements, element)
}

func (l *List) Clear() {
	l.Elements = []string{}
}

func (l *List) IsEmpty() bool {
	return len(l.Elements) == 0
}
