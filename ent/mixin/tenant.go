// Package mixin appid
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
)

// TenantID mixin 通常用于 sass 场景中标识不同的 租户
type TenantID struct {
	mixin.Schema
}

// Fields _
func (TenantID) Fields() []ent.Field {
	return []ent.Field{
		field.String("tenantID").
			MaxLen(63).
			Comment("租户 ID").
			Immutable().
			Annotations(ann.Field{
				PbIndex:   205,
				SassField: true,
				Validate:  "{max_len:63}",
			}),
	}
}

// Index _
func (TenantID) Index() []ent.Index {
	return []ent.Index{
		index.Fields("tenantID"),
	}
}

// TenantIDOnlyPortal mixin 通常用于 sass 场景中标识不同的 租户
type TenantIDOnlyPortal struct {
	mixin.Schema
}

// Fields _
func (TenantIDOnlyPortal) Fields() []ent.Field {
	return []ent.Field{
		field.String("tenantID").
			MaxLen(63).
			Comment("租户 ID").
			Immutable().
			Annotations(ann.Field{
				PbIndex:        205,
				SassField:      true,
				SassOnlyPortal: true,
				Validate:       "{max_len:63}",
			}),
	}
}

// Index _
func (TenantIDOnlyPortal) Index() []ent.Index {
	return []ent.Index{
		index.Fields("tenantID"),
	}
}
