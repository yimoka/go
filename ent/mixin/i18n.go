// Package mixin i18n 多语言字段
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/yimoka/go/ent/ann"
	"google.golang.org/protobuf/types/known/structpb"
)

type ExtraI18n struct {
	mixin.Schema
}

// Fields _
func (ExtraI18n) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("extraI18n", map[string]*structpb.Struct{}).
			Optional().
			Default(map[string]*structpb.Struct{}).
			Comment("国际化扩展信息").
			Annotations(ann.Field{
				PbIndex:          300,
				I18NFor:          "extra",
				BFFOnlyLocalLang: true,
				Query:            ann.FieldQuery{Disabled: true},
			}),
	}
}

// Content  mixin 通常用于表示 内容	存放 html 之类的大文本信息
type ContentI18n struct {
	mixin.Schema
}

// Fields _
func (ContentI18n) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("contentI18n", map[string]string{}).
			Optional().
			Default(make(map[string]string)).
			Comment("国际化内容").
			Annotations(ann.Field{
				PbIndex:          301,
				I18NFor:          "content",
				BFFOnlyLocalLang: true,
				XSSFilter:        true,
				Query:            ann.FieldQuery{Disabled: true},
				// 查询不返回 只在详情中返回
				NotQueryReply:    true,
				NotBffQueryReply: true,
			}),
	}
}

// SummaryI18n mixin 通常用于表示 简介
type SummaryI18n struct {
	mixin.Schema
}

// Fields _
func (SummaryI18n) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("summaryI18n", map[string]string{}).
			Optional().
			Default(make(map[string]string)).
			Comment("国际化简介").
			Annotations(ann.Field{
				PbIndex:          302,
				I18NFor:          "summary",
				BFFOnlyLocalLang: true,
				XSSFilter:        true,
				Query:            ann.FieldQuery{Disabled: true},
			}),
	}
}

// Cover  mixin 通常用于表示 封面 图片地址
type CoverI18n struct {
	mixin.Schema
}

// Fields _
func (CoverI18n) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("coverI18n", map[string]string{}).
			Optional().
			Default(make(map[string]string)).
			Comment("国际化封面").
			Annotations(ann.Field{
				PbIndex:          303,
				I18NFor:          "cover",
				BFFOnlyLocalLang: true,
				Query:            ann.FieldQuery{Disabled: true},
			}),
	}
}

// TitleI18n  mixin 通常用于表示 标题
type TitleI18n struct {
	mixin.Schema
}

// Fields _
func (TitleI18n) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("titleI18n", map[string]string{}).
			Optional().
			Default(make(map[string]string)).
			Comment("国际化标题").
			Annotations(ann.Field{
				PbIndex:             304,
				I18NFor:             "title",
				BFFOnlyLocalLang:    true,
				IndexJSONObjKeysLen: 7,
				Query: ann.FieldQuery{
					Like:    true,
					NotLike: true,
				},
			}),
	}
}

// 副标题国际化
type SubTitleI18n struct {
	mixin.Schema
}

// Fields _
func (SubTitleI18n) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("subTitleI18n", map[string]string{}).
			Optional().
			Default(make(map[string]string)).
			Comment("国际化副标题").
			Annotations(ann.Field{
				PbIndex:          305,
				I18NFor:          "subTitle",
				BFFOnlyLocalLang: true,
				Query:            ann.FieldQuery{Disabled: true},
			}),
	}
}
