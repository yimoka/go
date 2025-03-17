// Package mixin mail
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
)

// Remark mixin
type Remark struct {
	mixin.Schema
}

// Fields _
func (Remark) Fields() []ent.Field {
	return []ent.Field{
		field.String("remark").
			Comment("备注").
			MaxLen(511).
			Optional().
			Default("").
			Annotations(
				ann.Field{
					PbIndex: 214,
					Query: ann.FieldQuery{
						Like: true,
					},
					NotPortalAdd:         true,
					NotPortalEdit:        true,
					NotPortalQuery:       true,
					NotPortalQueryReply:  true,
					NotPortalDetailReply: true,
				},
			),
	}
}
