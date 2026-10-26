const binarySearch = (n: number[], element: number): boolean => {
  let low = 0;
  let mid = 0;
  let high = n.length - 1;

  while (low <= high) {
    mid = Math.floor(low + (high - low) / 2);
    if (n[mid] === element) {
      return true;
    } else if (n[mid] < element) {
      low = mid + 1;
    } else {
      high = mid;
    }
  }
  return false;
};

const result = binarySearch([1, 3, 67, 232, 240, 323, 391], 3);
console.log(result);
