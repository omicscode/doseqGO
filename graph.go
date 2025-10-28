package main

/*
 Gaurav Sablok
 codeprog@icloud.com
*/

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	ID         string
	Sequence   string
	Species    string
	StartPos   int
	StrandRank int
}

type Graph struct {
	Nodes map[string]*Node
	Edges map[string]map[string]bool
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
		Edges: make(map[string]map[string]bool),
	}
}

func (g *Graph) AddNode(id, seq, species string, start, sr int) {
	g.Nodes[id] = &Node{
		ID:         id,
		Sequence:   seq,
		Species:    species,
		StartPos:   start,
		StrandRank: sr,
	}
}

func (g *Graph) AddEdge(from, to string) {
	if _, ok := g.Edges[from]; !ok {
		g.Edges[from] = make(map[string]bool)
	}
	g.Edges[from][to] = true
}

func ParseGraph(scanner *bufio.Scanner) (*Graph, error) {
	g := NewGraph()

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 1 {
			continue
		}

		recordType := fields[0]

		switch recordType {
		case "S":
			if len(fields) < 5 {
				return nil, fmt.Errorf("invalid S line: %s", line)
			}
			id := fields[1]
			seq := fields[2]

			species := ""
			startPos := 0
			sr := 0

			for i := 3; i < len(fields); i++ {
				parts := strings.SplitN(fields[i], ":", 3)
				if len(parts) != 3 {
					continue
				}
				tag, typ, val := parts[0], parts[1], parts[2]
				switch tag {
				case "SN":
					if typ == "Z" {
						species = val
					}
				case "SO":
					if typ == "i" {
						if v, err := strconv.Atoi(val); err == nil {
							startPos = v
						}
					}
				case "SR":
					if typ == "i" {
						if v, err := strconv.Atoi(val); err == nil {
							sr = v
						}
					}
				}
			}

			g.AddNode(id, seq, species, startPos, sr)

		case "L":
			if len(fields) < 6 {
				return nil, fmt.Errorf("invalid L line: %s", line)
			}
			from := fields[1]
			fromOrient := fields[2]
			to := fields[3]
			toOrient := fields[4]

			var edgeFrom, edgeTo string
			if fromOrient == "+" {
				edgeFrom = from + "+"
			} else {
				edgeFrom = from + "-"
			}
			if toOrient == "+" {
				edgeTo = to + "+"
			} else {
				edgeTo = to + "-"
			}

			g.AddEdge(edgeFrom, edgeTo)
		}
	}

	return g, scanner.Err()
}

func (g *Graph) PrintGraph() {
	fmt.Println("=== NODES ===")
	for id, node := range g.Nodes {
		orient := "+"
		if node.StrandRank == 1 {
			orient = "-"
		}
		fmt.Printf("%s (%s) [%s] pos:%d len:%d\n", id, orient, node.Species, node.StartPos, len(node.Sequence))
	}

	fmt.Println("\n=== EDGES ===")
	for from, targets := range g.Edges {
		for to := range targets {
			fmt.Printf("%s -> %s\n", from, to)
		}
	}
}

func (g *Graph) DFS(start string, visited map[string]bool, path []string) [][]string {
	var results [][]string

	visited[start] = true
	path = append(path, start)

	if len(g.Edges[start]) == 0 {
		results = append(results, append([]string{}, path...))
	} else {
		for neighbor := range g.Edges[start] {
			if !visited[neighbor] {
				results = append(results, g.DFS(neighbor, visited, path)...)
			}
		}
	}

	visited[start] = false
	return results
}
