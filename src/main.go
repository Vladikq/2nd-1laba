package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ================== Вспомогательные функции ==================

// printUsage выводит справку по использованию
func printUsage(programName string) {
	fmt.Printf("Использование: %s --file <filename> --query 'command'\n", programName)
}

// readFile читает файл и возвращает его содержимое как слайс строк
func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeFile записывает содержимое в файл
func writeFile(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// removeStruct удаляет структуру данных из файла
func removeStruct(lines []string, structName string) []string {
	var result []string
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 0 || parts[0] != structName {
			result = append(result, line)
		}
	}
	return result
}

// findStruct находит строку со структурой данных
func findStruct(lines []string, structName string) (string, bool) {
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) > 0 && parts[0] == structName {
			return line, true
		}
	}
	return "", false
}

// parseStruct разбирает строку структуры данных
func parseStruct(line string) (string, []string) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return "", nil
	}
	return parts[0], parts[1:]
}

// ================== Массив (Array) ==================

type Array struct {
	elements []string
}

func NewArray() *Array {
	return &Array{elements: make([]string, 0)}
}

func (a *Array) AddToEnd(value string) {
	a.elements = append(a.elements, value)
}

func (a *Array) Add(index int, value string) error {
	if index < 0 || index > len(a.elements) {
		return errors.New("индекс вне диапазона")
	}
	a.elements = append(a.elements, "")
	copy(a.elements[index+1:], a.elements[index:])
	a.elements[index] = value
	return nil
}

func (a *Array) Get(index int) (string, error) {
	if index < 0 || index >= len(a.elements) {
		return "", errors.New("индекс вне диапазона")
	}
	return a.elements[index], nil
}

func (a *Array) Remove(index int) error {
	if index < 0 || index >= len(a.elements) {
		return errors.New("индекс вне диапазона")
	}
	a.elements = append(a.elements[:index], a.elements[index+1:]...)
	return nil
}

func (a *Array) Replace(index int, value string) error {
	if index < 0 || index >= len(a.elements) {
		return errors.New("индекс вне диапазона")
	}
	a.elements[index] = value
	return nil
}

func (a *Array) Size() int {
	return len(a.elements)
}

func (a *Array) ShowArray() {
	for _, elem := range a.elements {
		fmt.Println(elem)
	}
}

func (a *Array) String() string {
	return strings.Join(a.elements, " ")
}

// Функции для работы с массивом в файле
func readArrayFromFile(lines []string, name string) *Array {
	arr := NewArray()
	if line, found := findStruct(lines, name); found {
		_, values := parseStruct(line)
		for _, val := range values {
			arr.AddToEnd(val)
		}
	}
	return arr
}

func saveArrayToFile(filename, name string, arr *Array) error {
	lines, err := readFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	
	// Удаляем старую структуру
	lines = removeStruct(lines, name)
	
	// Добавляем новую
	if arr.Size() > 0 {
		newLine := name + " " + arr.String()
		lines = append(lines, newLine)
	}
	
	return writeFile(filename, lines)
}

