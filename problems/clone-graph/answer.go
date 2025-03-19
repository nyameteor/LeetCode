package main

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	nodeMap := make(map[*Node]*Node)
	newNode := dfsClone(node, nodeMap)
	return newNode
}

func dfsClone(node *Node, nodeMap map[*Node]*Node) *Node {
	if newNode, ok := nodeMap[node]; ok {
		return newNode
	}

	newNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}

	nodeMap[node] = newNode

	for i, neighbor := range node.Neighbors {
		newNode.Neighbors[i] = dfsClone(neighbor, nodeMap)
	}

	return newNode
}
