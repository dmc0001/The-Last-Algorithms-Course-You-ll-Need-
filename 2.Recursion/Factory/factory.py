def fatory(n):
  # Base case
  if (n == 1) :
    return 1
  # We shall recurse
  return fatory(n-1)*n

print(fatory(4))