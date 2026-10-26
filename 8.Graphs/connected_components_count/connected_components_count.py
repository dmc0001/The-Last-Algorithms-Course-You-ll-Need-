def connect_components_count(graph):
  visited = {} 
  count = 0
  for node in graph:
    if explore(graph=graph,visited=visited,node = node) :
      count +=1
  return count



def explore(graph,visited,node):
  if node in visited: 
    return False
  visited[node] = True
  neighbors = graph[node]

  for neighbor in neighbors:
     explore(graph=graph,visited=visited,node = neighbor)


  