// Команды для массива
func processArrayCommand(command, filename string) error {
	parts := strings.Fields(command)
	if len(parts) < 2 {
		return errors.New("неверная команда")
	}
	
	lines, _ := readFile(filename)
	arrName := parts[1]
	arr := readArrayFromFile(lines, arrName)
	
	switch parts[0] {
	case "MPUSH":
		if len(parts) < 3 {
			return errors.New("недостаточно аргументов")
		}
		arr.AddToEnd(parts[2])
		return saveArrayToFile(filename, arrName, arr)
		
	case "MPUSHIND":
		if len(parts) < 4 {
			return errors.New("недостаточно аргументов")
		}
		index, err := strconv.Atoi(parts[3])
		if err != nil {
			return errors.New("неверный индекс")
		}
		if err := arr.Add(index, parts[2]); err != nil {
			return err
		}
		return saveArrayToFile(filename, arrName, arr)
		
	case "MREMOVE":
		if len(parts) < 3 {
			return errors.New("недостаточно аргументов")
		}
		index, err := strconv.Atoi(parts[2])
		if err != nil {
			return errors.New("неверный индекс")
		}
		if err := arr.Remove(index); err != nil {
			return err
		}
		return saveArrayToFile(filename, arrName, arr)
		
	case "MREPLACE":
		if len(parts) < 4 {
			return errors.New("недостаточно аргументов")
		}
		index, err := strconv.Atoi(parts[3])
		if err != nil {
			return errors.New("неверный индекс")
		}
		if err := arr.Replace(index, parts[2]); err != nil {
			return err
		}
		return saveArrayToFile(filename, arrName, arr)
		
	case "MGET":
		if len(parts) < 3 {
			return errors.New("недостаточно аргументов")
		}
		index, err := strconv.Atoi(parts[2])
		if err != nil {
			return errors.New("неверный индекс")
		}
		value, err := arr.Get(index)
		if err != nil {
			return err
		}
		fmt.Println(value)
		return nil
		
	case "MSIZE":
		fmt.Println(arr.Size())
		return nil
		
	case "MPRINT":
		arr.ShowArray()
		return nil
		
	default:
		return errors.New("неизвестная команда")
	}
}

// ================== Односвязный список ==================

type SNode struct {
	Data string
	Next *SNode
}

