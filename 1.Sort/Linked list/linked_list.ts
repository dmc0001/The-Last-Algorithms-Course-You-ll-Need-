class LNode {
  value: any;
  next: LNode | null = null;
  constructor(value: any, next: LNode | null = null) {
    this.next = next;
    this.value = value;
  }
}
class LinkedList {
  head: LNode | null = null;
  constructor(head: LNode | null = null) {
    this.head = head;
  }
  public insert(newValue: any): any {
    const node = new LNode(newValue);
    if (this.head == null) {
      this.head = node;
      return;
    }
    let curr = this.head;
    while (curr.next) {
      curr = curr.next;
    }
    curr.next = node;
  }
  public display() {
    let curr = this.head;
    while (curr !== null) {
      console.log(curr.value);
      curr = curr.next;
    }
  }
}

let list = new LinkedList();
list.insert(1);
list.insert(12);
list.insert(13);
list.insert(14);
list.display();
