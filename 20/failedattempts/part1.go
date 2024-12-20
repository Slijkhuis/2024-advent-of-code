package failedattempts

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/Slijkhuis/2024-advent-of-code/pkg/aoc"
)

const cheatWeight = 1_000_000

func Part1TooSlow() {
	t := time.Now()

	// approach: make graph, if there's a wall and you're coming from a non-wall, add the node anyway,
	// but add a weight of `cheatWeight`. Coming from a wall, going to a another wall, should not have an edge at all.
	// From wall to non-wall is fine, also has a weight of 1.
	grid := aoc.BuildGridFromFile(os.Args[2])
	start, ok := grid.FindFirstCellWithValue('S')
	if !ok {
		aoc.Error("No start found")
	}
	end, ok := grid.FindFirstCellWithValue('E')
	if !ok {
		aoc.Error("No end found")
	}

	g := aoc.NewGraph[aoc.Point, rune]()

	visited := make(map[aoc.Point]bool)

	var buildGraph func(from *aoc.Node[aoc.Point, rune])
	buildGraph = func(from *aoc.Node[aoc.Point, rune]) {
		if _, ok := visited[from.Key]; ok {
			return
		}
		visited[from.Key] = true

		for _, dir := range aoc.NoDiagonals {
			to := from.Key.Move(dir)
			v, ok := grid.Data[to]
			if !ok {
				continue
			}

			if from.Value == '#' && v == '#' {
				continue
			}

			weight := 1
			if v == '#' {
				weight = cheatWeight
			}

			toNode := aoc.NewNode(to, v)
			g.AddNodeIgnoreDuplicate(toNode)
			g.AddWeightedEdge(from, toNode, weight)

			buildGraph(toNode)
		}
	}

	startNode := aoc.NewNode(start.Point, start.Value)
	g.AddNodeIgnoreDuplicate(startNode)
	buildGraph(startNode)
	aoc.Println(t, "built graph")

	endNode := aoc.NewNode(end.Point, end.Value)

	fastestPath, fastestWithoutCheats := g.FindShortestPath(startNode, endNode)
	aoc.Println(t, "found fastest time (no cheats)", fastestWithoutCheats)
	for _, node := range fastestPath {
		grid.Data[node.Key] = 'X'
	}
	aoc.Debug(grid.String())

	paths := findBetterPathsWithCheats(g, startNode, endNode, fastestWithoutCheats)
	aoc.Println(t, "found paths")

	cheatsPerSavings := map[int]int{}
	allSavings := []int{}
	cheatsThatSaveAtLeast100ps := 0
	for _, path := range paths {
		aoc.Debug(path.Weight, len(path.Nodes))

		trackTime := len(path.Nodes) - 1
		savings := fastestWithoutCheats - trackTime
		if savings <= 0 {
			continue
		}

		cheatsPerSavings[savings]++
		allSavings = append(allSavings, savings)

		if savings >= 100 {
			cheatsThatSaveAtLeast100ps++
		}
	}

	sort.IntSlice(allSavings).Sort()
	for _, savings := range aoc.Unique(allSavings) {
		aoc.Debug("There are", cheatsPerSavings[savings], "cheats that save", savings, "picoseconds.")
	}

	aoc.Println(t, "cheats that save at least 100 picoseconds:", cheatsThatSaveAtLeast100ps)
}

func findBetterPathsWithCheats(
	g *aoc.Graph[aoc.Point, rune],
	start, end *aoc.Node[aoc.Point, rune],
	fastestWithoutCheats int,
) []aoc.Path[aoc.Point, rune] {
	var allPaths []aoc.Path[aoc.Point, rune]
	var currentPath []*aoc.Node[aoc.Point, rune]
	visited := make(map[aoc.Point]bool)

	dfsInvocations := 0
	dfsSkips := 0

	dfs := func(current *aoc.Node[aoc.Point, rune], weight int) {}
	dfs = func(current *aoc.Node[aoc.Point, rune], weight int) {
		if dfsInvocations%1_000 == 0 {
			fmt.Println("dfs invocations:", dfsInvocations, "skips:", dfsSkips, "len(allPaths):", len(allPaths))
		}

		dfsInvocations++

		if aoc.DebugMode {
			if weight > fastestWithoutCheats+cheatWeight {
				dfsSkips++
				return
			}
		} else {
			if weight-99 > fastestWithoutCheats+cheatWeight {
				dfsSkips++
				return
			}
		}

		currentPath = append(currentPath, current)
		visited[current.Key] = true

		if current.Key == end.Key {
			pathCopy := make([]*aoc.Node[aoc.Point, rune], len(currentPath))
			copy(pathCopy, currentPath)
			allPaths = append(allPaths, aoc.Path[aoc.Point, rune]{Nodes: pathCopy, Weight: weight})
		} else {
			for edge := range g.Edges {
				if edge.From.Key == current.Key && !visited[edge.To.Key] {
					dfs(&edge.To, weight+edge.Weight)
				}
			}
		}

		currentPath = currentPath[:len(currentPath)-1]
		visited[current.Key] = false
	}

	dfs(start, 0)
	return allPaths
}
