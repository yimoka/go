// Package mixin mail
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
	"github.com/yimoka/go/utils"
)

// Mail mixin 邮箱
type Mail struct {
	mixin.Schema
}

// Fields _
func (Mail) Fields() []ent.Field {
	return []ent.Field{
		field.String("mail").
			Comment("邮箱").
			MaxLen(63).
			Optional().
			Annotations(
				ann.Field{
					PbIndex:     209,
					Validate:    "{ max_len:63, email:true }",
					MaskEncrypt: utils.MaskTypeEmail,
					Query: ann.FieldQuery{
						NotEq: true,
						In:    true,
						NotIn: true,
					},
				},
			),

		field.String("mailCipher").
			MaxLen(255).
			Optional().
			Sensitive().
			Annotations(
				ann.Field{
					OnlyData: true,
				},
			),
	}
}

// Index _
func (Mail) Index() []ent.Index {
	return []ent.Index{
		index.Fields("mailCipher"),
	}
}
