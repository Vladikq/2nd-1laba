#pragma once 
#include "includes.h"

struct Stack { 
    string* data;
    int head;
    size_t capacity;  

    Stack(size_t size);
    Stack();
    ~Stack();

    void push(string value);
    string pop();
    string peek();
    bool isEmpty();
    bool isFull(); 
    size_t size();
    void clear(); 
    size_t getCapacity(); 
};
