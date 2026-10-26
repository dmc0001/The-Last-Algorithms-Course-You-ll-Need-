class NodeTree {
  data: number;
  leftChild: NodeTree | null;
  rightChild: NodeTree | null;
  constructor(
    data: number,
    leftChild: NodeTree | null = null,
    rightChild: NodeTree | null = null,
  ) {
    this.data = data;
    this.leftChild = leftChild;
    this.rightChild = rightChild;
  }
}

// PRE-ORDER
const preOrederWalk = (tree: NodeTree | null, path: number[]) => {
  if (tree === null) return;
  path.push(tree.data);
  preOrederWalk(tree.leftChild, path);
  preOrederWalk(tree.rightChild, path);
};

const preOrder = (tree: NodeTree): number[] => {
  let path: number[] = [];
  preOrederWalk(tree, path);
  return path;
};

// IN-ORDER
const inOrederWalk = (tree: NodeTree | null, path: number[]) => {
  if (tree === null) return;
  inOrederWalk(tree.leftChild, path);
  path.push(tree.data);
  inOrederWalk(tree.rightChild, path);
};

const inOrder = (tree: NodeTree): number[] => {
  let path: number[] = [];
  inOrederWalk(tree, path);
  return path;
};

// POST-ORDER
const postOrederWalk = (tree: NodeTree | null, path: number[]) => {
  if (tree === null) return;
  postOrederWalk(tree.leftChild, path);
  postOrederWalk(tree.rightChild, path);
  path.push(tree.data);
};

const postOrder = (tree: NodeTree): number[] => {
  let path: number[] = [];
  postOrederWalk(tree, path);
  return path;
};

// # Create tree:
// #
// #         1
// #        / \
// #       2   3
// #      / \   \
// #     4   5   6
// #

let root: NodeTree = new NodeTree(
  1,
  new NodeTree(2, new NodeTree(4), new NodeTree(5)),
  new NodeTree(3, null, new NodeTree(6)),
);

let path = [];
let path2 = [];
let path3 = [];
path = preOrder(root);
path2 = inOrder(root);
path3 = postOrder(root);

console.log(path);
console.log(path2);
console.log(path3);
