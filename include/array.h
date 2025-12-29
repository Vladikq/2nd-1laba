#pragma once
#include "includes.h"

struct Array
{ 
    string *arr;
    size_t size;
    size_t volume;

    Array();
    ~Array();
    
    void ShowArray() const;
    void addToEnd(string value);
    void add(size_t index, string value);
    string getIndex(size_t index);
    void remove(size_t index);
    void replace(size_t index, string value);
    size_t getSize() const;
};
