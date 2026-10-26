class Point {
  x: number;
  y: number;
  constructor(x: number, y: number) {
    this.x = x;
    this.y = y;
  }
}
const dir = [
  [0, 1],
  [0, -1],
  [1, 0],
  [-1, 0],
];

const walk = (
  maze: string[][],
  wall: string,
  curr: Point,
  end: Point,
  seen: boolean[][],
  path: Point[],
): boolean => {
  //Base case
  //a. check if we are out of the map
  if (curr.x > maze[0].length || curr.y > maze.length) {
    return false;
  }
  //b. check we are not in wall
  if (maze[curr.x][curr.y] === wall) {
    return false;
  }
  if (curr.x == end.x && curr.y == end.y) {
    path.push(curr);
    return true;
  }
  if (seen[curr.x][curr.y]) {
    return false;
  }
  seen[curr.x][curr.y] = true;
  path.push(curr);

  // recursion
  for (let i = 0; i < dir.length; i++) {
    const nextStep = dir[i];
    if (
      walk(
        (maze = maze),
        (wall = wall),
        (curr = new Point(curr.x + nextStep[0], curr.y + nextStep[1])),
        (end = end),
        (seen = seen),
        (path = path),
      )
    ) {
      return true;
    }
  }
  path.pop();
  return false;
};

const solve = (
  maze: string[][],
  wall: string,
  start: Point,
  end: Point,
): Point[] => {
  let seen: boolean[][] = [];
  let path: Point[] = [];
  for (let i = 0; i < maze.length; i++) {
    const rowLength = maze[i].length;
    seen.push(new Array(rowLength).fill(false));
  }

  walk(maze, wall, start, end, seen, path);
  return path;
};

const maze = [
  ["S", "W", ".", "."],
  [".", "W", "W", "."],
  [".", ".", "W", "."],
  ["W", ".", ".", "E"],
];

const startPoint = new Point(0, 0);
const endPoint = new Point(3, 3);
const wallPoint = "W";

console.log("Starting the maze solver...");

const resultPath = solve(maze, wallPoint, startPoint, endPoint);

console.log("\nResulting Path:");

if (resultPath && resultPath.length > 0) {
  for (const point of resultPath) {
    console.log(`(${point.x}, ${point.y})`);
  }
} else {
  console.log("No path found.");
}
