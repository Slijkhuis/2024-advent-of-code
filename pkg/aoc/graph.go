package aoc

import (
	"container/heap"
	"math"
)

type Graph[K, T comparable] struct {
	Nodes map[K]*Node[K, T]
	Edges map[Edge[K, T]]struct{}
}

type StandardGraph Graph[Point, rune]

type Node[K, T comparable] struct {
	Key   K
	Value T
}

func NewNode[K, T comparable](key K, value T) *Node[K, T] {
	return &Node[K, T]{Key: key, Value: value}
}

type Edge[K, T comparable] struct {
	From   Node[K, T]
	To     Node[K, T]
	Weight int
}

type Path[K, T comparable] struct {
	Nodes  []*Node[K, T]
	Weight int
}

func NewGraph[K, T comparable]() *Graph[K, T] {
	return &Graph[K, T]{
		Nodes: map[K]*Node[K, T]{},
		Edges: map[Edge[K, T]]struct{}{},
	}
}

func NewStandardGraph() *Graph[Point, rune] {
	return &Graph[Point, rune]{
		Nodes: map[Point]*Node[Point, rune]{},
		Edges: map[Edge[Point, rune]]struct{}{},
	}
}

func (g *Graph[K, T]) NewOrExistingNode(key K, value T) *Node[K, T] {
	n, ok := g.Nodes[key]
	if !ok {
		n = NewNode(key, value)
		g.Nodes[key] = n
	}
	return n
}

func (g *Graph[K, T]) AddNode(n *Node[K, T]) {
	if existing, ok := g.Nodes[n.Key]; ok && existing != n {
		panic("Node already exists")
	}

	g.Nodes[n.Key] = n
}

func (g *Graph[K, T]) AddNodeIgnoreDuplicate(n *Node[K, T]) {
	g.Nodes[n.Key] = n
}

// AddEdge adds an edge to the graph. Returns true if the edge already exists.
func (g *Graph[K, T]) AddEdge(from, to *Node[K, T]) bool {
	if _, ok := g.Nodes[from.Key]; !ok {
		panic("Node (from) not found in graph")
	}

	if _, ok := g.Nodes[to.Key]; !ok {
		panic("Node (to) not found in graph")
	}

	edge := Edge[K, T]{From: *from, To: *to}
	_, ok := g.Edges[edge]
	if ok {
		return true
	}

	g.Edges[edge] = struct{}{}

	return false
}

// AddWeightedEdge adds a weighted edge to the graph. Returns true if the edge already exists.
func (g *Graph[K, T]) AddWeightedEdge(from, to *Node[K, T], weight int) bool {
	if _, ok := g.Nodes[from.Key]; !ok {
		panic("Node (from) not found in graph")
	}

	if _, ok := g.Nodes[to.Key]; !ok {
		panic("Node (to) not found in graph")
	}

	edge := Edge[K, T]{From: *from, To: *to, Weight: weight}
	_, ok := g.Edges[edge]
	if ok {
		return true
	}

	g.Edges[edge] = struct{}{}

	return false
}

// AreConnected (DFS)
func (g *Graph[K, T]) AreConnected(from, to *Node[K, T]) bool {
	visited := make(map[K]bool)
	return g.dfs(from, to, visited)
}

func (g *Graph[K, T]) dfs(current, target *Node[K, T], visited map[K]bool) bool {
	if current.Key == target.Key {
		return true
	}

	visited[current.Key] = true

	for edge := range g.Edges {
		if edge.From.Key == current.Key && !visited[edge.To.Key] {
			if g.dfs(g.Nodes[edge.To.Key], target, visited) {
				return true
			}
		}
	}

	return false
}

// FindAllPaths (BFS)
func (g *Graph[K, T]) FindAllPaths(from, to *Node[K, T]) []Path[K, T] {
	var paths []Path[K, T]
	queue := []Path[K, T]{{Nodes: []*Node[K, T]{from}, Weight: 0}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		last := path.Nodes[len(path.Nodes)-1]
		if last.Key == to.Key {
			paths = append(paths, path)
			continue
		}

		for edge := range g.Edges {
			if edge.From.Key == last.Key {
				newPath := Path[K, T]{
					Nodes:  append([]*Node[K, T]{}, path.Nodes...),
					Weight: path.Weight + edge.Weight,
				}
				newPath.Nodes = append(newPath.Nodes, g.Nodes[edge.To.Key])
				queue = append(queue, newPath)
			}
		}
	}

	return paths
}

