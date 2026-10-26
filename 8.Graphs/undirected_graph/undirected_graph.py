from collections import deque
class Pair:
  def __init__(self,first,second):
    self.first = first
    self.second = second


def has_path_bfs(graph,src,dst,seen):
  q = deque()
  q.append(src)
  while len(q) > 0 :

    curr = q.popleft()
    if curr in seen:
      continue
    seen[curr] = True
    if curr == dst:
      return True
    neighbors = graph[curr]
    for neighbor in neighbors:
      q.append(neighbor)
  return False

def has_path_dfs(graph,src,dst,seen):
  if src == dst:
    return True
  if src in seen:
    return False
  seen[src] = True

  for neighbor in graph[src]:
    if has_path_dfs(graph, neighbor, dst, seen):
      return True
  return False

def build_graph(edges):
  graph ={}
  for edge in edges:
    a = edge.first
    b = edge.second
    graph.setdefault(a, []).append(b)
    graph.setdefault(b, []).append(a)

  return graph


def undirect_graph(edges,src,dst):
  graph = build_graph(edges)
  seen = {}
  return has_path_bfs(graph=graph,src=src,dst=dst,seen=seen)


if __name__ == "__main__":
  # Undirected graph:
  #   1 - 2 - 3
  #   |       |
  #   4       5      6 - 7  (separate component)
  edges = [
    Pair(1, 2),
    Pair(2, 3),
    Pair(1, 4),
    Pair(3, 5),
    Pair(6, 7),
  ]

  print("adjacency list:", build_graph(edges))

  cases = [
    (1, 5, True),   # 1->2->3->5
    (4, 5, True),   # 4->1->2->3->5
    (1, 6, False),  # different component
    (6, 7, True),
    (2, 2, True),   # same node
  ]
  graph = build_graph(edges)
  for src, dst, expected in cases:
    bfs = has_path_bfs(graph, src, dst, {})
    dfs = has_path_dfs(graph, src, dst, {})
    status = "ok" if bfs == expected and dfs == expected else "FAIL"
    print(f"{src} -> {dst}: BFS={bfs} DFS={dfs} (expected {expected}) [{status}]")