class Node:

  def __init__(self, data, next=None):
    self.data = data
    self.next = next
    
class LinkedList:

  def __init__(self, head=None):
    self.head = head
  

  def insert(self,newValue):
    node = Node(data = newValue)

    if self.head == None:
      self.head = node
      return
    
    curr = self.head

    while curr.next:
      curr = curr.next

    curr.next = node
  
  def display(self):
    curr = self.head
    while curr:
      print(curr.data)
      curr = curr.next
    


if __name__ == "__main__":

  linkedList = LinkedList()
  linkedList.insert(1)
  linkedList.insert(12)
  linkedList.insert(123)
  linkedList.insert(1234)
  linkedList.insert(12345)

  curr = linkedList.head
  linkedList.display()



  


  

