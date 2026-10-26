class Heap {
  data: number[];
  idx: number;

  constructor() {
    this.data = [];
    this.idx = 0;
  }

  insert = (value: number) => {
    this.data.push(value);
    this.idx = this.data.length - 1;
    this.heapfyUp();
  };

  delete = () => {
    if (this.data.length === 0) {
      return;
    }
    const last = this.data.pop()!;

    if (this.data.length > 0) {
      this.data[0] = last;
      this.idx = 0;
      this.heapfyDown();
    }
  };

  private heapfyDown = () => {
    const n = this.data.length;
    const l = this.leftChild();
    const r = this.rightChild();
    let smallest = this.idx;
    if (l < n && this.data[l] < this.data[smallest]) {
      smallest = l;
    }
    if (r < n && this.data[r] < this.data[smallest]) {
      smallest = r;
    }
    if (smallest === this.idx) {
      return;
    }
    this.swap(smallest, this.idx);
    this.idx = smallest;
    this.heapfyDown();
  };

  private heapfyUp = () => {
    if (this.idx === 0) {
      return;
    }

    const p = this.parent();
    const curr = this.idx;

    if (this.data[p] <= this.data[curr]) {
      return;
    }
    this.swap(curr, p);
    this.idx = p;
    this.heapfyUp();
  };

  private parent = (): number => {
    return Math.floor((this.idx - 1) / 2);
  };

  private rightChild = (): number => {
    return this.idx * 2 + 2;
  };
  private leftChild = (): number => {
    return this.idx * 2 + 1;
  };
  private swap = (i: number, j: number) => {
    const temp = this.data[i];
    this.data[i] = this.data[j];
    this.data[j] = temp;
  };
}

export default Heap;
