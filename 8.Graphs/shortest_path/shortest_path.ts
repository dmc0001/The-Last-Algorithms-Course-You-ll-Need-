class Pair {
  first: number;
  second: number;
  constructor(first: number, second: number) {
    this.first = first;
    this.second = second;
  }
}
const buildGraph = (edges: Pair[]): Map<number, number[]> => {
  let graph: Map<number, number[]> = new Map();
  for (const edge of edges) {
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
const shortestPath = (edges: Pair[], src: number, dst: number): number => {
  let visited: Map<number, boolean> = new Map();
  const graph = buildGraph(edges);
  let q: [number, number][] = [[src, 0]];
  while (q.length > 0) {
    const [node, dist] = q.shift()!;
    if (dist === dst) {
      return dist;
    }
    const neighbors = graph.get(node) || [];
    for (const neighbor of neighbors) {
      if (!visited.has(neighbor)) {
        visited.set(neighbor, true);
        q.push([neighbor, dist + 1]);
      }
    }
  }
  return -1;
};
