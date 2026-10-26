const largestComponent = (graph: Map<number, number[]>): number => {
  let largest = 0;
  let visited = new Map<number, boolean>();
  for (const [node] of graph) {
    const size = exploreSize(graph, visited, node);
    if (size > largest) {
      largest = size;
    }
  }
  return largest;
};

const exploreSize = (
  graph: Map<number, number[]>,
  visited: Map<number, boolean>,
  curr: number,
): number => {
  if (visited.get(curr)) {
    return 0;
  }
  visited.set(curr, true);
  let size = 1;
  for (const neighbor of graph.get(curr) ?? []) {
    size += exploreSize(graph, visited, neighbor);
  }
  return size;
};

export {};
