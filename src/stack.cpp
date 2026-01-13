#include "../include/stack.h"

Stack::Stack(size_t size) : capacity(size), head(-1) {
    if(capacity == 0) capacity = 30;
    data = new string[capacity];
}

Stack::Stack() : capacity(30), head(-1) {
    data = new string[capacity];
}

Stack::~Stack() {
    delete[] data;
}

void Stack::push(string value) {
    if (head == capacity - 1) {
        // Автоматическое расширение
        size_t newCapacity = capacity * 2;
        string* newData = new string[newCapacity];
        
        for(size_t i = 0; i <= head; i++) {
            newData[i] = data[i];
        }
        
        delete[] data;
        data = newData;
        capacity = newCapacity;
    }
    data[++head] = value;
}

string Stack::pop() {
    if (head == -1) {
        throw underflow_error("Stack is empty");
    }
    return data[head--];
}

string Stack::peek() {
    if (head == -1) {
        throw underflow_error("Stack is empty");
    } 
    return data[head];
}

bool Stack::isEmpty() {
    return head == -1;
}

bool Stack::isFull() {
    return head == capacity - 1;
}

size_t Stack::size() {
    return head + 1;
}

// НОВЫЙ МЕТОД: очистка стека
void Stack::clear() {
    head = -1;
}

// НОВЫЙ МЕТОД: получение емкости
size_t Stack::getCapacity() {
    return capacity;
}
