const connectedComponentsCount = (graph: Map<number, number[]>): number => {
  let count = 0;
  let visited = new Map<number, boolean>();
  for (const [node] of graph) {
    if (explore(graph, visited, node)) {
      count = count + 1;
    }
  }
  return count;
};

const explore = (
  graph: Map<number, number[]>,
  visited: Map<number, boolean>,
  node: number,
): boolean => {
  if (visited.get(node)) {
    return false;
  }
  visited.set(node, true);
  for (const neighbor of graph.get(node) ?? []) {
    explore(graph, visited, neighbor);
  }
  return true;
};

export {};
