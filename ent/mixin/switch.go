// Package mixin switch
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
)

// Switch  mixin 通常用于表示 停用/启用
type Switch struct {
	mixin.Schema
}

// Fields _
func (Switch) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("switch").
			Default(false).
			Comment("开关").
			Annotations(ann.Field{
				PbIndex:     213,
				SwitchField: true,
			}),
	}
}

// Index _
func (Switch) Index() []ent.Index {
	return []ent.Index{
		index.Fields("switch"),
	}
}
