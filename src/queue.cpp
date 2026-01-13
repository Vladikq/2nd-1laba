#include "../include/queue.h"

Queue::Queue(int cap) : capacity(cap), front(0), rear(-1), size(0) {
    if(capacity <= 0) {
        capacity = 30; // Защита от некорректного размера
    }
    data = new string[capacity];
}

Queue::Queue() : capacity(30), front(0), rear(-1), size(0) {
    data = new string[capacity];
}

Queue::~Queue() {
    delete[] data;
}

bool Queue::isEmpty() {
    return size == 0;
}

bool Queue::isFull() {
    return size == capacity;
}

void Queue::push(string value) {
    if (isFull()) {
        // Автоматическое расширение очереди
        int newCapacity = capacity * 2;
        string* newData = new string[newCapacity];
        
        // Копируем элементы в новый массив
        for(int i = 0; i < size; i++) {
            newData[i] = data[(front + i) % capacity];
        }
        
        delete[] data;
        data = newData;
        front = 0;
        rear = size - 1;
        capacity = newCapacity;
    }
    
    rear = (rear + 1) % capacity;
    data[rear] = value;
    size++;
}

string Queue::pop() {
    if (isEmpty()) {
        throw underflow_error("Queue is empty"); // Правильное исключение
    }
    string value = data[front];
    front = (front + 1) % capacity;
    size--;
    return value;
}

string Queue::peek() {
    if (isEmpty()) {
        throw underflow_error("Queue is empty");
    }
    return data[front];
}

int Queue::Size() {
    return size;
}

// НОВЫЙ МЕТОД: очистка очереди
void Queue::clear() {
    front = 0;
    rear = -1;
    size = 0;
}

// НОВЫЙ МЕТОД: получение емкости
int Queue::Capacity() {
    return capacity;
}
