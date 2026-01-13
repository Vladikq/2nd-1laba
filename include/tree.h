#pragma once 
#include "includes.h"

struct NodeT { 
    int data;
    NodeT* left;
    NodeT* right;

    NodeT(int value);
};

struct FullBinaryTree { 
    NodeT* root;
    size_t size;

    FullBinaryTree();
    ~FullBinaryTree();

    void print();
    string toString();
    void insert(int value);
    void remove(int value); // НОВЫЙ МЕТОД
    bool search(int value); // НОВАЯ ПЕРЕГРУЗКА
    bool search(NodeT* node, int value);
    bool isFull();
    bool isFull(NodeT* node);
    bool isPerfect(); // НОВЫЙ МЕТОД
    string inOrder(); // НОВЫЙ МЕТОД
    void clear(); // НОВЫЙ МЕТОД
    size_t getSize() const; // НОВЫЙ МЕТОД
    
private:
    NodeT* _insert(NodeT* node, int value);
    NodeT* _remove(NodeT* node, int value);
    NodeT* _findMin(NodeT* node);
    string _toString(NodeT* node);
    string _inOrder(NodeT* node);
    void printTree(NodeT* node, int depth);
    void clear(NodeT* node);
    bool _isPerfect(NodeT* node, int height, int level);
    int getHeight(NodeT* node);
};
