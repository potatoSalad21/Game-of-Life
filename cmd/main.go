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
	tileNumX = screenWidth / tileSize
	tileNumY = screenHeight / tileSize
)

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	for row := 0; row <= tileNumX; row++ {
		for col := 0; col <= tileNumY; col++ {
			//      -TODO-
			// store rectangles in a matrix
			// rectangle state: 0 (black == dead) or 1 (white == alive)
			rl.DrawRectangle(
				int32(col*tileSize),
				int32(row*tileSize),
				tileSize,
				tileSize,
				rl.White)
		}
	}

	rl.EndDrawing()
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Game of Life")
	rl.SetTargetFPS(60)

	// TODO: allow users to place blocks in initial state
	// press a button to begin
	for !rl.WindowShouldClose() {
		render()
	}
}
