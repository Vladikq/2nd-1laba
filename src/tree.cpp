#include "../include/tree.h"


NodeT::NodeT(int value) { 
    data = value;
    left = nullptr;
    right = nullptr;                                
}

FullBinaryTree::FullBinaryTree() { 
    root = nullptr;
    size = 0;                                       // Пустое дерево при создании
}

void FullBinaryTree::print() { 
    printTree(root, 0);
    cout << endl;                                   // Запуск рекурсивного вывода
}

string FullBinaryTree::toString() {
    return _toString(root);                         // Обертка для рекурсивного метода
}

void FullBinaryTree::insert(int value) {
    root = _insert(root, value);                    // Рекурсивная вставка с обновлением корня
    size++;                                         // Увеличение счетчика элементов
}

NodeT* FullBinaryTree::_insert(NodeT* node, int value) { 
    if (node == nullptr) { 
        return new NodeT(value);                    // Базовый случай: создание листа
    }
    if (value < node->data) { 
        node->left = _insert(node->left, value);    // Рекурсия в левое поддерево
    } else {
        node->right = _insert(node->right, value);  // Рекурсия в правое поддерево
    }
    return node;                                    // Возврат текущего узла
}

bool FullBinaryTree::search(NodeT* node, int value) { 
    if (node == nullptr) return false;              // Базовый случай: не найдено
    if (node->data == value) return true;           // Базовый случай: значение найдено
    return search(node->left, value) || search(node->right, value);  // Рекурсивный поиск
}

bool FullBinaryTree::isFull() { 
    return isFull(root);                            // Запуск проверки с корня
}

bool FullBinaryTree::isFull(NodeT* node) { 
    if (node == nullptr) return true;               // Пустое поддерево - полное
    
    if (node->left == nullptr && node->right == nullptr) return true; // Узел-лист
    if (node->left != nullptr && node->right != nullptr) {            // Узел с двумя потомками
        return isFull(node->left) && isFull(node->right);             // Рекурсивная проверка
    }
    return false;                                    // только один потомок
}

string FullBinaryTree::_toString(NodeT* node) { 
    if (node == nullptr) return "";                 
    ostringstream oss;  
    oss << node->data << " ";                       // Pre-order обход (корень-лево-право)
    oss << _toString(node->left);                   
    oss << _toString(node->right);                  
    return oss.str();
}

void FullBinaryTree::printTree(NodeT* node, int depth) { 
    if (node == nullptr) return;                    
    printTree(node->right, depth + 1);              // Сначала правое поддерево
    cout << setw(4 * depth) << " " << node->data << endl; // Форматированный вывод
    printTree(node->left, depth + 1);               // Затем левое поддерево
}

void FullBinaryTree::clear(NodeT* node) { 
    if (node == nullptr) return;                    // Базовый случай
    clear(node->left);                              
    clear(node->right);                             
    delete node;                                    // Удаление текущего узла
}

FullBinaryTree::~FullBinaryTree() { 
    clear(root);                                    // Запуск рекурсивного удаления
}
