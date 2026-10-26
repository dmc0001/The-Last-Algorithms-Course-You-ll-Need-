const twoCristalBalls = (n: boolean[]): number => {
  let jumpAmount = Math.floor(Math.sqrt(n.length));
  let i = jumpAmount;
  while (i <= n.length - 1) {
    if (n[i] === true) {
      break;
    }
    i += jumpAmount;
  }
  let j = i - jumpAmount;
  while (j <= n.length) {
    if (n[j] == true) {
      return j;
    }
    j++;
  }
  return -1;
};

const result = twoCristalBalls([false, false, false, false, true, true, true]);
console.log(result);
