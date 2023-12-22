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
				DefaultFn: &ann.FieldFunc{
					PkgPath: []string{"github.com/yimoka/go/middleware/meta"},
					Place:   ann.PlaceService,
					GetStr:  "userID,_:=meta.GetUserID(ctx)",
					SetStr:  `if b.Creator==nil && userID != "" {b.Creator = &userID}`,
					// 用户 ID 只在 BFF 层获取
					BFF: ann.FnBFFTypeOnly,
				},
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
				UpdateFn: &ann.FieldFunc{
					PkgPath: []string{"github.com/yimoka/go/middleware/meta"},
					Place:   ann.PlaceService,
					GetStr:  "userID,_:=meta.GetUserID(ctx)",
					SetStr:  `if b.Updater==nil && userID != "" {b.Updater = &userID}`,
					BFF:     ann.FnBFFTypeOnly,
				},
				DefaultFn: &ann.FieldFunc{
					PkgPath: []string{"github.com/yimoka/go/middleware/meta"},
					Place:   ann.PlaceService,
					GetStr:  "userID,_:=meta.GetUserID(ctx)",
					SetStr:  `if b.Updater==nil && userID != "" {b.Updater = &userID}`,
					BFF:     ann.FnBFFTypeOnly,
				},
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
