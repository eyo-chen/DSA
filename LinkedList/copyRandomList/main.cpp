#include <unordered_map>
using namespace std;

// Definition for a Node.
class Node {
 public:
  int val;
  Node* next;
  Node* random;

  Node(int _val) {
    val = _val;
    next = NULL;
    random = NULL;
  }
};

// First Solution: Using Hash Table
class Solution {
 public:
  Node* copyRandomList(Node* head) {
    unordered_map<Node*, Node*> hashTable;
    Node* old = head;

    while (old != nullptr) {
      Node* newNode = new Node(old->val);
      hashTable[old] = newNode;

      old = old->next;
    }

    old = head;
    while (old != nullptr) {
      hashTable[old]->next = hashTable[old->next];
      hashTable[old]->random = hashTable[old->random];

      old = old->next;
    }

    return hashTable[head];
  }
};

class Solution {
 public:
  Node* copyRandomList(Node* head) {
    if (head == nullptr) return head;

    Node* old = head;
    while (old != nullptr) {
      Node* newNode = new Node(old->val);

      newNode->next = old->next;
      old->next = newNode;

      old = old->next->next;
    }

    old = head;
    while (old != nullptr) {
      old->next->random = old->random ? old->random->next : nullptr;
      old = old->next->next;
    }

    old = head;
    Node* newHead = old->next;
    while (old != nullptr) {
      Node* newNode = old->next;

      old->next = newNode->next;
      newNode->next = newNode->next ? newNode->next->next : nullptr;

      old = old->next;
    }

    return newHead;
  }
};