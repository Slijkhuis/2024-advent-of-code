package failedattempts

import (
	"os"
	"strings"
	"time"

	"github.com/Slijkhuis/2024-advent-of-code/pkg/aoc"
)

func Part1() {
	t := time.Now()
	g := aoc.NewGraph[string, struct{}]()

	for line := range aoc.LinesFromFile(os.Args[2]) {
		computers := strings.Split(line, "-")
		n1 := g.NewOrExistingNode(computers[0], struct{}{})
		n2 := g.NewOrExistingNode(computers[1], struct{}{})
		g.AddEdge(n1, n2)
		g.AddEdge(n2, n1)
	}

	aoc.Println(t, "built graph")

	cliques := g.BronKerbosch()
	aoc.Println(t, "found cliques")

	count := 0
	for _, clique := range cliques {
		if len(clique) < 3 {
			continue
		}

		aoc.Debug(strings.Join(aoc.Map(clique, func(n *aoc.Node[string, struct{}]) string { return n.Key }), ", "))

		numberOfTNodes := 0
		for _, node := range clique {
			if node.Key[0] == 't' {
				numberOfTNodes++
			}
		}
		if numberOfTNodes > 0 {
			if len(clique) == 3 {
				count++
			} else {
				l := len(clique)
				combinations := l * (l - 1) * (l - 2) / 6 // l! / (3! * (l - 3)!)

				// exclude combinations between computers that do not start with a t
				l2 := l - numberOfTNodes
				if l2 > 0 {
					combinations -= l2 * (l2 - 1) * (l2 - 2) / 6
				}

				aoc.Debug("l:", l, "l2:", l2, "combinations:", combinations)
				count += combinations
			}
		}
	}

	aoc.Println(t, count)
}
