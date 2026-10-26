def partation(n,low,high):
  pivot = n[low]
  i = low
  j = high

  while i < j:
    while i <= high and n[i] <= pivot:
        i += 1

    while j >= low and n[j] > pivot:
        j -= 1
    if i < j:
       n[j], n[i] = n[i], n[j]

  n[j], n[low] = n[low], n[j]

  return j
    
def quick_sort(n,low,high):
  if low < high:
    j = partation(n= n, low=low, high=high)
    quick_sort(n= n,low= low, high= j - 1)
    quick_sort(n= n,low= j + 1 ,high= high)



# --- MAIN FUNCTION FOR TESTING ---
if __name__ == "__main__":
  numbers = [7, 2, 9, 4, 1, 6, 3, 8, 5]

  print("Before sorting:")
  print(numbers)

  quick_sort(
      n=numbers,
      low=0,
      high=len(numbers) - 1
  )

  print("\nAfter sorting:")
  print(numbers)