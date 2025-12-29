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
    NodeT* _insert(NodeT* node, int value);
    bool search(NodeT* node, int value);
    bool isFull(NodeT* node); 
    bool isFull();
    string _toString(NodeT* node);
    void printTree(NodeT* node, int depth);
    void clear(NodeT* node);
};
