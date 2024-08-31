package main

import (
	maze "kevyn.fr/maze/src/Maze"
)

func main() {
	w := uint(16 * 5)
	h := uint(9 * 5)

	backtracking := &maze.GrowingTreeMaze{}
	backtracking.Init(w, h)
	backtracking.Generate(func(u []uint) int {
		return 0
	})
	backtracking.Draw("assets/backtracking.png")

	// prims := &maze.GrowingTreeMaze{}
	// prims.Init(w, h)
	// prims.Generate(func(u []uint) int {
	// 	return rand.Intn(len(u))
	// })
	// prims.Draw("assets/prims.png")

	// midGrowingTree := &maze.GrowingTreeMaze{}
	// midGrowingTree.Init(w, h)
	// midGrowingTree.Generate(func(u []uint) int {
	// 	return len(u) / 2
	// })
	// midGrowingTree.Draw("assets/mid.png")

	// oldGrowingTree := &maze.GrowingTreeMaze{}
	// oldGrowingTree.Init(w, h)
	// oldGrowingTree.Generate(func(u []uint) int {
	// 	return len(u) - 1
	// })
	// oldGrowingTree.Draw("assets/old.png")

	// old50new50GrowingTree := &maze.GrowingTreeMaze{}
	// old50new50GrowingTree.Init(w, h)
	// old50new50GrowingTree.Generate(func(u []uint) int {
	// 	if rand.Float32() < 0.5 {
	// 		return len(u) - 1
	// 	}

	// 	return 0
	// })
	// old50new50GrowingTree.Draw("assets/old-50.new-50.png")

	// rng50new50GrowingTree := &maze.GrowingTreeMaze{}
	// rng50new50GrowingTree.Init(w, h)
	// rng50new50GrowingTree.Generate(func(u []uint) int {
	// 	if rand.Float32() < 0.5 {
	// 		return rand.Intn(len(u))
	// 	}

	// 	return 0
	// })
	// rng50new50GrowingTree.Draw("assets/rng-50.new-50.png")

	// rng25new75GrowingTree := &maze.GrowingTreeMaze{}
	// rng25new75GrowingTree.Init(w, h)
	// rng25new75GrowingTree.Generate(func(u []uint) int {
	// 	if rand.Float32() < 0.25 {
	// 		return rand.Intn(len(u))
	// 	}

	// 	return 0
	// })
	// rng25new75GrowingTree.Draw("assets/rng-25.new-75.png")
}
