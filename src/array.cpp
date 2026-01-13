#include "../include/array.h"

Array::Array() : volume(10), size(0) {
    arr = new string[volume];
}

Array::~Array() {
    delete[] arr;
}

// ПРИВАТНЫЙ МЕТОД: расширение буфера
void Array::_resize(size_t newVolume) {
    string *newArr = new string[newVolume];
    for(size_t i = 0; i < size; ++i) {
        newArr[i] = arr[i];
    }
    delete[] arr;
    arr = newArr;
    volume = newVolume;
}

void Array::ShowArray() const {
    for(size_t i = 0; i < size; ++i) {
        cout << arr[i] << endl;
    }
    cout << endl;
}

void Array::addToEnd(string value) {
    if(size == volume) {
        _resize(volume * 2); // Увеличиваем в 2 раза
    }
    arr[size++] = value;
}

void Array::add(size_t index, string value) {
    if(index > size) return; // Допускаем index == size (добавление в конец)
    
    if(size == volume) {
        _resize(volume * 2);
    }
    
    // Сдвигаем элементы вправо
    for(size_t i = size; i > index; --i) {
        arr[i] = arr[i-1];
    }
    arr[index] = value;
    size++;
}

string Array::getIndex(size_t index) {
    if(index >= size) {
        throw out_of_range("Index out of range");
    }
    return arr[index];
}

void Array::remove(size_t index) {
    if(index >= size) return;
    
    // Сдвигаем элементы влево
    for(size_t i = index; i < size - 1; ++i) {
        arr[i] = arr[i+1];
    }
    size--;
    
    // Сжимаем массив если слишком пустой
    if(size < volume / 4 && volume > 10) {
        _resize(volume / 2);
    }
}

void Array::replace(size_t index, string value) {
    if(index >= size) return;
    arr[index] = value;
}

size_t Array::getSize() const {
    return size;
}

// НОВЫЙ МЕТОД: получение емкости
size_t Array::getVolume() const {
    return volume;
}

// НОВЫЙ МЕТОД: очистка массива
void Array::clear() {
    size = 0;
    delete[] arr;
    volume = 10;
    arr = new string[volume];
}
