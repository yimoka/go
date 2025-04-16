// Package mixin provides appid mixin functionality
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
		field.String("appid").
			MaxLen(63).
			Comment("appid").
			Immutable().
			Annotations(ann.Field{
				PbIndex:   205,
				SassField: true,
				Validate:  "{max_len:63}",
			}),
	}
}

// Index _
func (Appid) Index() []ent.Index {
	return []ent.Index{
		index.Fields("appid"),
	}
}

// AppidOnlyPortal mixin 通常用于 sass 场景中标识不同的 app
type AppidOnlyPortal struct {
	mixin.Schema
}

// Fields _
func (AppidOnlyPortal) Fields() []ent.Field {
	return []ent.Field{
		field.String("appid").
			Comment("appid").
			Immutable().
			Annotations(ann.Field{
				PbIndex:        205,
				SassField:      true,
				SassOnlyPortal: true,
			}),
	}
}

// Index _
func (AppidOnlyPortal) Index() []ent.Index {
	return []ent.Index{
		index.Fields("appid"),
	}
}
