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

func render(matrix [][]int8) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	if !started {
		fillBoard(matrix)
	}

	var color rl.Color
	for row := 0; row < tileNumY; row++ {
		for col := 0; col < tileNumX; col++ {
			if matrix[row][col] == 1 {
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
		}
	}

	rl.EndDrawing()
}

func fillBoard(matrix [][]int8) {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()

		col := mouseX / tileSize
		row := mouseY / tileSize
		matrix[row][col] = 1
	} else if rl.IsKeyPressed(rl.KeySpace) {
		started = true
	}
}
func main() {
	rl.InitWindow(screenWidth, screenHeight, "Game of Life")
	rl.SetTargetFPS(60)

	matrix := make([][]int8, tileNumX)
	for i, _ := range matrix {
		matrix[i] = make([]int8, tileNumY)
	}

	for !rl.WindowShouldClose() {
		render(matrix)
	}
}
