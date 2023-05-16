package Go_Data_Structures

type Data struct {
	Key   int
	Value string
}

type Node struct {
	Data Data
	Next *Node
}

type LinkedList struct {
	Head *Node
	Len  int
}

func NewData(key int, value string) Data {
	return Data{
		Key:   key,
		Value: value,
	}
}

func NewNode(data Data) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Len:  0,
	}
}

func (l *LinkedList) IsEmpty() bool {
	return l.Len == 0
}

func (l *LinkedList) Add(data Data) {
	node := NewNode(data)
	node.Next = l.Head
	l.Head = node
	l.Len++
}

func (l *LinkedList) Find(data Data) bool {
	node := l.Head

	for node != nil {
		if node.Data.Key == data.Key {
			return true
		}

		node = node.Next
	}

	return false
}

func (l *LinkedList) Delete(data Data) {
	var pNode *Node = nil
	node := l.Head

	for node != nil {
		if node.Data.Key == data.Key {
			if pNode == nil {
				l.Head = node.Next
			} else {
				pNode.Next = node.Next
			}
			l.Len--
		}

		pNode = node
		node = node.Next
	}
}

func (l *LinkedList) Next() *Data {
	node := l.Head

	if l.IsEmpty() {
		return nil
	}

	l.Head = node.Next
	l.Len--

	return &node.Data
}
