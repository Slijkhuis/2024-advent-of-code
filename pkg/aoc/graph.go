package aoc

type Graph[K, T comparable] struct {
	Nodes map[K]*Node[K, T]
	Edges map[Edge[K, T]]struct{}
}

type Node[K, T comparable] struct {
	Key   K
	Value T
}

func NewNode[K, T comparable](key K, value T) *Node[K, T] {
	return &Node[K, T]{Key: key, Value: value}
}

type Edge[K, T comparable] struct {
	From Node[K, T]
	To   Node[K, T]
}

func NewGraph[K, T comparable]() *Graph[K, T] {
	return &Graph[K, T]{
		Nodes: map[K]*Node[K, T]{},
		Edges: map[Edge[K, T]]struct{}{},
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
func (g *Graph[K, T]) FindAllPaths(from, to *Node[K, T]) [][]*Node[K, T] {
	var paths [][]*Node[K, T]
	queue := [][]*Node[K, T]{{from}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		last := path[len(path)-1]
		if last.Key == to.Key {
			paths = append(paths, path)
			continue
		}

		for edge := range g.Edges {
			if edge.From.Key == last.Key {
				newPath := append([]*Node[K, T]{}, path...)
				newPath = append(newPath, g.Nodes[edge.To.Key])
				queue = append(queue, newPath)
			}
		}
	}

	return paths
}
