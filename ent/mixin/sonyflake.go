// Package mixin implements the sonyflake ID generation for ent schema.
package mixin

import (
	"encoding/base64"
	"encoding/binary"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/sony/sonyflake"
	"github.com/yimoka/go/ent/ann"
	"github.com/yimoka/go/utils"
)

var (
	// defaultStartTime 是 Sonyflake 的起始时间：2023-01-01 00:00:00 UTC
	defaultStartTime = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	// globalSonyflake 是全局的 Sonyflake 实例
	globalSonyflake *sonyflake.Sonyflake
	once            sync.Once
)

// getSonyflake 返回单例的 Sonyflake 实例
func getSonyflake() *sonyflake.Sonyflake {
	once.Do(func() {
		globalSonyflake = sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime: defaultStartTime,
		})
	})
	return globalSonyflake
}

// SonyflakeID 实现了基于 Sonyflake 算法的 ID 生成器 mixin.
type SonyflakeID struct {
	mixin.Schema
}

// GenerateID 生成一个新的 ID。
// 该函数包含重试机制，在生成失败时最多重试 3 次。
// 如果所有重试都失败，将返回一个随机字符串作为降级策略。
func GenerateID() string {
	sf := getSonyflake()
	for i := 0; i < 3; i++ {
		if id, err := sf.NextID(); err == nil {
			buf := make([]byte, 8)
			binary.BigEndian.PutUint64(buf, id)
			return base64.RawURLEncoding.EncodeToString(buf)
		}
	}
	return utils.RandomStr(15)
}

// Fields 返回 SonyflakeID mixin 的字段定义。
func (SonyflakeID) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Comment("ID").
			Immutable().
			Unique().
			MaxLen(15).
			DefaultFunc(GenerateID).
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

// Indexes 返回 SonyflakeID mixin 的索引定义。
func (SonyflakeID) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}

// ParseID 解析 ID 字符串，返回生成时间、机器 ID 和序列号。
// 如果解析失败，将返回错误。
func ParseID(id string) (time.Time, uint16, uint16, error) {
	data, err := base64.RawURLEncoding.DecodeString(id)
	if err != nil {
		return time.Time{}, 0, 0, err
	}
	sfID := binary.BigEndian.Uint64(data)

	// Sonyflake 位分配:
	// +-----------------------------------------------------------------------------+
	// | 1 Bit Unused | 39 Bit Timestamp | 16 Bit Machine ID | 8 Bit Sequence ID |
	// +-----------------------------------------------------------------------------+

	timeOffset := sfID >> 24                  // 右移 24 位得到时间戳和未使用位
	timestamp := timeOffset & ((1 << 39) - 1) // 取低 39 位作为时间戳
	machineID := uint16((sfID >> 8) & 0xFFFF) // 取 16 位机器 ID
	sequence := uint16(sfID & 0xFF)           // 取 8 位序列号

	// 计算实际时间（Sonyflake 使用 10 毫秒作为时间单位）
	actualTime := defaultStartTime.Add(time.Duration(timestamp) * 10 * time.Millisecond)

	return actualTime, machineID, sequence, nil
}
