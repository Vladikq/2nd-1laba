#pragma once
#include "includes.h"

// ДОБАВЬ ЭТИ ОБЪЯВЛЕНИЯ:
struct Node {
    string data;
    Node* next;
    Node(string value);
};

struct DoubleNode { 
    string data;
    DoubleNode* next;
    DoubleNode* prev;
    DoubleNode(string value);
};

struct SinglyLinkedList { 
    Node* head;
    size_t elementCount;

    SinglyLinkedList();
    ~SinglyLinkedList();

    bool isEmpty() const;
    void print();
    void pushFront(string value);
    void pushBack(string value);
    void popFront();
    void popBack();
    void removeAt(string value);
    bool find(string value);
    void clearSList();
    Node* getHead() const;
};

struct DoubleLinkedList { 
    DoubleNode* head;
    DoubleNode* tail;
    size_t elementCount;

    DoubleLinkedList();
    ~DoubleLinkedList();

    bool isEmpty() const;
    void pushFront(string value);
    void pushBack(string value);
    void popFront();
    void popBack();
    void removeAt(string value);
    bool find(string value);
    void print();
    void clearDList();
    DoubleNode* getHead() const;  // УБЕДИСЬ ЧТО ЕСТЬ
};
