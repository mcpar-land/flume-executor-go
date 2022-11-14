package flume

import "fmt"

type PortType struct {
	Type  string
	Name  string
	Label string
	Color string
}

type NodeType struct {
	Type         string
	Label        string
	Description  string
	InitialWidth string
	Root         bool
	Inputs       PortFn
	Outputs      PortFn
}

type Port struct {
	Type  string
	Name  string
	Label string
}

type PortFn func(any, any, any) []Port

type FlumeConfig struct {
	portTypes    map[string]PortType
	nodeTypes    map[string]NodeType
	rootNodeType string
}

func (cfg FlumeConfig) AddPortType(port PortType) FlumeConfig {
	cfg.portTypes[port.Type] = port
	return cfg
}

func (cfg FlumeConfig) AddNodeType(node NodeType) FlumeConfig {
	cfg.nodeTypes[node.Type] = node
	return cfg
}

func (cfg FlumeConfig) AddRootNodeType(node NodeType) FlumeConfig {
	cfg.nodeTypes[node.Type] = node
	cfg.rootNodeType = node.Type
	return cfg
}

type ResolvePortsFn func(any, any, any) any
type ResolveNodesFn func(any, any, any, any) any
type ResolveRootNode func(any, any) any

// https://github.com/chrisjpatty/flume/blob/master/src/RootEngine.js

type RootEngine struct {
	Config       *FlumeConfig
	ResolvePorts ResolvePortsFn
	ResolveNodes ResolveNodesFn
}

func (r *RootEngine) ResolveRootNode(nodes Map, options *RootResolveOptions) (any, error) {
	var opts RootResolveOptions
	if options != nil {
		opts = *options
	}
	rootNode, err := nodes.RootNode()
	if err != nil {
		return nil, err
	}
	inputs := r.Config.nodeTypes[rootNode.Type].Inputs(nil, nil, nil) // TODO
	return nil, fmt.Errorf("Not yet implemented")
}

type RootResolveOptions struct {
	Context              any
	OnlyResolveConnected bool
	RootNodeId           string
	MaxLoops             int
}

type Map struct {
	Nodes map[string]MapNode
	Edges map[string]MapEdge
}

func (m *Map) RootNode() (*MapNode, error) {
	var roots []MapNode
	for _, node := range m.Nodes {
		if node.Root {
			roots = append(roots, node)
		}
	}
	if len(roots) > 1 {
		return nil, fmt.Errorf("Root engine must not be called with more than 1 root node.")
	} else if len(roots) == 0 {
		return nil, fmt.Errorf("Root node not found")
	}
	return &roots[0], nil
}

type MapNode struct {
	Id   string
	Type string
	Root bool
}

type MapEdge struct {
}
