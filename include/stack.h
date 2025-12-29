#pragma once 
#include "includes.h"

struct Stack { 
    string* data;
    int head;
    int capacity;  

    Stack(size_t size);
    Stack();
    ~Stack();

    void push(string value);
    string pop();
    string peek();
    bool isEmpty();
    size_t size();
};
