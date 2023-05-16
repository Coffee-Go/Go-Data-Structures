package Go_Data_Structures

import (
	"testing"
)

func TestNewMaxHeap(t *testing.T) {
	tests := []struct {
		name    string
		insert  []int
		extract []int
	}{
		{
			name:    "Case 1",
			insert:  []int{9, 1, 5, 7, 14, 34, 20, 8, 48, 50, 16},
			extract: []int{50, 48, 34, 20, 16, 14, 9, 8, 7, 5, 1},
		},
		{
			name:    "Case 2",
			insert:  []int{9, 1, 5, 7},
			extract: []int{9, 7, 5, 1, -1, -1, -1},
		},
		{
			name:    "Case 3",
			insert:  []int{9, 1, 5, 7, 14, 34, 20, 8, 48, 50, 16, 9, 1, 5, 7, 14, 34, 20, 8, 48, 50, 16, 9, 1, 5, 7, 14, 34, 20, 8, 48, 50, 16},
			extract: []int{50, 50, 50, 48, 48, 48, 34, 34, 34, 20, 20, 20, 16, 16, 16, 14, 14, 9, 9, 9, 8, 8, 14, 8, 7, 7, 7, 5, 5, 5, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := NewMaxHeap()

			for _, i := range tt.insert {
				heap.Insert(i)
			}

			for _, i := range tt.extract {
				if got := heap.Extract(); got != i {
					t.Errorf("Max Heap Extract() wont: %v Got: %v", i, got)
				}
			}
		})
	}
}
