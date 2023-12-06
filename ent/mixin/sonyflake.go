// Package mixin sonyflake
package mixin

import (
	"encoding/base64"
	"encoding/binary"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/sony/sonyflake"
	"github.com/yimoka/go/ent/ann"
	"github.com/yimoka/go/utils"
)

// SonyflakeID 使用 sonyflake 生产雪花 ID
type SonyflakeID struct {
	mixin.Schema
}

// Fields _
func (SonyflakeID) Fields() []ent.Field {
	s := sonyflake.NewSonyflake(sonyflake.Settings{})
	return []ent.Field{
		field.String("id").
			Comment("ID").
			Immutable().
			Unique().
			MaxLen(15).
			DefaultFunc(func() string {
				id, err := s.NextID()
				if err != nil {
					return utils.RandomStr(15)
				}
				buf := make([]byte, 8)
				binary.BigEndian.PutUint64(buf, id)
				return base64.StdEncoding.EncodeToString(buf)
			}).
			Annotations(ann.Field{
				PbIndex:    1,
				AutoCreate: true,
				Query: ann.FieldQuery{
					NotEq: true,
					In:    true,
					NotIn: true,
				},
			}),
	}
}

// Index _
func (SonyflakeID) Index() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}