func (g *Graph[K, T]) FindShortestPath(from, to *Node[K, T]) ([]*Node[K, T], int) {
	dist := make(map[K]int)
	prev := make(map[K]*Node[K, T])
	pq := make(PriorityQueue[K, T], 0)

	for key := range g.Nodes {
		dist[key] = math.MaxInt
		prev[key] = nil
	}

	dist[from.Key] = 0
	heap.Push(&pq, &Item[K, T]{node: from, priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item[K, T])
		u := item.node

		if u.Key == to.Key {
			var path []*Node[K, T]
			for u != nil {
				path = append([]*Node[K, T]{u}, path...)
				u = prev[u.Key]
			}
			return path, dist[to.Key]
		}

		for edge := range g.Edges {
			if edge.From.Key == u.Key {
				v := g.Nodes[edge.To.Key]
				alt := dist[u.Key] + edge.Weight
				if alt < dist[v.Key] {
					dist[v.Key] = alt
					prev[v.Key] = u
					heap.Push(&pq, &Item[K, T]{node: v, priority: alt})
				}
			}
		}
	}

	return nil, 0
}

func (g *Graph[K, T]) FindAllShortestPaths(from, to *Node[K, T]) ([][]*Node[K, T], int) {
	dist := make(map[K]int)
	prev := make(map[K][]*Node[K, T])
	pq := make(PriorityQueue[K, T], 0)

	for key := range g.Nodes {
		dist[key] = math.MaxInt
		prev[key] = []*Node[K, T]{}
	}

	dist[from.Key] = 0
	heap.Push(&pq, &Item[K, T]{node: from, priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item[K, T])
		u := item.node

		if item.priority > dist[u.Key] {
			continue
		}

		for edge := range g.Edges {
			if edge.From.Key == u.Key {
				v := g.Nodes[edge.To.Key]
				alt := dist[u.Key] + edge.Weight
				if alt < dist[v.Key] {
					dist[v.Key] = alt
					prev[v.Key] = []*Node[K, T]{u}
					heap.Push(&pq, &Item[K, T]{node: v, priority: alt})
				} else if alt == dist[v.Key] {
					prev[v.Key] = append(prev[v.Key], u)
				}
			}
		}
	}

	var paths [][]*Node[K, T]

	var backtrack func(*Node[K, T], []*Node[K, T])
	backtrack = func(n *Node[K, T], path []*Node[K, T]) {
		if n == from {
			paths = append(paths, Reverse(path))
			return
		}

		for _, prevNode := range prev[n.Key] {
			pathClone := append(path[:0:0], path...)
			backtrack(prevNode, append(pathClone, prevNode))
		}
	}

	backtrack(to, []*Node[K, T]{to})

	return paths, dist[to.Key]
}

func (g *Graph[K, T]) MustFindFirstNodeWithValue(value T) *Node[K, T] {
	for _, node := range g.Nodes {
		if node.Value == value {
			return node
		}
	}
	panic("Node not found")
}

// https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func (g *Graph[K, T]) BronKerbosch() [][]*Node[K, T] {
	R := make(map[K]*Node[K, T])
	X := make(map[K]*Node[K, T])
	P := CopyMap(g.Nodes)

	var cliques [][]*Node[K, T]
	g.bronKerbosch(R, P, X, &cliques)
	return cliques
}

func (g *Graph[K, T]) bronKerbosch(R, P, X map[K]*Node[K, T], cliques *[][]*Node[K, T]) {
	if len(P) == 0 && len(X) == 0 {
		*cliques = append(*cliques, Map(Keys(R), func(key K) *Node[K, T] { return R[key] }))
		return
	}

	localP := CopyMap(P)

	for _, node := range localP {
		newR := CopyMap(R)
		newR[node.Key] = node

		neighbors := g.Neighbors(node)

		newP := SliceToMap(Intersection(Values(P), neighbors), func(node *Node[K, T]) (K, *Node[K, T]) {
			return node.Key, node
		})
		newX := SliceToMap(Intersection(Values(X), neighbors), func(node *Node[K, T]) (K, *Node[K, T]) {
			return node.Key, node
		})

		g.bronKerbosch(newR, newP, newX, cliques)

		delete(P, node.Key)
		X[node.Key] = node
	}
}

func (g *Graph[K, T]) Neighbors(node *Node[K, T]) []*Node[K, T] {
	var neighbors []*Node[K, T]
	for edge := range g.Edges {
		if edge.From.Key == node.Key {
			if neighbor, exists := g.Nodes[edge.To.Key]; exists {
				neighbors = append(neighbors, neighbor)
			}
		}
	}
	return neighbors
}
