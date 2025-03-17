// Package mixin switch
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
)

// Content  mixin 通常用于表示 内容	存放 html 之类的大文本信息
type Content struct {
	mixin.Schema
}

// Fields _
func (Content) Fields() []ent.Field {
	return []ent.Field{
		field.Text("content").
			Optional().
			Default("").
			Comment("内容").
			Annotations(ann.Field{
				PbIndex:   216,
				XSSFilter: true,
				Query: ann.FieldQuery{
					// 不允许查询 如需 like 则启用搜索
					Disabled: true,
				},
				// 查询不返回 只在详情中返回
				NotQueryReply:       true,
				NotPortalQueryReply: true,
			}),
	}
}

// Index _
func (Content) Index() []ent.Index {
	return []ent.Index{}
}
