// Package mixin switch
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
)

// Title  mixin 通常用于表示 标题
type Title struct {
	mixin.Schema
}

// Fields _
func (Title) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			MaxLen(511).
			Comment("标题").
			Annotations(ann.Field{
				PbIndex: 219,
				Query: ann.FieldQuery{
					Like:    true,
					NotLike: true,
				},
			}),
	}
}

// Index _
func (Title) Index() []ent.Index {
	return []ent.Index{}
}
