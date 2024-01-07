// Package mixin phone
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
	"github.com/yimoka/go/utils"
)

// Phone mixin 手机号
type Phone struct {
	mixin.Schema
}

// Fields _
func (Phone) Fields() []ent.Field {
	return []ent.Field{
		field.String("phonePrefix").
			Comment("区号").
			MaxLen(7).
			Optional().
			Default("").
			Annotations(
				ann.Field{
					PbIndex: 207,
					Query: ann.FieldQuery{
						NotEq: true,
						In:    true,
						NotIn: true,
					},
				},
			),

		field.String("phone").
			Comment("手机号码").
			MaxLen(15).
			Optional().
			Annotations(
				ann.Field{
					PbIndex:     208,
					MaskEncrypt: utils.MaskTypePhone,
					Query: ann.FieldQuery{
						NotEq: true,
						In:    true,
						NotIn: true,
					},
				},
			),

		field.String("phoneCipher").
			MaxLen(255).
			Sensitive().
			Optional().
			Annotations(
				ann.Field{
					OnlyData: true,
				},
			),
	}
}

// Index _
func (Phone) Index() []ent.Index {
	return []ent.Index{
		index.Fields("phoneCipher"),
	}
}
