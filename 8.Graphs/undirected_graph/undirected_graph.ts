class Pair {
  first: number;
  second: number;
  constructor(first: number, second: number) {
    this.first = first;
    this.second = second;
  }
}

const undirectPath = () => {};

const buildGraph = (edges: Pair[]): Map<number, number[]> => {
  let graph: Map<number, number[]> = new Map();
  for (let edge of edges) {
    const a = edge.first;
    const b = edge.second;

    if (!graph.has(a)) {
      graph.set(a, []);
    }
    graph.get(a)!.push(b);

    if (!graph.has(b)) {
      graph.set(b, []);
    }
    graph.get(b)!.push(a);
  }

  return graph;
};

const hasPathBFS = (
  graph: Map<number, number[]>,
  src: number,
  dst: number,
  seen: Map<number, boolean>,
): boolean => {
  let queue = [src];
  while (queue.length > 0) {
    const value = queue.shift()!;
    if (value === dst) {
      return true;
    }
    if (seen.get(value)) {
      continue;
    }
    seen.set(value, true);
    const neighbors = graph.get(value) || [];

    for (const neighbor of neighbors) {
      queue.push(neighbor);
    }
  }
  return false;
};

const hasPathDFS = (
  graph: Map<number, number[]>,
  src: number,
  dst: number,
  seen: Map<number, boolean>,
): boolean => {
  if (src === dst) {
    return true;
  }
  if (seen.get(src)) {
    return false;
  }
  seen.set(src, true);
  const neighbors = graph.get(src) || [];
  for (const neighbor of neighbors) {
    if (hasPathDFS(graph, neighbor, dst, seen)) {
      return true;
    }
  }
  return false;
};

const undirectGraph = (edges: Pair[], nodeA: number, nodeB: number) => {
  const graph = buildGraph(edges);
  const seen: Map<number, boolean> = new Map();
  return hasPathDFS(graph, nodeA, nodeB, seen);
};

// Undirected graph:
//   1 - 2 - 3
//   |       |
//   4       5      6 - 7  (separate component)
const edges = [
  new Pair(1, 2),
  new Pair(2, 3),
  new Pair(1, 4),
  new Pair(3, 5),
  new Pair(6, 7),
];

const g = buildGraph(edges);
console.log("adjacency list:", g);

const cases: [number, number, boolean][] = [
  [1, 5, true],  // 1->2->3->5
  [4, 5, true],  // 4->1->2->3->5
  [1, 6, false], // different component
  [6, 7, true],
  [2, 2, true],  // same node
];

for (const [src, dst, expected] of cases) {
  const bfs = hasPathBFS(g, src, dst, new Map());
  const dfs = hasPathDFS(g, src, dst, new Map());
  const status = bfs === expected && dfs === expected ? "ok" : "FAIL";
  console.log(`${src} -> ${dst}: BFS=${bfs} DFS=${dfs} (expected ${expected}) [${status}]`);
}

export {};
