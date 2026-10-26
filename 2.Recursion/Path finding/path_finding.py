class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y

# Directions: [x, y] offsets
dir = [
    [-1, 0], # Left
    [1, 0],  # Right
    [0, -1], # Up
    [0, 1],  # Down
]

def walk(maze, wall, curr, end, seen, path):
    # 1. Base cases:
    
    # a. Off the map (MUST BE FIRST to prevent index errors)
    if curr.x < 0 or curr.y < 0 or curr.y >= len(maze) or curr.x >= len(maze[0]):
        return False

    # b. It's a wall (Changed 'or' to 'and' so it only blocks the specific tile)
    if maze[curr.x][curr.y]== wall:
        return False

    # c. It's the end
    if curr.x == end.x and curr.y == end.y:
        path.append(end)
        return True
    
    # d. If we have seen it
    if seen[curr.y][curr.x]:
        return False
  
    # 2. Recursive steps:
    # Mark as seen so we don't step backward into an infinite loop
    seen[curr.y][curr.x] = True
    path.append(curr)

    # Fixed the range syntax here
    for i in range(len(dir)):
        newDir = dir[i]
        # Accessing list items by index instead of .x and .y
        next_step = Point(x=newDir[0] + curr.x, y=newDir[1] + curr.y)
        
        if walk(maze=maze, wall=wall, curr=next_step, end=end, seen=seen, path=path):
            return True
    
    # Backtrack if none of the directions panned out
    path.pop()
    return False


def solve(maze, wall, start, end):
    seen = []
    path = []
    # Fixed the range syntax here
    for i in range(len(maze)):
        row_length = len(maze[i])
        seen.append([False] * row_length)
        
    walk(maze, wall, start, end, seen, path)
    return path


# --- MAIN FUNCTION FOR TESTING ---
if __name__ == "__main__":
    maze = [
        ['S', 'W', '.', '.'],
        ['.', 'W', 'W', '.'],
        ['.', '.', 'W', '.'],
        ['W', '.', '.', 'E']
    ]
    
    
    start_point = Point(0, 0)
    end_point = Point(3, 3)
    wall_point = 'W'
    
    print("Starting the maze solver...")
    result_path = solve(maze, wall_point, start_point, end_point)
    
    print("\nResulting Path:")
    if result_path:
        for point in result_path:
            print(f"({point.x}, {point.y})")
    else:
        print("No path found.")


         