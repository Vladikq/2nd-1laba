package main

import "fmt"

// ListNode - узел списка
type ListNode struct {
	Value string
	Next  *ListNode
	Prev  *ListNode // nil для односвязного списка
}

// List - интерфейс для списка
type List interface {
	AddFront(value string)
	AddBack(value string)
	Remove(value string) bool
	Contains(value string) bool
	Size() int
	Print()
}

// ================== Singly Linked List ==================

type SinglyList struct {
	head *ListNode
	size int
}

func NewSinglyList() *SinglyList {
	return &SinglyList{}
}

func (sl *SinglyList) AddFront(value string) {
	sl.head = &ListNode{Value: value, Next: sl.head}
	sl.size++
}

func (sl *SinglyList) AddBack(value string) {
	newNode := &ListNode{Value: value}
	
	if sl.head == nil {
		sl.head = newNode
	} else {
		current := sl.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	sl.size++
}

func (sl *SinglyList) Remove(value string) bool {
	if sl.head == nil {
		return false
	}
	
	// Удаление из головы
	if sl.head.Value == value {
		sl.head = sl.head.Next
		sl.size--
		return true
	}
	
	current := sl.head
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}
	
	if current.Next != nil {
		current.Next = current.Next.Next
		sl.size--
		return true
	}
	
	return false
}

func (sl *SinglyList) Contains(value string) bool {
	current := sl.head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (sl *SinglyList) Size() int {
	return sl.size
}

func (sl *SinglyList) Print() {
	current := sl.head
	for current != nil {
		fmt.Printf("%s -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

// ================== Doubly Linked List ==================

type DoublyList struct {
	head *ListNode
	tail *ListNode
	size int
}

func NewDoublyList() *DoublyList {
	return &DoublyList{}
}

func (dl *DoublyList) AddFront(value string) {
	newNode := &ListNode{Value: value, Next: dl.head}
	
	if dl.head != nil {
		dl.head.Prev = newNode
	} else {
		dl.tail = newNode
	}
	
	dl.head = newNode
	dl.size++
}

func (dl *DoublyList) AddBack(value string) {
	newNode := &ListNode{Value: value, Prev: dl.tail}
	
	if dl.tail != nil {
		dl.tail.Next = newNode
	} else {
		dl.head = newNode
	}
	
	dl.tail = newNode
	dl.size++
}

func (dl *DoublyList) Remove(value string) bool {
	current := dl.head
	
	for current != nil {
		if current.Value == value {
			// Обновляем ссылки соседних узлов
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				dl.head = current.Next
			}
			
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				dl.tail = current.Prev
			}
			
			dl.size--
			return true
		}
		current = current.Next
	}
	
	return false
}

func (dl *DoublyList) Contains(value string) bool {
	current := dl.head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (dl *DoublyList) Size() int {
	return dl.size
}

func (dl *DoublyList) Print() {
	current := dl.head
	for current != nil {
		fmt.Printf("%s <-> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func (dl *DoublyList) PrintReverse() {
	current := dl.tail
	for current != nil {
		fmt.Printf("%s <-> ", current.Value)
		current = current.Prev
	}
	fmt.Println("nil")
}

func main() {
	fmt.Println("=== Singly Linked List ===")
	slist := NewSinglyList()
	slist.AddFront("C")
	slist.AddFront("B")
	slist.AddFront("A")
	slist.AddBack("D")
	slist.Print()
	fmt.Printf("Size: %d\n", slist.Size())
	fmt.Printf("Contains 'B': %v\n", slist.Contains("B"))
	
	slist.Remove("B")
	slist.Print()
	
	fmt.Println("\n=== Doubly Linked List ===")
	dlist := NewDoublyList()
	dlist.AddFront("3")
	dlist.AddFront("2")
	dlist.AddFront("1")
	dlist.AddBack("4")
	dlist.Print()
	dlist.PrintReverse()
	fmt.Printf("Size: %d\n", dlist.Size())
	
	dlist.Remove("2")
	dlist.Print()
}
