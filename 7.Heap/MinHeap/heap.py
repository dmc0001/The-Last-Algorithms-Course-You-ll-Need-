class Heap:
  def __init__(self):
    self.data = []
    self.idx = 0

  def insert(self, value):
    self.data.append(value)
    self.idx = len(self.data) - 1
    self.heapfy_up()

  def delete(self):
    n = len(self.data)
    if n == 0:
      return None

    root = self.data[0]
    last = self.data.pop()     

    if self.data:                
      self.data[0] = last
      self.idx = 0
      self.heapfy_down()

    return root

  def heapfy_down(self):
    n = len(self.data)
    l = self.left_child()
    r = self.right_child()
    smallest = self.idx

    if l < n and self.data[l] < self.data[smallest]:
      smallest = l
    if r < n and self.data[r] < self.data[smallest]:
      smallest = r

    if smallest == self.idx:
      return

    self.data[smallest], self.data[self.idx] = self.data[self.idx], self.data[smallest]
    self.idx = smallest
    self.heapfy_down()

  def heapfy_up(self):
    if self.idx == 0:
      return

    p = self.parent()
    curr = self.idx

    if self.data[p] <= self.data[curr]:
      return

    self.data[p], self.data[curr] = self.data[curr], self.data[p]
    self.idx = p
    self.heapfy_up()

  def parent(self):
    return (self.idx - 1) // 2

  def right_child(self):
    return (self.idx * 2) + 2

  def left_child(self):
    return (self.idx * 2) + 1
