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
		field.Int32("tenantID").
			Min(1000).
			Comment("租户 ID").
			Immutable().
			Annotations(ann.Field{
				PbIndex:   205,
				SassField: true,
			}),
	}
}

// Index _
func (TenantID) Index() []ent.Index {
	return []ent.Index{
		index.Fields("tenantID"),
	}
}

// TenantIDOnlyBFF mixin 通常用于 sass 场景中标识不同的 租户
type TenantIDOnlyBFF struct {
	mixin.Schema
}

// Fields _
func (TenantIDOnlyBFF) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("tenantID").
			Min(1000).
			Comment("租户 ID").
			Immutable().
			Annotations(ann.Field{
				PbIndex:     205,
				SassField:   true,
				SassOnlyBFF: true,
			}),
	}
}

// Index _
func (TenantIDOnlyBFF) Index() []ent.Index {
	return []ent.Index{
		index.Fields("tenantID"),
	}
}
