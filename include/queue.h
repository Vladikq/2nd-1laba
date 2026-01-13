#pragma once 
#include "includes.h"

struct Queue { 
    string* data;
    size_t size;
    int front;     
    int rear;
    int capacity;  

    Queue(int cap);  
    Queue();
    ~Queue();

    void push(string value);
    bool isEmpty();
    bool isFull(); 
    string pop();   
    string peek();
    int Size();
    void clear(); 
    int Capacity(); 
};
