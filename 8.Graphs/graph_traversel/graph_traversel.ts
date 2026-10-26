const breathFirstSearch = (graph: Map<number, number[]>, source: number) => {
  let queue = [source];
  while (queue.length > 0) {
    const value = queue.shift()!;
    console.log(value);
    const neighbors = graph.get(value) || [];
    for (const neighbor of neighbors) {
      queue.push(neighbor);
    }
  }
};

const depthFirstSearch = (graph: Map<number, number[]>, source: number) => {
  console.log(source);
  const neighbors = graph.get(source) || [];
  for (const neighbor of neighbors) {
    depthFirstSearch(graph, neighbor);
  }
};

// Tree / DAG so the visited-free traversals terminate.
//        0
//      / | \
//     1  2  3
//     |     |
//     4     5
const graph = new Map<number, number[]>([
  [0, [1, 2, 3]],
  [1, [4]],
  [2, []],
  [3, [5]],
  [4, []],
  [5, []],
]);

console.log("BFS:");
breathFirstSearch(graph, 0);

console.log("DFS:");
depthFirstSearch(graph, 0);

export {};
