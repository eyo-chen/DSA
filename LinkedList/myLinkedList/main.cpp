class Node {
 public:
  int value;
  Node* next;
  Node(int val) : value(val), next(nullptr) {}
};

class MyLinkedList {
 public:
  MyLinkedList() {
    head = nullptr;
    size = 0;
  }

  int get(int index) {
    if (index < 0 || index >= size) return -1;

    Node* ptr = head;
    for (int i = 0; i < index; i++) {
      ptr = ptr->next;
    }

    return ptr->value;
  }

  void addAtHead(int val) {
    Node* newNode = new Node(val);
    newNode->next = head;
    head = newNode;
    size++;
  }

  void addAtTail(int val) {
    Node* newNode = new Node(val);
    size++;

    if (head == nullptr) {
      head = newNode;
      return;
    }

    Node* ptr = head;
    while (ptr->next != nullptr) {
      ptr = ptr->next;
    }

    ptr->next = newNode;
  }

  void addAtIndex(int index, int val) {
    // Note that the condition is index > size, not index >= size
    // when index == size, it means add to the tail
    if (index < 0 || index > size) return;

    if (index == 0) {
      addAtHead(val);
      return;
    }

    Node* newNode = new Node(val);
    Node* ptr = head;
    for (int i = 0; i < index - 1; i++) {
      ptr = ptr->next;
    }

    newNode->next = ptr->next;
    ptr->next = newNode;
    size++;
  }

  void deleteAtIndex(int index) {
    if (index < 0 || index >= size) return;

    Node* ptr = head;
    size--;

    if (index == 0) {
      head = head->next;
      delete ptr;
      return;
    }

    for (int i = 0; i < index - 1; i++) {
      ptr = ptr->next;
    }

    Node* deletedNode = ptr->next;
    ptr->next = deletedNode->next;
    delete deletedNode;
  }

 private:
  Node* head;
  int size;
};
