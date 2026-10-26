from collections import deque

def breath_first_search(graph,source):
  queue = deque()
  queue.append(source)
  while len(queue) != 0:
    value = queue.popleft()
    print(value)
    for neighbor in graph[value]:
      queue.append(neighbor)



def depth_first_search(graph,source):
  print(source)
  neighbors = graph[source]
  for neighbor in neighbors:
    depth_first_search(graph=graph,source=neighbor)


if __name__ == "__main__":
  # Tree / DAG so the visited-free traversals terminate.
  #        0
  #      / | \
  #     1  2  3
  #     |     |
  #     4     5
  graph = {
    0: [1, 2, 3],
    1: [4],
    2: [],
    3: [5],
    4: [],
    5: [],
  }

  print("BFS:")
  breath_first_search(graph, 0)

  print("DFS:")
  depth_first_search(graph, 0)
