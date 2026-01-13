// internal/parser/ast.go
package parser

type NodeType string

const (
	NodeProgram  NodeType = "Program"
	NodeFunction NodeType = "Function"
	NodeLoop     NodeType = "Loop"
	NodeIf       NodeType = "If"
)

type ASTNode struct {
	Type     NodeType
	Children []*ASTNode
	Value    string
}

type AST struct {
	Root *ASTNode
}

func NewAST() *AST {
	return &AST{
		Root: &ASTNode{
			Type: NodeProgram,
		},
	}
}

func (a *AST) AddNode(parent *ASTNode, nodeType NodeType, value string) *ASTNode {
	node := &ASTNode{
		Type:  nodeType,
		Value: value,
	}
	parent.Children = append(parent.Children, node)
	return node
}
