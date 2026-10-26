const linearSearch = (n: number[], element: number): boolean => {
  for (let index = 0; index < n.length; index++) {
    if (element === n[index]) {
      return true;
    }
  }
  return false;
};

const result = linearSearch([1, 3, 67, 232, 0, 23, 1], 3);
console.log(result);
