package cmarkparser

import (
	"bytes"
	"fmt"
)

func NewInlineText(cont []byte) Node {
	var el Node

	el.Type = InlineText
	el.Content = cont

	return el
}

func NewParagraph(cont []byte) Node {
	var el Node

	el.Type = Par
	el.Content = cont

	return el
}

func NewThematicBreak(t byte) Node {
	var el Node

	el.Type = TBreak
	el.Content = []byte{t}

	return el
}

func NewHeading(level uint, content []byte) Node {
	var el Node
	var t NodeType

	switch level {
	case 1:
		t = H1
	case 2:
		t = H2
	case 3:
		t = H3
	case 4:
		t = H4
	case 5:
		t = H5
	case 6:
		t = H6
	}
	el.Type = t
	el.Content = content

	return el
}

func (n *Node) Empty() bool {
	return n.Type == None
}

func (d *Document) String() string {
	var buffer bytes.Buffer
	for _, c := range d.Children {
		buffer.WriteString(fmt.Sprintf("  %s\n", c.String()))
	}
	return buffer.String()
}

func (n *Node) String() string {
	var buffer bytes.Buffer
	if len(n.Content) > 0 {
		buffer.WriteString(fmt.Sprintf("[%s] %s", n.Type.String(), string(n.Content)))
	} else {
		buffer.WriteString(fmt.Sprintf("[%s]", n.Type.String()))
	}
	if len(n.Children) > 0 {
		buffer.WriteString("\n[")
	}
	for _, c := range n.Children {
		buffer.WriteString(fmt.Sprintf("  %s\n", c.String()))
	}
	if len(n.Children) > 0 {
		buffer.WriteString("]")
	}
	return buffer.String()
}

type NodeType uint8

const (
	None NodeType = iota
	Doc
	InlineText
	H1
	H2
	H3
	H4
	H5
	H6
	Par
	TBreak
)

var nodeTypeMap = map[string]NodeType{
	"nil": None,
	"doc": Doc,
	"txt": InlineText,
	"h1":  H1,
	"h2":  H2,
	"h3":  H3,
	"h4":  H4,
	"h5":  H5,
	"h6":  H6,
	"par": Par,
	"tbr": TBreak,
}

func getNodeType(s string) NodeType {
	return nodeTypeMap[s]
}

func (nt *NodeType) String() string {
	for key, node := range nodeTypeMap {
		if node == *nt {
			return key
		}
	}
	return "[nil]"
}

func (n *Nodes) String() string {
	var s string
	for k, v := range []Node(*n) {
		s += fmt.Sprintf("{%s}\n%s", k, v)
	}
	return s
}

type Document struct {
	Children Nodes
}

type Node struct {
	Type       NodeType
	Content    []byte
	Children   Nodes
	Attributes Attributes
}

type (
	Nodes      []Node
	Attributes map[string]string
)
