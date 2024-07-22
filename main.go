package main

import maze "kevyn.fr/maze/src/Maze"

func main() {
	backtracking := &maze.BacktrackingMaze{}
	backtracking.Init(40, 40)
	backtracking.Generate()
	backtracking.Draw("backtracking.png")

	prims := &maze.PrimsMaze{}
	prims.Init(40, 40)
	prims.Generate()
	prims.Draw("prims.png")
}
