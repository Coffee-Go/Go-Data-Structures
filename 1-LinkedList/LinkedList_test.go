package LinkedList

import (
	"testing"
)

func TestLinkedList_Add(t *testing.T) {
	tests := []struct {
		name string
		add  []Data
		next []string
	}{
		{
			name: "Case 1",
			add: func() []Data {
				data := make([]Data, 4)

				data[0] = NewData(1, "A")
				data[1] = NewData(2, "B")
				data[2] = NewData(3, "C")
				data[3] = NewData(4, "D")

				return data
			}(),
			next: []string{"D", "C", "A"},
		},
	}
	for _, tt := range tests {
		l := NewLinkedList()

		for _, v := range tt.add {
			l.Add(v)
		}

		if got := l.Head.Data.Key; got != 4 {
			t.Errorf("LinkedList Add() Wont: %v, Got: %v", 4, got)
		}

		if got := l.Find(tt.add[1]); got != true {
			t.Errorf("LinkedList Find() Wont: %v, Got: %v", true, got)
		}

		l.Delete(tt.add[1])

		if got := l.Find(tt.add[1]); got != false {
			t.Errorf("LinkedList Find() Wont: %v, Got: %v", false, got)
		}

		for _, v := range tt.next {
			if got := l.Next(); got.Value != v {
				t.Errorf("LinkedList Next() Wont: %v, Got: %v", v, got)
			}
		}
	}
}
