from heap import Heap


def main():
  h = Heap()

  values = [5, 3, 8, 1, 9, 2, 7, 4, 6, 0]
  print("Inserting:", values)
  for v in values:
    h.insert(v)

  print("Internal array:", h.data)

  print("\nExtracting min one by one:")
  extracted = []
  while h.data:
    extracted.append(h.delete())
  print("Extracted order:", extracted)

  print("\nIs it sorted ascending? ", extracted == sorted(values))

  print("\nEdge cases:")
  print("  delete on empty heap ->", h.delete())

  single = Heap()
  single.insert(42)
  print("  single element delete ->", single.delete())
  print("  heap empty after?       ->", single.data == [])


if __name__ == "__main__":
  main()
