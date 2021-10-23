package main

// AbstractSyntaxTree has all nodes in it.
type AbstractSyntaxTree struct {
	node Node
}

// Node describes a basic node in the Abstract Syntax Tree
type Node interface {
}

// IntegerLiteralNode describes a basic int
type IntegerLiteralNode struct {
	value int
}

// FloatLiteralNode describes a basic float
type FloatLiteralNode struct {
	value float64
}

// ExpressionNode a basic expression
type ExpressionNode struct {
	operator string
	left     Node
	right    Node
}

// VariableAccessNode describes a variable, that is accessed by the user
type VariableAccessNode struct {
	name string
}

// DefineVariableNode describes a new variable created by the user
type DefineVariableNode struct {
	name  string
	value Node
}
