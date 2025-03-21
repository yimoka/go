// Package ann Field 字段配置
package ann

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/samber/lo"
	"github.com/yimoka/go/utils"

	"entgo.io/ent/entc/gen"
)

// Field _
type Field struct {
	// PB 的序号
	PbIndex int

	// 表示单字段唯一索引（可忽略 sass 字段，减少 key 长度，并支持某些场景与 sass 字段无关的查询，例如通过域名/appid 反查配置）
	OnlyUnique bool

	// 特殊字段定义 之所以放在字段里定义是为了可以直接使用字段的 mixin
	// 是否是操作字段 一张表只能有一个操作字段 当 false 默认为 ID
	OperationField bool
	// 是否是 sass 字段 一张表只能有一个 sass 字段
	SassField bool
	// sass 只在 portal 服务使用
	SassOnlyPortal bool
	// 软删除字段 一张表只能有一个软删除字段 并且必须是 bool 类型
	SoftDeleteField bool
	// 行数据开关 停/启用的字段 一张表只能有一个行数据开关字段 并且必须是 bool 类型
	SwitchField bool
	// Portal 行数据开关的字段 只在 Portal 生成开关相关的方法
	SwitchPortalField bool

	//  用于生成 pb 的 validate 文档 https://github.com/bufbuild/protoc-gen-validate
	Validate string

	// 在查询的 Reply 不返回
	NotQueryReply bool
	// 在详情的 Reply 不返回
	NotDetailReply bool
	// 表示在 Portal 层 Add 时 不需要该字段，例如后台的 备注 字段
	NotPortalAdd bool
	// 表示在 Portal 层 Edit 时 不需要该字段，例如后台的 备注 字段
	NotPortalEdit bool
	// 表示在 Portal 层 Query 时 不需要该字段，例如后台的 备注 字段
	NotPortalQuery bool
	// 在 Portal 层的列表 Reply 不返回
	NotPortalQueryReply bool
	// 在 Portal 层的详情 Reply 不返回
	NotPortalDetailReply bool

	// 查询的配置
	Query FieldQuery
	// 仅是数据层的字段 如辅助加解密的字段
	OnlyData bool
	// 是否只从 meta 中获取, 如 sass 中的 appid
	OnlyMeta bool

	// 自动生成的值，不接受传参 如创建时间
	AutoCreate bool
	// 自动更新的值，不接受传参 如更新时间
	AutoUpdate bool

	// 默认值 用于 生成默认值的方法 ent 默认值不带 ctx，无法获取用户相关的数据
	DefaultFn *FieldFunc
	// 更新值 用于 生成更新值的方法 ent 默认值不带 ctx，无法获取用户相关的数据
	UpdateFn *FieldFunc

	// 当有值时 pb 使用 int64 并根据值是 秒还是毫秒在 service 层转换 time.Time
	PBTimeToType PBTimeType

	// 表示该字段存储需要加密, 仅在 data 层使用
	Encrypt bool
	// 掩码存储, 用于存储敏感数据, 仅在 data 层使用 如手机号
	// 请确保有 字段名 + Cipher 的字段用于存储加密后的数据 并确保 onlyData 为 true
	MaskEncrypt utils.MaskType
	// 独立加密, 用于存储敏感数据, 仅在 data 层使用 如 密码通过生成每一行独立的 nonce 保证安全
	// 请确保有 字段名 + Nonce 的字段用于存储随机生成的 nonce 并确保 onlyData 为 true
	// 请确保有 字段名 + Cipher 的字段用于存储加密后的数据 并确保 onlyData 为 true
	RowIrreversibleEncrypt bool

	// 是否需要 xss 过滤
	XSSFilter bool

	// 多语言字段，为空则不支持多语言，不为空则表示该字段为值字段的多语言值
	// 不为空则必须有对应的默认内容字段，必有为 json 类型 且值与值字段相同
	// 例如：值为 title 时，则必须有 title 字段
	// 本身字段名默认为 xxxI18n
	// 值为 json 类型，key 为语言代码，value 为对应的值
	// 例如：{"zh-CN":"标题","en-US":"title"}
	I18NFor string
	// Portal 层仅本地语言,为 true 在 Portal 层的无 I18N 输入，输出默认内容与和本地语言的内容
	PortalOnlyLocalLang bool

	// 索引
	// JSON 数组 CAST 索引 item 的长度 为 0 时不生成索引
	// ALTER TABLE `articles` ADD INDEX category_ids ((CAST(`category_ids` as char(15) ARRAY)))
	IndexJSONArrItemLen int

	// JSON 对象 keys CAST 索引 item 的长度 为 0 时不生成索引
	// ALTER TABLE `articles` ADD INDEX title_i18n ((CAST(JSON_KEYS(`title_i18n`) as char(7) ARRAY)))
	IndexJSONObjKeysLen int
}

// FnHandleType 处理方式
type FnHandleType string

