const factory = (n: number): number => {
  //Base Case
  if (n === 1) {
    return 1;
  }
  //We shall recurse
  return n * factory(n - 1);
};

console.log(factory(4));
