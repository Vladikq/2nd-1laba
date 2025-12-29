#include "../include/array.h"

Array::Array() : volume(10), size(0) {              // Инициализация через список инициализации
    arr = new string[volume];                       // Выделение памяти начального объема
}

Array::~Array() {  
    delete[] arr;                                   // Освобождение динамического массива
}

void Array::ShowArray() const { 
    for(size_t i = 0; i < size; ++i) { 
        cout << arr[i] << endl;                     // Построчный вывод элементов
    }
    cout << endl;
}

void Array::addToEnd(string value) { 
    string *newArr = new string[size + 1];          // Создание нового массива на 1 элемент больше
    for(size_t i = 0; i < size; ++i) { 
        newArr[i] = arr[i];                         // Копирование существующих элементов
    }
    delete[] arr;                                   
    arr = newArr;                                   // Перенаправление указателя
    arr[size++] = value;                            // Добавление нового элемента и инкремент размера
}

void Array::add(size_t index, string value) { 
    if(index >= size) return;                       // Проверка корректности индекса

    string *newArr = new string[size + 1];          // Новый массив с увеличенным размером
    for(size_t i = 0; i < index; ++i) {  
        newArr[i] = arr[i];                         
    }
    newArr[index] = value;                          // Вставка нового значения
    for(size_t i = index; i < size; ++i) {  
        newArr[i + 1] = arr[i];                     // Копирование элементов после индекса со сдвигом
    }
    delete[] arr;
    arr = newArr;
    size++;                                         // Увеличение размера
}

string Array::getIndex(size_t index) { 
    if(index >= size) {
        throw out_of_range("Index out of range");   // Проверка выхода за границы
    }
    return arr[index];                              // Возврат элемента по индексу
}

void Array::remove(size_t index) { 
    if(index >= size) return;                       // Проверка валидности индекса

    string *newArr = new string[size - 1];          
    for(size_t i = 0; i < index; ++i) {
        newArr[i] = arr[i];                         // Копирование элементов до удаляемого
    }
    for(size_t i = index + 1; i < size; ++i) { 
        newArr[i - 1] = arr[i];                     // Копирование элементов после удаляемого со сдвигом
    }
    delete[] arr;
    arr = newArr;
    size--;                                         // Уменьшение размера
}

void Array::replace(size_t index, string value) { 
    if(index >= size) return;                       // Проверка границ
    arr[index] = value;                             // Замена значения по индексу
}

size_t Array::getSize() const {  
    return size;                                    // Возврат текущего размера
}
