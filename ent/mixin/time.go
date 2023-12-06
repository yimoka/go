// Package mixin time
package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
)

// CreateTime mixin 创建时间 时间精度为秒,并且突破了 2038 年的限制
type CreateTime struct {
	mixin.Schema
}

// Fields _
func (CreateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createTime").
			Comment("创建时间").
			SchemaType(map[string]string{"mysql": "datetime"}).
			Default(time.Now).
			Immutable().
			Annotations(
				ann.Field{
					PbIndex:      201,
					AutoCreate:   true,
					PBTimeToType: ann.PBTimeTypeSecond,
					Query:        ann.FieldQuery{Range: true},
				},
			),
	}
}

// Index _
func (CreateTime) Index() []ent.Index {
	return []ent.Index{
		index.Fields("createTime"),
	}
}

// UpdateTime 更新时间 时间精度为秒,并且突破了 2038 年的限制
type UpdateTime struct {
	mixin.Schema
}

// Fields _
func (UpdateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("updateTime").
			Comment("更新时间").
			SchemaType(map[string]string{"mysql": "datetime"}).
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(
				ann.Field{
					PbIndex:      202,
					AutoCreate:   true,
					AutoUpdate:   true,
					PBTimeToType: ann.PBTimeTypeSecond,
					Query:        ann.FieldQuery{Range: true},
				},
			),
	}
}

// Index _
func (UpdateTime) Index() []ent.Index {
	return []ent.Index{
		index.Fields("updateTime"),
	}
}
