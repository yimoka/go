// Package ann Edge 边配置
package ann

import (
	"encoding/json"
	"log"

	"entgo.io/ent/entc/gen"
)

// Edge _
type Edge struct {
	// PB 的序号
	PbIndex int
}

const edgeKey = "edge"

// Name of the annotation. Used by the codegen templates.
func (Edge) Name() string {
	return edgeKey
}

// GetEdgeConfig 获取边的配置
func GetEdgeConfig(node *gen.Edge) *Edge {
	ann := node.Annotations[edgeKey]
	if ann == nil {
		return &Edge{}
	}
	data, err := json.Marshal(ann)
	if err != nil {
		log.Fatal(err)
	}
	var conf Edge
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Fatal(err)
	}
	return &conf
}

// GetEdgesConfig 获取所有边的配置
func GetEdgesConfig(node *gen.Type) map[string]*Edge {
	edges := make(map[string]*Edge)
	for _, edge := range node.Edges {
		edges[edge.Name] = GetEdgeConfig(edge)
	}
	return edges
}
