#include "../include/list.h"

// Реализация конструкторов узлов
Node::Node(string value) { 
    data = value;
    next = nullptr;
}

DoubleNode::DoubleNode(string value) {  
    data = value;
    next = nullptr;
    prev = nullptr;
}

// SinglyLinkedList
SinglyLinkedList::SinglyLinkedList() { 
    head = nullptr;
    elementCount = 0;
}

SinglyLinkedList::~SinglyLinkedList() { 
    clearSList(); // Используем clearSList вместо popFront в цикле
}

bool SinglyLinkedList::isEmpty() const { 
    return elementCount == 0;
}

void SinglyLinkedList::print() { 
    Node* current = head;
    while(current) { 
        cout << current->data << " ";  
        current = current->next;
    }
    cout << endl;
}

void SinglyLinkedList::pushFront(string value) { 
    Node* newNode = new Node(value);
    newNode->next = head;
    head = newNode;
    elementCount++;
}

void SinglyLinkedList::pushBack(string value) {
    Node* newNode = new Node(value);
    if(head == nullptr) { 
        head = newNode;
    } else {
        Node* current = head; 
        while(current->next != nullptr) { 
            current = current->next;
        }
        current->next = newNode;
    }
    elementCount++;  
}

void SinglyLinkedList::popFront() { 
    if(head == nullptr) return;
    Node* temp = head;
    head = head->next;
    delete temp;
    elementCount--;
}

void SinglyLinkedList::popBack() { 
    if(head == nullptr) return;
    
    if(head->next == nullptr) { 
        // Только один элемент
        delete head;
        head = nullptr;
    } else {
        Node* current = head;
        // Ищем предпоследний элемент
        while(current->next->next != nullptr) {
            current = current->next;
        }
        delete current->next;
        current->next = nullptr;
    }
    elementCount--;
}

void SinglyLinkedList::removeAt(string value) { 
    if(isEmpty()) return;
    
    if(head->data == value) {
        popFront();
        return;
    }
    
    Node* current = head;
    while(current->next != nullptr && current->next->data != value) { 
        current = current->next;
    }
    
    if(current->next != nullptr) {
        Node* nodeToDelete = current->next;
        current->next = nodeToDelete->next;
        delete nodeToDelete;
        elementCount--;
    }
}

bool SinglyLinkedList::find(string value) { 
    Node* current = head;
    while(current != nullptr) {
        if(current->data == value) { 
            return true;
        }
        current = current->next;
    }
    return false;
}

// НОВАЯ ФУНКЦИЯ: Получить элемент по индексу
string SinglyLinkedList::getAt(size_t index) {
    if(index >= elementCount) {
        throw out_of_range("Index out of range");
    }
    
    Node* current = head;
    for(size_t i = 0; i < index; i++) {
        current = current->next;
    }
    return current->data;
}

// НОВАЯ ФУНКЦИЯ: Вставить элемент по индексу
void SinglyLinkedList::insertAt(size_t index, string value) {
    if(index > elementCount) {
        throw out_of_range("Index out of range");
    }
    
    if(index == 0) {
        pushFront(value);
        return;
    }
    
    if(index == elementCount) {
        pushBack(value);
        return;
    }
    
    Node* newNode = new Node(value);
    Node* current = head;
    
    // Находим элемент перед нужной позицией
    for(size_t i = 0; i < index - 1; i++) {
        current = current->next;
    }
    
    newNode->next = current->next;
    current->next = newNode;
    elementCount++;
}

// НОВАЯ ФУНКЦИЯ: Удалить элемент по индексу
void SinglyLinkedList::removeAt(size_t index) {
    if(index >= elementCount) {
        throw out_of_range("Index out of range");
    }
    
    if(index == 0) {
        popFront();
        return;
    }
    
    Node* current = head;
    for(size_t i = 0; i < index - 1; i++) {
        current = current->next;
    }
    
    Node* nodeToDelete = current->next;
    current->next = nodeToDelete->next;
    delete nodeToDelete;
    elementCount--;
}

// НОВАЯ ФУНКЦИЯ: Заменить элемент по индексу
void SinglyLinkedList::replaceAt(size_t index, string value) {
    if(index >= elementCount) {
        throw out_of_range("Index out of range");
    }
    
    Node* current = head;
    for(size_t i = 0; i < index; i++) {
        current = current->next;
    }
    current->data = value;
}

void SinglyLinkedList::clearSList() { 
    while(!isEmpty()) {  
        popFront();
    }
}

Node* SinglyLinkedList::getHead() const { 
    return head;
}

size_t SinglyLinkedList::size() const {
    return elementCount;
}

// НОВАЯ ФУНКЦИЯ: Обратный порядок
void SinglyLinkedList::reverse() {
    if(head == nullptr || head->next == nullptr) return;
    
    Node* prev = nullptr;
    Node* current = head;
    Node* next = nullptr;
    
    while(current != nullptr) {
        next = current->next;
        current->next = prev;
        prev = current;
        current = next;
    }
    head = prev;
}

// DoubleLinkedList
DoubleLinkedList::DoubleLinkedList() { 
    head = nullptr;
    tail = nullptr;
    elementCount = 0;  
}

DoubleLinkedList::~DoubleLinkedList() { 
    clearDList();
}

DoubleNode* DoubleLinkedList::getHead() const {
    return head;
}

DoubleNode* DoubleLinkedList::getTail() const {
    return tail;
}

bool DoubleLinkedList::isEmpty() const { 
    return elementCount == 0;
}

void DoubleLinkedList::pushFront(string value) { 
    DoubleNode* newNode = new DoubleNode(value);
    
    if(head == nullptr) {
        head = tail = newNode;
    } else {
        newNode->next = head;
        head->prev = newNode;
        head = newNode;
    }
    elementCount++; 
}

