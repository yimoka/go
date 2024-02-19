// Package mixin appid
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
)

// Appid mixin 通常用于 sass 场景中标识不同的 app
type Appid struct {
	mixin.Schema
}

// Fields _
func (Appid) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("appid").
			Min(1000).
			Comment("appid").
			Immutable().
			Annotations(ann.Field{
				PbIndex:   205,
				SassField: true,
			}),
	}
}

// Index _
func (Appid) Index() []ent.Index {
	return []ent.Index{
		index.Fields("appid"),
	}
}

// AppidOnlyBFF mixin 通常用于 sass 场景中标识不同的 app
type AppidOnlyBFF struct {
	mixin.Schema
}

// Fields _
func (AppidOnlyBFF) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("appid").
			Min(1000).
			Comment("appid").
			Immutable().
			Annotations(ann.Field{
				PbIndex:     205,
				SassField:   true,
				SassOnlyBFF: true,
			}),
	}
}

// Index _
func (AppidOnlyBFF) Index() []ent.Index {
	return []ent.Index{
		index.Fields("appid"),
	}
}