type SinglyLinkedList struct {
	head *SNode
	size int
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (sll *SinglyLinkedList) IsEmpty() bool {
	return sll.size == 0
}

func (sll *SinglyLinkedList) PushFront(value string) {
	newNode := &SNode{Data: value, Next: sll.head}
	sll.head = newNode
	sll.size++
}

func (sll *SinglyLinkedList) PushBack(value string) {
	newNode := &SNode{Data: value}
	if sll.head == nil {
		sll.head = newNode
	} else {
		current := sll.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	sll.size++
}

func (sll *SinglyLinkedList) PopFront() error {
	if sll.head == nil {
		return errors.New("список пуст")
	}
	sll.head = sll.head.Next
	sll.size--
	return nil
}

func (sll *SinglyLinkedList) PopBack() error {
	if sll.head == nil {
		return errors.New("список пуст")
	}
	if sll.head.Next == nil {
		sll.head = nil
	} else {
		current := sll.head
		for current.Next.Next != nil {
			current = current.Next
		}
		current.Next = nil
	}
	sll.size--
	return nil
}

func (sll *SinglyLinkedList) Remove(value string) bool {
	if sll.head == nil {
		return false
	}
	
	if sll.head.Data == value {
		sll.head = sll.head.Next
		sll.size--
		return true
	}
	
	current := sll.head
	for current.Next != nil && current.Next.Data != value {
		current = current.Next
	}
	
	if current.Next != nil {
		current.Next = current.Next.Next
		sll.size--
		return true
	}
	
	return false
}

func (sll *SinglyLinkedList) Find(value string) bool {
	current := sll.head
	for current != nil {
		if current.Data == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (sll *SinglyLinkedList) Print() {
	current := sll.head
	for current != nil {
		fmt.Printf("%s ", current.Data)
		current = current.Next
	}
	fmt.Println()
}

func (sll *SinglyLinkedList) String() string {
	var result []string
	current := sll.head
	for current != nil {
		result = append(result, current.Data)
		current = current.Next
	}
	return strings.Join(result, " ")
}

func (sll *SinglyLinkedList) Size() int {
	return sll.size
}

// Функции для работы со списком в файле
func readSListFromFile(lines []string, name string) *SinglyLinkedList {
	sll := NewSinglyLinkedList()
	if line, found := findStruct(lines, name); found {
		_, values := parseStruct(line)
		for _, val := range values {
			sll.PushBack(val)
		}
	}
	return sll
}

func saveSListToFile(filename, name string, sll *SinglyLinkedList) error {
	lines, err := readFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	
	lines = removeStruct(lines, name)
	
	if sll.Size() > 0 {
		newLine := name + " " + sll.String()
		lines = append(lines, newLine)
	}
	
	return writeFile(filename, lines)
}

// ================== Двусвязный список ==================

type DNode struct {
	Data string
	Next *DNode
	Prev *DNode
}

type DoublyLinkedList struct {
	head *DNode
	tail *DNode
	size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.size == 0
}

func (dll *DoublyLinkedList) PushFront(value string) {
	newNode := &DNode{Data: value, Next: dll.head}
	if dll.head != nil {
		dll.head.Prev = newNode
	} else {
		dll.tail = newNode
	}
	dll.head = newNode
	dll.size++
}

func (dll *DoublyLinkedList) PushBack(value string) {
	newNode := &DNode{Data: value, Prev: dll.tail}
	if dll.tail != nil {
		dll.tail.Next = newNode
	} else {
		dll.head = newNode
	}
	dll.tail = newNode
	dll.size++
}

func (dll *DoublyLinkedList) PopFront() error {
	if dll.head == nil {
		return errors.New("список пуст")
	}
	
	dll.head = dll.head.Next
	if dll.head != nil {
		dll.head.Prev = nil
	} else {
		dll.tail = nil
	}
	dll.size--
	return nil
}

func (dll *DoublyLinkedList) PopBack() error {
	if dll.tail == nil {
		return errors.New("список пуст")
	}
	
	dll.tail = dll.tail.Prev
	if dll.tail != nil {
		dll.tail.Next = nil
	} else {
		dll.head = nil
	}
	dll.size--
	return nil
}

func (dll *DoublyLinkedList) Remove(value string) bool {
	current := dll.head
	for current != nil {
		if current.Data == value {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				dll.head = current.Next
			}
			
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				dll.tail = current.Prev
			}
			
			dll.size--
			return true
		}
		current = current.Next
	}
	return false
}

func (dll *DoublyLinkedList) Find(value string) bool {
	current := dll.head
	for current != nil {
		if current.Data == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (dll *DoublyLinkedList) Print() {
	current := dll.head
	for current != nil {
		fmt.Printf("%s ", current.Data)
		current = current.Next
	}
	fmt.Println()
}

func (dll *DoublyLinkedList) String() string {
	var result []string
	current := dll.head
	for current != nil {
		result = append(result, current.Data)
		current = current.Next
	}
	return strings.Join(result, " ")
}

func (dll *DoublyLinkedList) Size() int {
	return dll.size
}

// Функции для работы с двусвязным списком в файле
func readDListFromFile(lines []string, name string) *DoublyLinkedList {
	dll := NewDoublyLinkedList()
	if line, found := findStruct(lines, name); found {
		_, values := parseStruct(line)
		for _, val := range values {
			dll.PushBack(val)
		}
	}
	return dll
}

func saveDListToFile(filename, name string, dll *DoublyLinkedList) error {
	lines, err := readFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	
	lines = removeStruct(lines, name)
	
	if dll.Size() > 0 {
		newLine := name + " " + dll.String()
		lines = append(lines, newLine)
	}
	
	return writeFile(filename, lines)
}

// Команды для списков
func processListCommand(command, filename string) error {
	parts := strings.Fields(command)
	if len(parts) < 2 {
		return errors.New("неверная команда")
	}
	
	lines, _ := readFile(filename)
	
	switch parts[0] {
	// Односвязный список
	case "LPUSHS":
		if len(parts) < 4 {
			return errors.New("недостаточно аргументов")
		}
		sll := readSListFromFile(lines, parts[1])
		if parts[2] == "back" {
			sll.PushBack(parts[3])
		} else if parts[2] == "front" {
			sll.PushFront(parts[3])
		}
		return saveSListToFile(filename, parts[1], sll)
		
	case "LPOPS":
		if len(parts) < 3 {
			return errors.New("недостаточно аргументов")
		}
		sll := readSListFromFile(lines, parts[1])
		if parts[2] == "back" {
			if err := sll.PopBack(); err != nil {
				return err
			}
		} else if parts[2] == "front" {
			if err := sll.PopFront(); err != nil {
				return err
			}
		}
		return saveSListToFile(filename, parts[1], sll)
		
	case "LREMOVES":
		if len(parts) < 3 {
			return errors.New("недостаточно аргументов")
		}
		sll := readSListFromFile(lines, parts[1])
		if !sll.Remove(parts[2]) {
			return errors.New("элемент не найден")
		}
		return saveSListToFile(filename, parts[1], sll)
		
	case "LGETS":
		if len(parts) < 3 {
			return errors.New("недостаточно аргументов")
		}
		sll := readSListFromFile(lines, parts[1])
		found := sll.Find(parts[2])
		fmt.Println(found)
		return nil
		
	case "LPRINTS":
		if len(parts) < 2 {
			return errors.New("недостаточно аргументов")
		}
		sll := readSListFromFile(lines, parts[1])
		sll.Print()
		return nil
	
	// Двусвязный список
	case "LPUSHB", "LPUSHF":
		if len(parts) < 3 {
			return errors.New("недостаточно аргументов")
		}
		dll := readDListFromFile(lines, parts[1])
		if parts[0] == "LPUSHB" {
			dll.PushBack(parts[2])
		} else {
			dll.PushFront(parts[2])
		}
		return saveDListToFile(filename, parts[1], dll)
		
	case "LPOPB", "LPOPF":
		if len(parts) < 2 {
			return errors.New("недостаточно аргументов")
		}
		dll := readDListFromFile(lines, parts[1])
		if parts[0] == "LPOPB" {
			if err := dll.PopBack(); err != nil {
				return err
			}
		} else {
			if err := dll.PopFront(); err != nil {
				return err
			}
		}
		return saveDListToFile(filename, parts[1], dll)
		
	case "LREMOVE":
		if len(parts) < 3 {
			return errors.New("недостаточно аргументов")
		}
		dll := readDListFromFile(lines, parts[1])
		if !dll.Remove(parts[2]) {
			return errors.New("элемент не найден")
		}
		return saveDListToFile(filename, parts[1], dll)
		
	case "LGET":
		if len(parts) < 3 {
			return errors.New("недостаточно аргументов")
		}
		dll := readDListFromFile(lines, parts[1])
		found := dll.Find(parts[2])
		fmt.Println(found)
		return nil
		
	case "LPRINT":
		if len(parts) < 2 {
			return errors.New("недостаточно аргументов")
		}
		dll := readDListFromFile(lines, parts[1])
		dll.Print()
		return nil
		
	default:
		return errors.New("неизвестная команда")
	}
}

// ================== Стек (Stack) ==================

type Stack struct {
	elements []string
}

func NewStack() *Stack {
	return &Stack{elements: make([]string, 0)}
}

func (s *Stack) Push(value string) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() (string, error) {
	if len(s.elements) == 0 {
		return "", errors.New("стек пуст")
	}
	lastIdx := len(s.elements) - 1
	value := s.elements[lastIdx]
	s.elements = s.elements[:lastIdx]
	return value, nil
}

func (s *Stack) Peek() (string, error) {
	if len(s.elements) == 0 {
		return "", errors.New("стек пуст")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) String() string {
	return strings.Join(s.elements, " ")
}

// Функции для работы со стеком в файле
func readStackFromFile(lines []string, name string) *Stack {
	stack := NewStack()
	if line, found := findStruct(lines, name); found {
		_, values := parseStruct(line)
		for _, val := range values {
			stack.Push(val)
		}
	}
	return stack
}

func saveStackToFile(filename, name string, stack *Stack) error {
	lines, err := readFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	
	lines = removeStruct(lines, name)
	
	if stack.Size() > 0 {
		// Сохраняем в обратном порядке для правильного восстановления
		tempStack := NewStack()
		elements := make([]string, stack.Size())
		copy(elements, stack.elements)
		for i := len(elements) - 1; i >= 0; i-- {
			tempStack.Push(elements[i])
		}
		newLine := name + " " + tempStack.String()
		lines = append(lines, newLine)
	}
	
	return writeFile(filename, lines)
}

// Команды для стека
func processStackCommand(command, filename string) error {
	parts := strings.Fields(command)
	if len(parts) < 2 {
		return errors.New("неверная команда")
	}
	
	lines, _ := readFile(filename)
	stackName := parts[1]
	stack := readStackFromFile(lines, stackName)
	
	switch parts[0] {
	case "SPUSH":
		if len(parts) < 3 {
			return errors.New("недостаточно аргументов")
		}
		stack.Push(parts[2])
		return saveStackToFile(filename, stackName, stack)
		
	case "SPOP":
		if _, err := stack.Pop(); err != nil {
			return err
		}
		return saveStackToFile(filename, stackName, stack)
		
	case "SPRINT":
		fmt.Println(stack.String())
		return nil
		
	default:
		return errors.New("неизвестная команда")
	}
}

// ================== Очередь (Queue) ==================

type Queue struct {
	elements []string
}

func NewQueue() *Queue {
	return &Queue{elements: make([]string, 0)}
}

func (q *Queue) Push(value string) {
	q.elements = append(q.elements, value)
}

func (q *Queue) Pop() (string, error) {
	if len(q.elements) == 0 {
		return "", errors.New("очередь пуста")
	}
	value := q.elements[0]
	q.elements = q.elements[1:]
	return value, nil
}

func (q *Queue) Peek() (string, error) {
	if len(q.elements) == 0 {
		return "", errors.New("очередь пуста")
	}
	return q.elements[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func (q *Queue) String() string {
	return strings.Join(q.elements, " ")
}

// Функции для работы с очередью в файле
func readQueueFromFile(lines []string, name string) *Queue {
	queue := NewQueue()
	if line, found := findStruct(lines, name); found {
		_, values := parseStruct(line)
		for _, val := range values {
			queue.Push(val)
		}
	}
	return queue
}

func saveQueueToFile(filename, name string, queue *Queue) error {
	lines, err := readFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	
	lines = removeStruct(lines, name)
	
	if queue.Size() > 0 {
		newLine := name + " " + queue.String()
		lines = append(lines, newLine)
	}
	
	return writeFile(filename, lines)
}

// Команды для очереди
func processQueueCommand(command, filename string) error {
	parts := strings.Fields(command)
	if len(parts) < 2 {
		return errors.New("неверная команда")
	}
	
	lines, _ := readFile(filename)
	queueName := parts[1]
	queue := readQueueFromFile(lines, queueName)
	
	switch parts[0] {
	case "QPUSH":
		if len(parts) < 3 {
			return errors.New("недостаточно аргументов")
		}
		queue.Push(parts[2])
		return saveQueueToFile(filename, queueName, queue)
		
	case "QPOP":
		if _, err := queue.Pop(); err != nil {
			return err
		}
		return saveQueueToFile(filename, queueName, queue)
		
	case "QPRINT":
		fmt.Println(queue.String())
		return nil
		
	default:
		return errors.New("неизвестная команда")
	}
}

// ================== Дерево (Binary Tree) ==================

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
	size int
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (bt *BinaryTree) Insert(value int) {
	bt.Root = bt.insertRec(bt.Root, value)
	bt.size++
}

func (bt *BinaryTree) insertRec(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{Value: value}
	}
	
	if value < node.Value {
		node.Left = bt.insertRec(node.Left, value)
	} else if value > node.Value {
		node.Right = bt.insertRec(node.Right, value)
	}
	return node
}

func (bt *BinaryTree) Search(value int) bool {
	return bt.searchRec(bt.Root, value)
}

func (bt *BinaryTree) searchRec(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}
	
	if value == node.Value {
		return true
	} else if value < node.Value {
		return bt.searchRec(node.Left, value)
	} else {
		return bt.searchRec(node.Root, value)
	}
}

func (bt *BinaryTree) IsFull() bool {
	return bt.isFullRec(bt.Root)
}

func (bt *BinaryTree) isFullRec(node *TreeNode) bool {
	if node == nil {
		return true
	}
	
	if node.Left == nil && node.Right == nil {
		return true
	}
	
	if node.Left != nil && node.Right != nil {
		return bt.isFullRec(node.Left) && bt.isFullRec(node.Right)
	}
	
	return false
}

func (bt *BinaryTree) InOrder() []int {
	var result []int
	bt.inOrderRec(bt.Root, &result)
	return result
}

func (bt *BinaryTree) inOrderRec(node *TreeNode, result *[]int) {
	if node != nil {
		bt.inOrderRec(node.Left, result)
		*result = append(*result, node.Value)
		bt.inOrderRec(node.Right, result)
	}
}

func (bt *BinaryTree) Print() {
	values := bt.InOrder()
	for _, val := range values {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

func (bt *BinaryTree) String() string {
	values := bt.InOrder()
	strValues := make([]string, len(values))
	for i, val := range values {
		strValues[i] = strconv.Itoa(val)
	}
	return strings.Join(strValues, " ")
}

func (bt *BinaryTree) Size() int {
	return bt.size
}

// Функции для работы с деревом в файле
func readTreeFromFile(lines []string, name string) *BinaryTree {
	tree := NewBinaryTree()
	if line, found := findStruct(lines, name); found {
		_, values := parseStruct(line)
		for _, valStr := range values {
			if val, err := strconv.Atoi(valStr); err == nil {
				tree.Insert(val)
			}
		}
	}
	return tree
}

func saveTreeToFile(filename, name string, tree *BinaryTree) error {
	lines, err := readFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	
	lines = removeStruct(lines, name)
	
	if tree.Size() > 0 {
		newLine := name + " " + tree.String()
		lines = append(lines, newLine)
	}
	
	return writeFile(filename, lines)
}

// Команды для дерева
func processTreeCommand(command, filename string) error {
	parts := strings.Fields(command)
	if len(parts) < 2 {
		return errors.New("неверная команда")
	}
	
	lines, _ := readFile(filename)
	treeName := parts[1]
	tree := readTreeFromFile(lines, treeName)
	
	switch parts[0] {
	case "TPUSH":
		if len(parts) < 3 {
			return errors.New("недостаточно аргументов")
		}
		value, err := strconv.Atoi(parts[2])
		if err != nil {
			return errors.New("неверное значение")
		}
		tree.Insert(value)
		return saveTreeToFile(filename, treeName, tree)
		
	case "TSEARCH":
		if len(parts) < 3 {
			return errors.New("недостаточно аргументов")
		}
		value, err := strconv.Atoi(parts[2])
		if err != nil {
			return errors.New("неверное значение")
		}
		found := tree.Search(value)
		fmt.Println(found)
		return nil
		
	case "TCHECK":
		fmt.Println(tree.IsFull())
		return nil
		
	case "TPRINT":
		tree.Print()
		return nil
		
	default:
		return errors.New("неизвестная команда")
	}
}

// ================== Основная функция ==================

func main() {
	if len(os.Args) != 5 {
		printUsage(os.Args[0])
		os.Exit(1)
	}

	var filename, query string

	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "--file" {
			if i+1 < len(os.Args) {
				filename = os.Args[i+1]
				i++
			} else {
				printUsage(os.Args[0])
				os.Exit(1)
			}
		} else if os.Args[i] == "--query" {
			if i+1 < len(os.Args) {
				query = os.Args[i+1]
				i++
			} else {
				printUsage(os.Args[0])
				os.Exit(1)
			}
		}
	}

	if query == "" {
		fmt.Println("Ошибка: должна быть указана команда")
		os.Exit(1)
	}

	var err error
	
	switch query[0] {
	case 'M':
		err = processArrayCommand(query, filename)
	case 'L':
		err = processListCommand(query, filename)
	case 'S':
		err = processStackCommand(query, filename)
	case 'Q':
		err = processQueueCommand(query, filename)
	case 'T':
		err = processTreeCommand(query, filename)
	default:
		fmt.Println("Ошибка: неизвестная структура данных")
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}
}
