def binary_search(n, element):
  low = 0
  high = len(n)-1
  while low<= high:
    mid = low+(high-low)//2
    if (n[mid]== element):
      return True
    elif(n[mid]<element):
      low = mid+1
    else:
      high= mid
  return False

result = binary_search([1, 3, 67, 232, 240, 323, 391], 4)
print(result)