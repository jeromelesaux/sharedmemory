package model

import "encoding/json"

type Element struct {
	Value    any      `json:"value"`
	Next     *Element `json:"next"`
	Previous *Element `json:"-"`
}

type LinkedList struct {
	Root Element `json:"root"`
	Len  int     `json:"len"`
}

func NewLinkedList() LinkedList {
	l := LinkedList{}
	l.Root.Next = &l.Root
	l.Root.Previous = &l.Root
	l.Len = 0
	return l
}

func (l *LinkedList) Add(value any) {
	var e *Element
	e = &l.Root
	for x := 0; x < l.Len; x++ {
		e = e.Next
	}
	e.Next = &Element{Value: value, Next: &l.Root, Previous: e}
	l.Len++
}

type jsonLinkedList struct {
	Elems []any `json:"element"`
}

func (l LinkedList) MarshalJSON() ([]byte, error) {
	if l.Len == 0 {
		return []byte{}, nil
	}

	j := jsonLinkedList{Elems: make([]any, l.Len)}
	e := l.Root.Next
	for i := 0; i < l.Len; i++ {
		j.Elems[i] = e.Value
		e = e.Next
	}
	return json.Marshal(&j)
}

func (l *LinkedList) UnmarshalJSON(data []byte) error {
	var j jsonLinkedList
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	l.Len = len(j.Elems)
	e := &l.Root
	for _, v := range j.Elems {
		e.Next = &Element{Value: v, Previous: e, Next: &l.Root}
		e = e.Next
	}
	return nil
}
