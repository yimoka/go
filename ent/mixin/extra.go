// Package mixin switch
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
	"google.golang.org/protobuf/types/known/structpb"
)

// Extra  mixin 通常用于表示 扩展信息
type Extra struct {
	mixin.Schema
}

// Fields _
func (Extra) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("extra", &structpb.Struct{}).
			Optional().
			Default(&structpb.Struct{}).
			Comment("扩展信息").
			Annotations(ann.Field{
				PbIndex: 215,
				Query:   ann.FieldQuery{Disabled: true},
			}),
	}
}

// Index _
func (Extra) Index() []ent.Index {
	return []ent.Index{}
}