void DoubleLinkedList::pushBack(string value) { 
    DoubleNode* newNode = new DoubleNode(value);
    
    if(tail == nullptr) { 
        head = tail = newNode;
    } else { 
        newNode->prev = tail;
        tail->next = newNode;
        tail = newNode;
    }
    elementCount++;  
}

void DoubleLinkedList::popFront() { 
    if(head == nullptr) return;
    
    if(head == tail) {
        delete head;
        head = tail = nullptr;
    } else {
        DoubleNode* temp = head;
        head = head->next;
        head->prev = nullptr;
        delete temp;
    }
    elementCount--;  
}

void DoubleLinkedList::popBack() { 
    if(tail == nullptr) return;
    
    if(head == tail) {
        delete tail;
        head = tail = nullptr;
    } else {
        DoubleNode* temp = tail;
        tail = tail->prev;
        tail->next = nullptr;
        delete temp;
    }
    elementCount--;  
}

void DoubleLinkedList::removeAt(string value) { 
    DoubleNode* current = head;
    while(current) { 
        if(current->data == value) { 
            // Удаление из начала
            if(current == head) {
                popFront();
            } 
            // Удаление из конца
            else if(current == tail) {
                popBack();
            } 
            // Удаление из середины
            else {
                current->prev->next = current->next;
                current->next->prev = current->prev;
                delete current;
                elementCount--;
            }
            return;
        }
        current = current->next;
    }
}

bool DoubleLinkedList::find(string value) { 
    DoubleNode* current = head; 
    while(current) { 
        if(current->data == value) { 
            return true;
        }
        current = current->next;
    }
    return false;
}

void DoubleLinkedList::print() { 
    DoubleNode* current = head;
    while(current) { 
        cout << current->data << " ";
        current = current->next;
    }
    cout << endl;
}

void DoubleLinkedList::printReverse() { 
    DoubleNode* current = tail;
    while(current) { 
        cout << current->data << " ";
        current = current->prev;
    }
    cout << endl;
}

// НОВАЯ ФУНКЦИЯ: Получить элемент по индексу
string DoubleLinkedList::getAt(size_t index) {
    if(index >= elementCount) {
        throw out_of_range("Index out of range");
    }
    
    DoubleNode* current;
    if(index < elementCount / 2) {
        // Начинаем с начала
        current = head;
        for(size_t i = 0; i < index; i++) {
            current = current->next;
        }
    } else {
        // Начинаем с конца
        current = tail;
        for(size_t i = elementCount - 1; i > index; i--) {
            current = current->prev;
        }
    }
    return current->data;
}

// НОВАЯ ФУНКЦИЯ: Вставить элемент по индексу
void DoubleLinkedList::insertAt(size_t index, string value) {
    if(index > elementCount) {
        throw out_of_range("Index out of range");
    }
    
    if(index == 0) {
        pushFront(value);
        return;
    }
    
    if(index == elementCount) {
        pushBack(value);
        return;
    }
    
    DoubleNode* newNode = new DoubleNode(value);
    DoubleNode* current;
    
    // Выбираем оптимальный путь для поиска
    if(index < elementCount / 2) {
        current = head;
        for(size_t i = 0; i < index; i++) {
            current = current->next;
        }
    } else {
        current = tail;
        for(size_t i = elementCount - 1; i > index; i--) {
            current = current->prev;
        }
    }
    
    // Вставляем перед current
    newNode->prev = current->prev;
    newNode->next = current;
    current->prev->next = newNode;
    current->prev = newNode;
    
    elementCount++;
}

// НОВАЯ ФУНКЦИЯ: Удалить элемент по индексу
void DoubleLinkedList::removeAt(size_t index) {
    if(index >= elementCount) {
        throw out_of_range("Index out of range");
    }
    
    if(index == 0) {
        popFront();
        return;
    }
    
    if(index == elementCount - 1) {
        popBack();
        return;
    }
    
    DoubleNode* current;
    
    // Выбираем оптимальный путь для поиска
    if(index < elementCount / 2) {
        current = head;
        for(size_t i = 0; i < index; i++) {
            current = current->next;
        }
    } else {
        current = tail;
        for(size_t i = elementCount - 1; i > index; i--) {
            current = current->prev;
        }
    }
    
    // Удаляем current
    current->prev->next = current->next;
    current->next->prev = current->prev;
    delete current;
    elementCount--;
}

// НОВАЯ ФУНКЦИЯ: Заменить элемент по индексу
void DoubleLinkedList::replaceAt(size_t index, string value) {
    if(index >= elementCount) {
        throw out_of_range("Index out of range");
    }
    
    DoubleNode* current;
    
    if(index < elementCount / 2) {
        current = head;
        for(size_t i = 0; i < index; i++) {
            current = current->next;
        }
    } else {
        current = tail;
        for(size_t i = elementCount - 1; i > index; i--) {
            current = current->prev;
        }
    }
    
    current->data = value;
}

void DoubleLinkedList::clearDList() { 
    while(!isEmpty()) { 
        popFront();
    }
}

size_t DoubleLinkedList::size() const {
    return elementCount;
}

// НОВАЯ ФУНКЦИЯ: Обратный порядок
void DoubleLinkedList::reverse() {
    if(head == nullptr || head == tail) return;
    
    DoubleNode* current = head;
    DoubleNode* temp = nullptr;
    
    // Меняем head и tail
    temp = head;
    head = tail;
    tail = temp;
    
    // Меняем указатели у всех узлов
    current = head;
    while(current != nullptr) {
        temp = current->next;
        current->next = current->prev;
        current->prev = temp;
        current = temp;
    }
}
