// Package data operate.go
package data

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
		string(OpAdd):     "添加",
		string(OpDel):     "删除",
		string(OpSoftDel): "软删除",
		string(OpEdit):    "编辑",
		string(OpEnable):  "启用",
		string(OpDisable): "禁用",
		string(OpRecover): "恢复",
		string(OpUnknown): "未知",
	}
}

// String _
func (o OpType) String() string {
	return string(o)
}

// Label _
func (o OpType) Label() string {
	return OpTypeLabels()[string(o)]
}

// Values _
func (o OpType) Values() []string {
	return OpTypeValues()
}
