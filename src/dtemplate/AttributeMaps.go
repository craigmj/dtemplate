package dtemplate

type MapNode struct {
	Node *Node
	Pos  []int
}

func newMapNode(n *Node, p []int) *MapNode {
	pos := make([]int, len(p))
	copy(pos, p)
	return &MapNode{
		Node: n,
		Pos:  p,
	}
}

type AttributeMap map[string][]*MapNode

func NewAttributeMap() AttributeMap {
	return AttributeMap(map[string][]*MapNode{})
}

func MapNodes(attr string, removeAttr bool, node *Node) []*MapNode {
	nodes := []*MapNode{}
	// Check whether the parent node is a attributed node
	name := node.GetAttribute(attr)
	if `` != name {
		nodes = append(nodes, newMapNode(node, []int{}))
		// Remove the attribute after capturing the index
		if removeAttr {
			node.RemoveAttribute(attr)
		}
	}
	mapNodes_recurse(attr, removeAttr, node, []int{}, &nodes)
	return nodes
}

func mapNodes_recurse(attr string, removeAttr bool, parent *Node, path []int, nodes *[]*MapNode) {
	i := 0
	n := parent.FirstChild()
	lenp := len(path)
	cpath := make([]int, lenp+1)
	copy(cpath, path)
	for nil != n {
		cpath[lenp] = i
		if n.IsElement() {
			name := n.GetAttribute(attr)
			if `` != name {
				*nodes = append(*nodes, newMapNode(n, cpath))
				// Remove the attribute after capturing the index
				if removeAttr {
					n.RemoveAttribute(attr)
				}
			}
			// Only if this is an element, we recurse
			mapNodes_recurse(attr, removeAttr, n, cpath, nodes)
		}
		n = n.NextSibling()
		i++
	}
}
