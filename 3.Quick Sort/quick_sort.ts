const partition = (n: number[], low: number, high: number): number => {
  const pivot = n[low];

  let i = low;
  let j = high;

  while (i < j) {
    while (i <= high && n[i] <= pivot) {
      i++;
    }

    while (j >= low && n[j] > pivot) {
      j--;
    }

    if (i < j) {
      swap(n, i, j);
    }
  }

  swap(n, low, j);

  return j;
};

const quickSort = (n: number[], low: number, high: number): void => {
  if (low < high) {
    const j = partition(n, low, high);

    quickSort(n, low, j - 1);
    quickSort(n, j + 1, high);
  }
};

const swap = (arr: number[], x: number, y: number): void => {
  const temp = arr[x];
  arr[x] = arr[y];
  arr[y] = temp;
};

const n = [7, 2, 9, 4, 1, 6, 3, 8, 5];

console.log("Before:", n);

quickSort(n, 0, n.length - 1);

console.log("After:", n);
