def buble_sort(n):
  i=0
  while i < len(n):
    j = 0
    while j < len(n) - i - 1:
      if n[j] > n[j+1]:
        n[j+1],n[j] = n[j], n[j+1]
      j+=1
    i+=1
  return n

result = buble_sort([1,7,6,3,100,2,10])
print(result)