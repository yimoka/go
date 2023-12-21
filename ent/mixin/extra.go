// Package mixin switch
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
)

// Extra  mixin 通常用于表示 扩展信息
type Extra struct {
	mixin.Schema
}

// Fields _
func (Extra) Fields() []ent.Field {
	return []ent.Field{
		field.Text("extra").
			Optional().
			Default("").
			Comment("扩展信息").
			Annotations(ann.Field{
				PbIndex: 215,
			}),
	}
}

// Index _
func (Extra) Index() []ent.Index {
	return []ent.Index{}
}
