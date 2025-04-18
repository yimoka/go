// Package mixin user
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
)

// CreatorByStaff  mixin 创建人
type CreatorByStaff struct {
	mixin.Schema
}

// Fields _
func (CreatorByStaff) Fields() []ent.Field {
	return []ent.Field{
		field.String("creatorByStaff").
			Comment("创建人").
			MaxLen(15).
			Optional().
			Immutable().
			Annotations(ann.Field{
				PbIndex:              220,
				Validate:             "{max_len:15}",
				AutoCreate:           true,
				NotPortalAdd:         true,
				NotPortalEdit:        true,
				NotPortalQuery:       true,
				NotPortalQueryReply:  true,
				NotPortalDetailReply: true,
				DefaultFn: &ann.FieldFunc{
					PkgPath: []string{"github.com/yimoka/go/middleware/meta"},
					Place:   ann.PlaceService,
					GetStr:  "staffID,_:=meta.GetStaffID(ctx)",
					SetStr:  `if b.CreatorByStaff==nil && staffID != "" {b.CreatorByStaff = &staffID}`,
					// 员工不在 Portal 层获取
					Portal: ann.FnPortalTypeNot,
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
func (CreatorByStaff) Index() []ent.Index {
	return []ent.Index{
		index.Fields("creatorByStaff"),
	}
}

// UpdaterByStaff 更新人 mixin
type UpdaterByStaff struct {
	mixin.Schema
}

// Fields _
func (UpdaterByStaff) Fields() []ent.Field {
	return []ent.Field{
		field.String("updaterByStaff").
			Comment("更新人").
			MaxLen(15).
			Optional().
			Annotations(ann.Field{
				PbIndex:              221,
				Validate:             "{max_len:15}",
				AutoCreate:           true,
				AutoUpdate:           true,
				NotPortalAdd:         true,
				NotPortalEdit:        true,
				NotPortalQuery:       true,
				NotPortalQueryReply:  true,
				NotPortalDetailReply: true,
				UpdateFn: &ann.FieldFunc{
					PkgPath: []string{"github.com/yimoka/go/middleware/meta"},
					Place:   ann.PlaceService,
					GetStr:  "staffID,_:=meta.GetStaffID(ctx)",
					SetStr:  `if b.UpdaterByStaff==nil && staffID != "" {b.UpdaterByStaff = &staffID}`,
					// 员工不在 Portal 层获取
					Portal: ann.FnPortalTypeNot,
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
func (UpdaterByStaff) Index() []ent.Index {
	return []ent.Index{
		index.Fields("updaterByStaff"),
	}
}
