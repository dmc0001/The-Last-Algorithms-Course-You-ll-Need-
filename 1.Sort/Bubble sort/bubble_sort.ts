const bubbleSort = (n: number[]) => {
  for (let i = 0; i < n.length; i++) {
    for (let j = 0; j < n.length - i - 1; j++) {
      if (n[j] > n[j + 1]) {
        let temp = n[j];
        n[j] = n[j + 1];
        n[j + 1] = temp;
      }
    }
  }
  return n;
};

let result = bubbleSort([1, 7, 6, 3, 100, 2, 10]);
console.log(result);
