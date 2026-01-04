package tree

import (
	"fmt"
	"strings"
)

// Node представляет узел бинарного дерева
type Node struct {
	data  int
	left  *Node
	right *Node
}

// NewNode создает новый узел с указанным значением
func NewNode(value int) *Node {
	return &Node{
		data:  value,
		left:  nil,
		right: nil,
	}
}

// FullBinaryTree представляет полное бинарное дерево
type FullBinaryTree struct {
	root *Node
	size int
}

// NewFullBinaryTree создает новое пустое бинарное дерево
func NewFullBinaryTree() *FullBinaryTree {
	return &FullBinaryTree{
		root: nil,
		size: 0,
	}
}

// Insert добавляет значение в дерево
func (t *FullBinaryTree) Insert(value int) {
	t.root = t.insert(t.root, value)
	t.size++
}

// insert рекурсивно вставляет значение в дерево
func (t *FullBinaryTree) insert(node *Node, value int) *Node {
	if node == nil {
		return NewNode(value)
	}
	
	if value < node.data {
		node.left = t.insert(node.left, value)
	} else {
		node.right = t.insert(node.right, value)
	}
	
	return node
}

// Search ищет значение в дереве
func (t *FullBinaryTree) Search(value int) bool {
	return t.search(t.root, value)
}

// search рекурсивно ищет значение в дереве
func (t *FullBinaryTree) search(node *Node, value int) bool {
	if node == nil {
		return false
	}
	
	if node.data == value {
		return true
	}
	
	// Рекурсивный поиск в обоих поддеревьях
	return t.search(node.left, value) || t.search(node.right, value)
}

// IsFull проверяет, является ли дерево полным
func (t *FullBinaryTree) IsFull() bool {
	return t.isFull(t.root)
}

// isFull рекурсивно проверяет, является ли поддерево полным
func (t *FullBinaryTree) isFull(node *Node) bool {
	if node == nil {
		return true
	}
	
	// Если это лист
	if node.left == nil && node.right == nil {
		return true
	}
	
	// Если есть оба потомка, проверяем рекурсивно
	if node.left != nil && node.right != nil {
		return t.isFull(node.left) && t.isFull(node.right)
	}
	
	// Если только один потомок - дерево не полное
	return false
}

// Print выводит дерево в консоль
func (t *FullBinaryTree) Print() {
	t.printTree(t.root, 0)
	fmt.Println()
}

// printTree рекурсивно выводит дерево с отступами
func (t *FullBinaryTree) printTree(node *Node, depth int) {
	if node == nil {
		return
	}
	
	// Сначала правое поддерево
	t.printTree(node.right, depth+1)
	
	// Вывод текущего узла с отступами
	indent := strings.Repeat("    ", depth)
	fmt.Printf("%s%d\n", indent, node.data)
	
	// Затем левое поддерево
	t.printTree(node.left, depth+1)
}

// ToString возвращает строковое представление дерева (pre-order обход)
func (t *FullBinaryTree) ToString() string {
	var result strings.Builder
	t.toString(t.root, &result)
	return result.String()
}

// toString рекурсивно собирает строковое представление
func (t *FullBinaryTree) toString(node *Node, sb *strings.Builder) {
	if node == nil {
		return
	}
	
	sb.WriteString(fmt.Sprintf("%d ", node.data))
	t.toString(node.left, sb)
	t.toString(node.right, sb)
}

// Size возвращает количество элементов в дереве
func (t *FullBinaryTree) Size() int {
	return t.size
}

// Clear очищает дерево
func (t *FullBinaryTree) Clear() {
	t.clear(t.root)
	t.root = nil
	t.size = 0
}

// clear рекурсивно очищает память узлов
func (t *FullBinaryTree) clear(node *Node) {
	if node == nil {
		return
	}
	
	t.clear(node.left)
	t.clear(node.right)
	// В Go не нужно явно освобождать память
}

// Вспомогательные методы для обходов дерева

// InOrder возвращает значения при инфиксном обходе (левый-корень-правый)
func (t *FullBinaryTree) InOrder() []int {
	result := make([]int, 0, t.size)
	t.inOrder(t.root, &result)
	return result
}

func (t *FullBinaryTree) inOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	
	t.inOrder(node.left, result)
	*result = append(*result, node.data)
	t.inOrder(node.right, result)
}

// PreOrder возвращает значения при префиксном обходе (корень-левый-правый)
func (t *FullBinaryTree) PreOrder() []int {
	result := make([]int, 0, t.size)
	t.preOrder(t.root, &result)
	return result
}

func (t *FullBinaryTree) preOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	
	*result = append(*result, node.data)
	t.preOrder(node.left, result)
	t.preOrder(node.right, result)
}

// PostOrder возвращает значения при постфиксном обходе (левый-правый-корень)
func (t *FullBinaryTree) PostOrder() []int {
	result := make([]int, 0, t.size)
	t.postOrder(t.root, &result)
	return result
}

func (t *FullBinaryTree) postOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	
	t.postOrder(node.left, result)
	t.postOrder(node.right, result)
	*result = append(*result, node.data)
}

// FindMin возвращает минимальное значение в дереве
func (t *FullBinaryTree) FindMin() (int, error) {
	if t.root == nil {
		return 0, fmt.Errorf("tree is empty")
	}
	
	node := t.root
	for node.left != nil {
		node = node.left
	}
	return node.data, nil
}

// FindMax возвращает максимальное значение в дереве
func (t *FullBinaryTree) FindMax() (int, error) {
	if t.root == nil {
		return 0, fmt.Errorf("tree is empty")
	}
	
	node := t.root
	for node.right != nil {
		node = node.right
	}
	return node.data, nil
}
