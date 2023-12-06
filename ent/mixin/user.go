// Package mixin user
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
)

// Creator  mixin 创建人
type Creator struct {
	mixin.Schema
}

// Fields _
func (Creator) Fields() []ent.Field {
	return []ent.Field{
		field.String("creator").
			Comment("创建人").
			MaxLen(15).
			Optional().
			Immutable().
			Annotations(ann.Field{
				PbIndex:    203,
				AutoCreate: true,
				DefaultFn:  "meta.GetUserID(ctx)",
				Query: ann.FieldQuery{
					NotEq: true,
					In:    true,
					NotIn: true,
				},
			}),
	}
}

// Index _
func (Creator) Index() []ent.Index {
	return []ent.Index{
		index.Fields("creator"),
	}
}

// Updater 更新人 mixin
type Updater struct {
	mixin.Schema
}

// Fields _
func (Updater) Fields() []ent.Field {
	return []ent.Field{
		field.String("updater").
			Comment("更新人").
			MaxLen(15).
			Optional().
			Annotations(ann.Field{
				PbIndex:    204,
				AutoCreate: true,
				AutoUpdate: true,
				DefaultFn:  "meta.GetUserID(ctx)",
				UpdateFn:   "meta.GetUserID(ctx)",
				Query: ann.FieldQuery{
					NotEq: true,
					In:    true,
					NotIn: true,
				},
			}),
	}
}

// Index _
func (Updater) Index() []ent.Index {
	return []ent.Index{
		index.Fields("updater"),
	}
}