const (
	// TError 返回错误
	TError FnHandleType = "error"
	// TIgnore 忽略
	TIgnore FnHandleType = "ignore"
	// TDefault 默认值
	TDefault FnHandleType = "default"
)

// FieldFunc 字段取值的方法
// 示例
//
//	 &ann.FieldFunc{
//		PkgPath: `"github.com/yimoka/go/middleware/meta"`,
//		Place:   ann.PlaceData,
//		GetStr:  "userID,_:=meta.GetUserID(ctx)",
//		SetStr:  `if b.Creator==nil && userID != "" {b.Creator = &userID}`, service 的写法 b.Field=xxx
//		SetStr: `if b.Creator==nil && userID != "" {db.SetCreator(userID)}`, data 层的写法 db.SetField(xxx)
//	}
type FieldFunc struct {
	PkgPath []string
	// 执行的方法
	GetStr string
	SetStr string
	// 执行的层  TODO biz 层待支持
	Place  Place
	Portal FnPortalType
}

// FnPortalType Portal 层的执行方式
type FnPortalType string

const (
	// FnPortalTypeDefault 默认 都执行
	FnPortalTypeDefault FnPortalType = ""
	// FnPortalTypeOnly 只在 Portal 层执行
	FnPortalTypeOnly FnPortalType = "only"
	// FnPortalTypeNot 不在 Portal 层执行
	FnPortalTypeNot FnPortalType = "not"
)

// PBTimeType pb 的时间类型
type PBTimeType string

const (
	// PBTimeTypeSecond 秒
	PBTimeTypeSecond PBTimeType = "second"
	// PBTimeTypeMillisecond 毫秒
	PBTimeTypeMillisecond PBTimeType = "milli"
)

// FieldQuery 字段的查询
type FieldQuery struct {
	// 默认支持 eq 不启用时全部关闭
	Disabled bool
	// 不等于
	NotEq bool
	// 是否开启包含
	In bool
	// 是否开启不包含
	NotIn bool
	// 是否开启模糊查询
	Like bool
	// 是否开启不模糊查询
	NotLike bool
	// 是否开启范围查询
	Range bool
}

// FieldNameKey 字段的注解名称
const FieldNameKey = "Field"

// Name of the annotation. Used by the codegen templates.
func (Field) Name() string {
	return FieldNameKey
}

// GetFieldConfig 获取字段的配置
func GetFieldConfig(node *gen.Field) *Field {
	ann := node.Annotations[FieldNameKey]
	if ann == nil {
		return &Field{}
	}
	data, err := json.Marshal(ann)
	if err != nil {
		log.Fatal(err)
	}
	var conf Field
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Fatal(err)
	}
	return &conf
}

// GetFieldsConfig 获取所有字段的配置
func GetFieldsConfig(node *gen.Type) map[string]*Field {
	fields := make(map[string]*Field)
	indexMap := make(map[int]bool)

	idConfig := GetFieldConfig(node.ID)
	if idConfig.PbIndex == 0 {
		idConfig.PbIndex = 1
	}
	indexMap[idConfig.PbIndex] = true
	fields[node.ID.Name] = idConfig

	for _, field := range node.Fields {
		config := GetFieldConfig(field)
		fields[field.Name] = config
		if config.PbIndex == 0 && !config.OnlyData {
			log.Fatalf("字段 %s 的 pbIndex 未配置", field.Name)
		}

		if config.PbIndex == 1 && !strings.EqualFold(field.Name, "id") {
			log.Fatalf("字段 %s 的 pbIndex 不能等于 1, 默认 1 预留给 ID", field.Name)
		}

		if indexMap[config.PbIndex] && config.PbIndex != 0 && !strings.EqualFold(field.Name, "id") {
			log.Fatalf("字段 %s 的 pbIndex 重复", field.Name)
		}
		indexMap[config.PbIndex] = true
		if config.OnlyUnique {
			_, iB := lo.Find(node.Indexes, func(item *gen.Index) bool {
				return len(item.Columns) == 1 && strings.EqualFold(item.Columns[0], field.Name)
			})
			if !iB {
				log.Fatalf("字段 %s 的 onlyUnique 为 true 时必须有单字段 Unique 索引", field.Name)
			}
		}
		if config.I18NFor != "" {
			// 必须是一个 json 字段
			if !field.IsJSON() {
				log.Fatalf("字段 %s 为多语言字段时时值必须为 json 类型", field.Name)
			}
			// 必须是一个 map 并且 key 为 string
			if !strings.HasPrefix(field.Type.String(), "map[string]") {
				log.Fatalf("字段 %s 为多语言字段时时值必须为 map[string]* 类型", field.Name)
			}
			cField := strings.TrimSuffix(field.Name, "I18n")
			if _, ok := fields[cField]; !ok {
				log.Fatalf("字段 %s 为多语言字段时必须有对应的值字段 %s", field.Name, cField)
			}
		}
	}
	return fields
}
