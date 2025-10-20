package wcodegen

import (
	"fmt"
	"log"
)

// Build a map[id]*TaskNode, flatten children recursively
func flattenNodes(nodes []*TaskNode, m map[string]*TaskNode) {
	for _, n := range nodes {
		if _, exists := m[n.Id]; !exists {
			m[n.Id] = n
		}
		if len(n.Children) > 0 {
			flattenNodes(n.Children, m)
		}
		if len(n.Dependencies) > 0 {
			flattenNodes(n.Dependencies, m)
		}
	}
}

// TopoSort using Kahn's algorithm
func topoSort(nodes []*TaskNode) ([]*TaskNode, error) {
	nodeMap := make(map[string]*TaskNode)
	flattenNodes(nodes, nodeMap)
	log.Println("已展开")
	// Build indegree map
	indegree := make(map[string]int)
	for id := range nodeMap {
		indegree[id] = 0
	}
	for _, n := range nodeMap {
		for _, dep := range n.Dependencies {

			if _, ok := indegree[dep.Id]; ok {
				indegree[n.Id]++
			} else {
				log.Printf("Warning: dependency %s not found for node %s\n", dep.Id, n.Id)
			}
		}
	}

	// Collect nodes with zero indegree
	queue := []*TaskNode{}
	for id, deg := range indegree {
		if deg == 0 {
			queue = append(queue, nodeMap[id])
		}
	}

	var sorted []*TaskNode
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		sorted = append(sorted, n)

		for _, other := range nodeMap {
			for _, dep := range other.Dependencies {
				if dep.Id == n.Id {
					indegree[other.Id]--
					if indegree[other.Id] == 0 {
						queue = append(queue, other)
					}
				}
			}
		}
	}

	// Check if there is a cycle
	if len(sorted) != len(nodeMap) {
		return nil, fmt.Errorf("cycle detected in DAG")
	}

	return sorted, nil
}
