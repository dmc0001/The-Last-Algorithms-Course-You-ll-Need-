package main

import "fmt"

type Point struct {
	x int
	y int
}

var dir = [][]int{
	{0, -1}, // Up    -> X: 0, Y: -1
	{1, 0},  // Right -> X: 1, Y: 0
	{-1, 0}, // Left  -> X: -1, Y: 0
	{0, 1},  // Down  -> X: 0, Y: 1
}

func walk(maze [][]string, wall string, curr Point, end Point, seen [][]bool, path *[]Point) bool {

	//Base case
	//1.check we are not out of the map
	if curr.x < 0 || curr.y < 0 || curr.y >= len(maze) || curr.x >= len(maze[0]) {
		return false
	}
	//2.check we are not in wall
	if maze[curr.y][curr.x] == wall {
		return false
	}
	//3. check we are not at the end
	if curr.x == end.x && curr.y == end.y {
		*path = append(*path, curr)
		return true
	}
	//4. check we have not seen this before
	if seen[curr.y][curr.x] {
		return false
	}

	seen[curr.y][curr.x] = true
	*path = append(*path, curr)

	for i := 0; i < len(dir); i++ {
		nextStep := dir[i]
		nextPoint := Point{x: curr.x + nextStep[0], y: curr.y + nextStep[1]}
		if walk(maze, wall, nextPoint, end, seen, path) {
			return true
		}
	}
	*path = (*path)[:len(*path)-1]
	return false
}

func solve(maze [][]string, wall string, start Point, end Point) []Point {
	path := []Point{}

	// Dynamically build the 'seen' 2D matrix matching the maze size
	seen := make([][]bool, len(maze))
	for i := 0; i < len(maze); i++ {
		seen[i] = make([]bool, len(maze[i]))
	}

	// Pass the memory address of path using '&'
	walk(maze, wall, start, end, seen, &path)
	return path
}

func main() {
	maze := [][]string{
		{"S", "W", ".", "."},
		{".", "W", "W", "."},
		{".", ".", "W", "."},
		{"W", ".", "W", "E"},
		{"W", ".", "W", "."},
		{"W", ".", ".", "."},
	}

	startPoint := Point{x: 0, y: 0}
	endPoint := Point{x: 3, y: 3}
	wallMarker := "W"

	fmt.Println("Starting the maze solver...")
	resultPath := solve(maze, wallMarker, startPoint, endPoint)

	fmt.Println("\nResulting Path:")
	if len(resultPath) > 0 {
		for _, point := range resultPath {
			fmt.Printf("(%d, %d)\n", point.x, point.y)
		}
	} else {
		fmt.Println("No path found.")
	}
}
