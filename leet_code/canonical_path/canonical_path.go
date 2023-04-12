package canonical_path

import "strings"

type PathNode struct {
	Node   string
	Parent *PathNode
	Child  *PathNode
}

func (p *PathNode) append(component string) *PathNode {
	switch component {
	case "/":
		if p.Node != "/" {
			p.Child = &PathNode{Node: "/", Parent: p}
			return p.Child
		} else {
			return p
		}
	case ".":
		return p
	case "..":
		if p.Parent == nil { //this handles /../ noop
			return p
		}
		p.Parent.Parent.Child = nil
		return p.Parent.Parent //because we always have a node with a / at the end
	default:
		p.Child = &PathNode{Node: component, Parent: p}
		return p.Child
	}
}

func (p *PathNode) string() string {
	if p.Child == nil {
		if p.Node == "/" && p.Parent != nil { //remove trailing /
			return ""
		}
		return p.Node
	}
	return p.Node + p.Child.string()
}

func simplifyPath(path string) string {
	pathComponents := strings.Split(path, "/")
	pathLinkedList := &PathNode{Node: "/", Parent: nil}
	pathLinkedListHead := pathLinkedList

	for _, component := range pathComponents {
		if component == "" {
			continue
		}
		pathLinkedList = pathLinkedList.append("/")
		pathLinkedList = pathLinkedList.append(component)
	}

	return pathLinkedListHead.string()
}
