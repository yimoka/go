// Package mixin del
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
)

// Del mixin 软删除
type Del struct {
	mixin.Schema
}

// Fields _
func (Del) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("del").
			Default(false).
			Comment("软删除").
			Annotations(ann.Field{
				PbIndex:         206,
				SoftDeleteField: true,
				OnlyData:        true,
			}),
	}
}

// Index _
func (Del) Index() []ent.Index {
	return []ent.Index{
		index.Fields("del"),
	}
}
