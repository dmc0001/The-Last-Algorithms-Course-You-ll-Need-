from dataclasses import dataclass
from collections import deque


@dataclass
class Pair:
    first: int
    second: int


def shortest_path(edges, src, dst):
    graph = build_graph(edges)
    visited = {src}
    q = deque()
    q.append((src, 0))
    while q:
        node, dist = q.popleft()
        if node == dst:
            return dist
        for neighbor in graph.get(node, []):
            if neighbor not in visited:
                visited.add(neighbor)
                q.append((neighbor, dist + 1))
    return -1


def build_graph(edges):
    graph = {}
    for edge in edges:
        a, b = edge.first, edge.second
        graph.setdefault(a, []).append(b)
        graph.setdefault(b, []).append(a)
    return graph
