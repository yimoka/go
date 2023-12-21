// Package mixin switch
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
)

// Cover  mixin 通常用于表示 封面 图片地址
type Cover struct {
	mixin.Schema
}

// Fields _
func (Cover) Fields() []ent.Field {
	return []ent.Field{
		field.String("cover").
			MaxLen(2047).
			Optional().
			Default("").
			Comment("封面").
			Annotations(ann.Field{
				PbIndex: 218,
				Query: ann.FieldQuery{
					// 图片地址不允许查询
					Disabled: true,
				},
			}),
	}
}

// Index _
func (Cover) Index() []ent.Index {
	return []ent.Index{}
}
