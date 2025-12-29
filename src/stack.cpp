#include "../include/stack.h"

Stack::Stack(size_t size) : capacity(size), head(-1) { 
    data = new string[capacity];                    // Выделение памяти под массив
}

Stack::Stack() : capacity(30), head(-1) { 
    data = new string[capacity];                    // Конструктор по умолчанию
}

Stack::~Stack() {
    delete[] data;                                  // Освобождение памяти
}

void Stack::push(string value) { 
    if (head == capacity - 1) { 
        throw overflow_error(" Stack is full");     // Проверка переполнения
    }
    data[++head] = value;                           // Инкремент + запись
}

string Stack::pop() { 
    if (head == -1) { 
        throw underflow_error(" Stack is empty");   // Проверка пустоты
    }
    return data[head--];                            // Возврат + декремент
}

string Stack::peek() { 
    if (head == -1) { 
        throw underflow_error(" Stack is empty");   // Проверка пустоты
    } 
    return data[head];                              // Чтение без изменения
}

bool Stack::isEmpty() {  
    return head == -1;                              // Пуст если head = -1
}

size_t Stack::size() { 
    return head + 1;                                // Размер = индекс + 1
}
