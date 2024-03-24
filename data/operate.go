/* cSpell:disable */
// Package data operate.go
package data

import "github.com/yimoka/go/lang"

// OpType 数据操作类型
type OpType string

const (
	// OpAdd 添加
	OpAdd OpType = "add"
	// OpDel 删除
	OpDel OpType = "del"
	// OpSoftDel 软删除
	OpSoftDel OpType = "softDel"
	// OpEdit 编辑/修改
	OpEdit OpType = "edit"
	// OpEnable 启用
	OpEnable OpType = "enable"
	// OpDisable 禁用
	OpDisable OpType = "disable"
	// OpRecover 恢复 从软删除恢复
	OpRecover OpType = "recover"
	// OpUnknown 未知
	OpUnknown OpType = "unknown"
)

// OpTypeValues _
func OpTypeValues() []string {
	return []string{
		string(OpAdd),
		string(OpDel),
		string(OpSoftDel),
		string(OpEdit),
		string(OpEnable),
		string(OpDisable),
		string(OpRecover),
		string(OpUnknown),
	}
}

// OpTypeLabels _
func OpTypeLabels() map[string]string {
	return map[string]string{
		string(OpAdd):     "add",
		string(OpDel):     "delete",
		string(OpSoftDel): "soft delete",
		string(OpEdit):    "edit",
		string(OpEnable):  "enable",
		string(OpDisable): "disable",
		string(OpRecover): "recover",
		string(OpUnknown): "unknown",
	}
}

// 多语言支持
var opTypeLangMap = map[string]map[string]string{
	"zh": {
		string(OpAdd):     "添加",
		string(OpDel):     "删除",
		string(OpSoftDel): "软删除",
		string(OpEdit):    "编辑",
		string(OpEnable):  "启用",
		string(OpDisable): "禁用",
		string(OpRecover): "恢复",
		string(OpUnknown): "未知",
	},
	"ru": {
		string(OpAdd):     "добавить",
		string(OpDel):     "удалить",
		string(OpSoftDel): "мягкое удаление",
		string(OpEdit):    "редактировать",
		string(OpEnable):  "включить",
		string(OpDisable): "запретить",
		string(OpRecover): "восстановить",
		string(OpUnknown): "неизвестный",
	},
}

// OpTypeLangLabels _
func OpTypeLangLabels(langs ...string) (map[string]string, bool) {
	return lang.MatchContent(opTypeLangMap, langs)
}

// String _
func (o OpType) String() string {
	return string(o)
}

// Label _
func (o OpType) Label() string {
	return OpTypeLabels()[string(o)]
}

// LangLabel _
func (o OpType) LangLabel(langs ...string) string {
	labels, ok := OpTypeLangLabels(langs...)
	if !ok {
		return o.Label()
	}
	return labels[string(o)]
}

// Values _
func (o OpType) Values() []string {
	return OpTypeValues()
}
