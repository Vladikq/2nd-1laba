#include "../include/queue.h"

Queue::Queue(int cap) : capacity(cap), front(0), rear(-1), size(0) { 
    data = new string[capacity];                    // Инициализация кольцевого буфера
}

Queue::Queue() : capacity(30), front(0), rear(-1), size(0) { 
    data = new string[capacity];                    // Конструктор по умолчанию
}

Queue::~Queue() { 
    delete[] data;                                  // Освобождение памяти
}

bool Queue::isEmpty() { 
    return size == 0;                               // Проверка по счетчику элементов
}

void Queue::push(string value) {  
    if (size == capacity) { 
        throw overflow_error(" Queue is full ");    // Проверка переполнения
    }
    rear = (rear + 1) % capacity;                   // Циклическое перемещение rear
    data[rear] = value;                             // Запись в конец очереди
    size++;                                         // Увеличение счетчика
}

string Queue::pop() { 
    if (size == 0) { 
        throw overflow_error(" Queue is empty");    // Проверка пустоты
    }
    string value = data[front];                     // Чтение из начала
    front = (front + 1) % capacity;                 // Циклическое перемещение front
    size--;                                         // Уменьшение счетчика
    return value;
} 

string Queue::peek() { 
    if (isEmpty()) { 
        throw underflow_error(" Queue is empty ");  // Проверка пустоты
    }
    return data[front];                             // Чтение без удаления
}

int Queue::Size() { 
    return size;                                    // Текущее количество элементов
}
