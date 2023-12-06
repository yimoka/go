// Package mixin secret
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
)

// Secret ent
type Secret struct {
	mixin.Schema
}

// Fields Secret
func (Secret) Fields() []ent.Field {
	return []ent.Field{
		field.String("secretID").
			Comment("密钥ID").
			MaxLen(127).
			Annotations(ann.Field{
				PbIndex: 211,
				Query: ann.FieldQuery{
					NotEq: true,
					In:    true,
					NotIn: true,
				},
			}),

		field.String("secretKey").
			Optional().
			MaxLen(255).
			Default("").
			Comment("密钥Key").
			Annotations(ann.Field{
				PbIndex:        212,
				Encrypt:        true,
				NotQueryReply:  true,
				NotDetailReply: true,
			},
			),
	}
}

// Indexes PhoneMxin
func (Secret) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("secretID"),
	}
}
