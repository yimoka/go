// Package mixin password
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
)

// Password mixin 密码
type Password struct {
	mixin.Schema
}

// Fields _
func (Password) Fields() []ent.Field {
	return []ent.Field{
		field.String("password").
			MaxLen(255).
			Optional().
			Default("").
			Comment("密码").
			Annotations(
				ann.Field{
					PbIndex:                210,
					RowIrreversibleEncrypt: true,
					NotQueryReply:          true,
					NotDetailReply:         true,
					NotPortalQuery:         true,
					NotPortalQueryReply:    true,
					NotPortalDetailReply:   true,
					Query: ann.FieldQuery{
						Disabled: true,
					},
				},
			),

		field.String("passwordNonce").
			MaxLen(15).
			Sensitive().
			Optional().
			Default("").
			Annotations(
				ann.Field{
					OnlyData: true,
				},
			),
	}
}
