
def largest_component(graph):
  visited = set()
  largest = 0
  for node in graph:
    size = explore_size(graph=graph,visited=visited,curr=node)
    if  size > largest:
      largest = size

  return largest




def explore_size(graph,visited,curr):
  if curr in visited:
    return 0 
  visited.add(curr)
  size = 1
  neighbors = graph[curr]
  for neighbor in neighbors:
    size += explore_size(graph=graph,visited=visited,curr=neighbor)

  return size