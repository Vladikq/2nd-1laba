#include "../include/tree.h"

NodeT::NodeT(int value) {
    data = value;
    left = nullptr;
    right = nullptr;
}

FullBinaryTree::FullBinaryTree() {
    root = nullptr;
    size = 0;
}

FullBinaryTree::~FullBinaryTree() {
    clear(root);
}

void FullBinaryTree::print() {
    printTree(root, 0);
    cout << endl;
}

string FullBinaryTree::toString() {
    return _toString(root);
}

void FullBinaryTree::insert(int value) {
    root = _insert(root, value);
    size++;
}

NodeT* FullBinaryTree::_insert(NodeT* node, int value) {
    if (node == nullptr) {
        return new NodeT(value);
    }
    if (value < node->data) {
        node->left = _insert(node->left, value);
    } else if (value > node->data) {
        node->right = _insert(node->right, value);
    }
    // Если значение уже существует, ничего не делаем
    return node;
}

// УДАЛЕНИЕ УЗЛА
void FullBinaryTree::remove(int value) {
    root = _remove(root, value);
    if(root) size--;
}

NodeT* FullBinaryTree::_remove(NodeT* node, int value) {
    if (node == nullptr) return nullptr;
    
    if (value < node->data) {
        node->left = _remove(node->left, value);
    } else if (value > node->data) {
        node->right = _remove(node->right, value);
    } else {
        // Найден узел для удаления
        
        // Случай 1: узел без потомков или с одним потомком
        if (node->left == nullptr) {
            NodeT* temp = node->right;
            delete node;
            return temp;
        } else if (node->right == nullptr) {
            NodeT* temp = node->left;
            delete node;
            return temp;
        }
        
        // Случай 2: узел с двумя потомками
        // Находим минимальный элемент в правом поддереве
        NodeT* temp = _findMin(node->right);
        node->data = temp->data;
        node->right = _remove(node->right, temp->data);
    }
    return node;
}

// Поиск минимального элемента
NodeT* FullBinaryTree::_findMin(NodeT* node) {
    while (node && node->left != nullptr) {
        node = node->left;
    }
    return node;
}

// Улучшенный поиск (для BST)
bool FullBinaryTree::search(NodeT* node, int value) {
    if (node == nullptr) return false;
    if (node->data == value) return true;
    
    if (value < node->data) {
        return search(node->left, value);
    } else {
        return search(node->right, value);
    }
}

// НОВЫЙ МЕТОД: поиск с возвратом bool
bool FullBinaryTree::search(int value) {
    return search(root, value);
}

bool FullBinaryTree::isFull() {
    return isFull(root);
}

bool FullBinaryTree::isFull(NodeT* node) {
    if (node == nullptr) return true;
    
    if (node->left == nullptr && node->right == nullptr) return true;
    if (node->left != nullptr && node->right != nullptr) {
        return isFull(node->left) && isFull(node->right);
    }
    return false;
}

// НОВЫЙ МЕТОД: проверка на полноту (все узлы имеют 0 или 2 потомка)
bool FullBinaryTree::isPerfect() {
    int height = getHeight(root);
    return _isPerfect(root, height, 0);
}

int FullBinaryTree::getHeight(NodeT* node) {
    if (node == nullptr) return 0;
    return 1 + max(getHeight(node->left), getHeight(node->right));
}

bool FullBinaryTree::_isPerfect(NodeT* node, int height, int level) {
    if (node == nullptr) return true;
    
    if (node->left == nullptr && node->right == nullptr) {
        return (height == level + 1);
    }
    
    if (node->left == nullptr || node->right == nullptr) {
        return false;
    }
    
    return _isPerfect(node->left, height, level + 1) &&
           _isPerfect(node->right, height, level + 1);
}

string FullBinaryTree::_toString(NodeT* node) {
    if (node == nullptr) return "";
    ostringstream oss;
    oss << node->data << " ";
    oss << _toString(node->left);
    oss << _toString(node->right);
    return oss.str();
}

// НОВЫЙ МЕТОД: обход в порядке in-order
string FullBinaryTree::inOrder() {
    return _inOrder(root);
}

string FullBinaryTree::_inOrder(NodeT* node) {
    if (node == nullptr) return "";
    ostringstream oss;
    oss << _inOrder(node->left);
    oss << node->data << " ";
    oss << _inOrder(node->right);
    return oss.str();
}

void FullBinaryTree::printTree(NodeT* node, int depth) {
    if (node == nullptr) return;
    printTree(node->right, depth + 1);
    cout << setw(4 * depth) << " " << node->data << endl;
    printTree(node->left, depth + 1);
}

void FullBinaryTree::clear(NodeT* node) {
    if (node == nullptr) return;
    clear(node->left);
    clear(node->right);
    delete node;
}

// НОВЫЙ МЕТОД: очистка всего дерева
void FullBinaryTree::clear() {
    clear(root);
    root = nullptr;
    size = 0;
}

// НОВЫЙ МЕТОД: получение размера
size_t FullBinaryTree::getSize() const {
    return size;
}
