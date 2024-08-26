package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	rl.EndDrawing()
}

func main() {
	rl.InitWindow(800, 450, "Game of Life")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		render()
	}
}
