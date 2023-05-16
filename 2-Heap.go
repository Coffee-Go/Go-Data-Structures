package Go_Data_Structures

// MaxHeap struct
type MaxHeap struct {
	array []int
}

// NewMaxHeap create a new max heap
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		array: make([]int, 0, 10),
	}
}

// Insert item to heap
func (h *MaxHeap) Insert(data int) {
	h.array = append(h.array, data)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract item from heap
func (h *MaxHeap) Extract() int {
	lenHeap := len(h.array) - 1

	if lenHeap < 0 {
		return -1
	}

	data := h.array[0]
	h.array[0] = h.array[lenHeap]
	h.array = h.array[:lenHeap]
	h.maxHeapifyDown(0)

	return data
}

// maxHeapifyUp
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[h.parent(index)] < h.array[index] {
		h.swap(h.parent(index), index)
		index = h.parent(index)
	}
}

// maxHeapifyDown
func (h *MaxHeap) maxHeapifyDown(index int) {
	lenHeap := len(h.array) - 1

	for index < lenHeap {
		childToCompare := h.childToCompare(index)

		if h.array[childToCompare] < h.array[index] {
			break
		}

		h.swap(childToCompare, index)
		index = childToCompare
	}
}

// swap
func (h *MaxHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

// parent
func (h *MaxHeap) parent(index int) int {
	return (index - 1) / 2
}

// left
func (h *MaxHeap) left(index int) int {
	return 2*index + 1
}

// right
func (h *MaxHeap) right(index int) int {
	return 2*index + 2
}

// childToCompare
func (h *MaxHeap) childToCompare(index int) int {
	lenHeap := len(h.array) - 1
	childToCompare := 0
	l, r := h.left(index), h.right(index)

	if l >= lenHeap || r > lenHeap {
		childToCompare = lenHeap
	} else if h.array[l] > h.array[r] {
		childToCompare = l
	} else {
		childToCompare = r
	}

	return childToCompare
}
