package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 800
	tileSize     = 20
)

var (
	started bool = false

	tileNumX int = screenWidth / tileSize
	tileNumY int = screenHeight / tileSize
)

func newGrid() [][]int8 {
	grid := make([][]int8, tileNumX)
	for i, _ := range grid {
		grid[i] = make([]int8, tileNumY)
	}

	return grid
}

func fillBoard(grid [][]int8) {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()

		col := mouseX / tileSize
		row := mouseY / tileSize
		grid[row][col] = 1
	} else if rl.IsKeyPressed(rl.KeySpace) {
		started = true
	}
}

func countNeighbors(grid [][]int8, x, y int) int8 {
	var sum int8

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			col := (x + i + tileNumX) % tileNumX
			row := (y + j + tileNumY) % tileNumY

			sum += grid[col][row]
		}
	}
	sum -= grid[x][y]

	return sum
}

func render(grid [][]int8) [][]int8 {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	prevGrid := grid
	var nextGrid [][]int8
	if !started {
		nextGrid = grid
	} else {
		nextGrid = newGrid()
	}

	var color rl.Color
	for row := 0; row < tileNumY; row++ {
		for col := 0; col < tileNumX; col++ {
			state := prevGrid[row][col]

			if state == 1 {
				color = rl.White
			} else {
				color = rl.Black
			}

			rl.DrawRectangle(
				int32(col*tileSize),
				int32(row*tileSize),
				tileSize,
				tileSize,
				color)

			if !started {
				continue
			}

			sum := countNeighbors(prevGrid, row, col)
			if state == 0 && sum == 3 {
				nextGrid[row][col] = 1
			} else if state == 1 && (sum < 2 || sum > 3) {
				nextGrid[row][col] = 0
			} else {
				nextGrid[row][col] = state
			}
		}
	}

	rl.EndDrawing()
	return nextGrid
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Game of Life")
	rl.SetTargetFPS(60)

	grid := newGrid()
	for !rl.WindowShouldClose() {
		if !started {
			fillBoard(grid)
		}
		grid = render(grid)
	}
}
