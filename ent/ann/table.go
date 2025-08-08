// Package ann Table
package ann

import (
	"encoding/json"
	"log"

	"entgo.io/ent/entc/gen"
	"github.com/yimoka/api/common"
)

// Table 数据库表的注解
type Table struct {
	// 表名 用于展示
	TableName string `json:"tableName"`

	// 默认每页数量 默认 20
	DefaultPageSize int `json:"defaultPageSize"`
	// 最大每页数量 默认 1000
	MaxPageSize int `json:"maxPageSize"`
	// 默认排序
	DefaultSortOrder []*common.SortOrder `json:"defaultSort"`

	// 是否启用返回所有记录 默认不启用
	EnableQueryAll bool `json:"enableQueryAll"`
	// 缓存 缓存以单表在 data 层进行存储，边数据，为防边循环和与data数据一致，只取一层边关系数据
	Cache TableCache `json:"cache"`
	// 聚合与分组
	Aggregates []Aggregate `json:"aggregate"`
	// 树的父级字段名 为空则不启用 值为字段名
	TreeParent string `json:"treeParent"`
	// 边
	WithEdge map[string]WithConfig `json:"withEdge"`
	// json 字段的 ids 带上对应表的数据
	WithTableByJSON map[string]WithTableByJSON `json:"withTableByJSON"`
	// 服务
	WithServices []WithService `json:"withServices"`
	// 自定义方法配置
	Custom TableCustom `json:"custom"`
	// 是否生成 Http 接口 默认为空不生成 仅提供 rpc 如有值则为接口前缀
	HTTPAPI string `json:"httpApi"`
	// Portal 提供给前端的接口, 通过是部分用于管理的字段不返回给前端
	PortalAPI TableFn `json:"portalApi"`
	// 加密的密钥 Key 如不填则使用默认
	SecretKey string `json:"secretKey"`

	OpProtectConfig OpProtectConfig `json:"operationProtectionConfig"`

	MutationConfig MutationConfig `json:"mutationConfig"`

	// 是否是日志操作表
	IsOpLogTable bool `json:"isOpLogTable"`

	// 权限配置的ID
	PermissionsID string `yaml:"permissionsParentID"`

	// 自定义字段 业务逻辑字段 不存在于数据库 仅用于代码生成
	CustomFields []CustomField `json:"customFields"`

	// 切面编程 用于在代码生成时 对代码进行切面处理
	// 支持在  service biz 层对自动生成的每个 fn 进行处理
	ServiceBefore       TableFn `json:"serviceBefore"`
	ServiceAfter        TableFn `json:"serviceAfter"`
	ServicePortalBefore TableFn `json:"servicePortalBefore"`
	ServicePortalAfter  TableFn `json:"servicePortalAfter"`
	BizBefore           TableFn `json:"bizBefore"`
	BizAfter            TableFn `json:"bizAfter"`
}

// Place 代码生成的层
type Place string

const (
	// PlaceProto proto 层
	PlaceProto Place = "proto"
	// PlaceService service 层
	PlaceService Place = "service"
	// PlaceBiz biz 层
	PlaceBiz Place = "biz"
	// PlaceData data 层
	PlaceData Place = "data"
)

const tableNameKey = "Table"

// Name _
func (Table) Name() string {
	return tableNameKey
}

// GetTableConfig 获取表配置
func GetTableConfig(node *gen.Type) *Table {
	ann := node.Annotations[tableNameKey]
	if ann == nil {
		return &Table{}
	}
	data, err := json.Marshal(ann)
	if err != nil {
		log.Fatal(err)
	}
	var conf Table
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Fatal(err)
	}
	if conf.DefaultPageSize == 0 {
		conf.DefaultPageSize = 20
	}
	if conf.MaxPageSize == 0 {
		conf.MaxPageSize = 1000
	}
	return &conf
}

// TableCache 表缓存配置
type TableCache struct {
	// 详情缓存 支持配置多个字段 请确保字段的值唯一
	Detail          []string                `json:"detail"`
	DetailBizBefore map[string]TableCacheFn `json:"detailBizBefore"`
	DetailBizAfter  map[string]TableCacheFn `json:"detailBizAfter"`

	DetailBefore       map[string]TableCacheFn `json:"detailBefore"`
	DetailAfter        map[string]TableCacheFn `json:"detailAfter"`
	PortalDetailFn     map[string]TableCacheFn `json:"portalDetailFn"`
	PortalDetailBefore map[string]TableCacheFn `json:"portalDetailBefore"`
	PortalDetailAfter  map[string]TableCacheFn `json:"portalDetailAfter"`

	// 启用查询所有缓存
	QueryAll          bool         `json:"queryAll"`
	QueryAllBizBefore TableCacheFn `json:"queryAllBizBefore"`
	QueryAllBizAfter  TableCacheFn `json:"queryAllBizAfter"`

	QueryAllBefore TableCacheFn `json:"queryAllBefore"`
	QueryAllAfter  TableCacheFn `json:"queryAllAfter"`
	// Portal层查询所有缓存 启用的方法 all 只有三个方法 get set del
	PortalQueryAllFn     TableCacheFn `json:"portalQueryAll"`
	PortalQueryAllBefore TableCacheFn `json:"portalQueryAllBefore"`
	PortalQueryAllAfter  TableCacheFn `json:"portalQueryAllAfter"`

	// 查询缓存 支持配置多个字段
	Query          []string                `json:"query"`
	QueryBizBefore map[string]TableCacheFn `json:"queryBizBefore"`
	QueryBizAfter  map[string]TableCacheFn `json:"queryBizAfter"`

	QueryBefore map[string]TableCacheFn `json:"queryBefore"`
	QueryAfter  map[string]TableCacheFn `json:"queryAfter"`
	// Portal层查询缓存 启用的方法
	PortalQueryFn     map[string]TableCacheFn `json:"portalQueryFn"`
	PortalQueryBefore map[string]TableCacheFn `json:"portalQueryBefore"`
	PortalQueryAfter  map[string]TableCacheFn `json:"portalQueryAfter"`

	// 是否不开启穿透保护
	// 穿透保护查询不到的数据 会缓存一个空值
	// 详情缓存 返回 404 列表缓存 返回 空数组
	NotPreventPenetration bool `json:"notPreventPenetration"`
}

// TableCacheFn _
type TableCacheFn struct {
	Get  bool
	MGet bool
	Set  bool
	MSet bool
	Del  bool
}

// WithConfig _
type WithConfig struct {
	Detail bool
	Rows   bool
	Add    bool
	Edit   bool
	Query  bool
}

// WithTableByJSON json 字段的 ids 带上对应表的数据
type WithTableByJSON struct {
	Table  string
	Detail bool
	Rows   bool
}

// TableCustom _
type TableCustom struct {
	Service TableFn
	Biz     TableFn
}

// TableFn _
type TableFn struct {
	Add        bool `json:"add"`
	BatchAdd   bool `json:"batchAdd"`
	Edit       bool `json:"edit"`
	BatchEdit  bool `json:"batchEdit"`
	Detail     bool `json:"detail"`
	Multi      bool `json:"multi"`
	QueryOne   bool `json:"queryOne"`
	Query      bool `json:"query"`
	Count      bool `json:"count"`
	List       bool `json:"list"`
	All        bool `json:"all"`
	DelOne     bool `json:"delOne"`
	Del        bool `json:"del"`
	EnableOne  bool `json:"enableOne"`
	DisableOne bool `json:"disableOne"`
	Enable     bool `json:"enable"`
	Disable    bool `json:"disable"`
	Tree       bool `json:"tree"`
}

// WithServiceEdgeCondition _
type WithServiceEdgeCondition struct {
	Field string
	Value string
}

// WithServiceEdge 带上其他微服务数据
type WithServiceEdge struct {
	Type         string // O2M M2O M2M 暂时支持 M2O
	EdgeName     string // 边名
	Field        string // 字段名
	DataField    string // 增加存放数据的字段名
	Reply        string // 响应字段名
	QueryType    string // 查询类型 缓存 or db 默认缓存
	IsWithDetail bool   // 是否带上详情
	IsWithRows   bool   // 是否带上列表
	Condition    WithServiceEdgeCondition
}

// WithService 带上其他微服务数据
type WithService struct {
	Name      string
	ProtoPath string
	Package   string
	Usecase   string
	Edges     []WithServiceEdge
}

// Aggregate 聚合与分组
type Aggregate struct {
	// 聚合函数名，为空则 GroupBy 组合
	FuName  string
	GroupBy []string
	// 如果 Aggregate 为空，则默认为 count
	Aggregate []AggregateConf
	// 当没有查询条件时 是否缓存数据
	IsAllCache bool
}

// AggregateConf _
type AggregateConf struct {
	Type  AggregateConfType
	Field string
	As    string
}

// AggregateConfType 聚合类型
type AggregateConfType string

const (
	// AggregateTypeCount 聚合类型 count
	AggregateTypeCount AggregateConfType = "count"
	// AggregateTypeSum 聚合类型 sum
	AggregateTypeSum AggregateConfType = "sum"
	// AggregateTypeMean 聚合类型 mean
	AggregateTypeMean AggregateConfType = "mean"
	// AggregateTypeMax 聚合类型 max
	AggregateTypeMax AggregateConfType = "max"
	// AggregateTypeMin 聚合类型 min
	AggregateTypeMin AggregateConfType = "min"
)

// OpProtectConfig 操作保护的配置 添加 / 编辑 / 删除 / 启用 / 禁用 根据此配置在 biz 层进行插入函数，函数给于用户自行编写
type OpProtectConfig struct {
	// 添加配置
	Add bool `json:"add"`
	// 编辑配置
	Edit bool `json:"edit"`
	// 删除配置
	Del bool `json:"del"`
	// 启用配置
	Enable bool `json:"enable"`
	// 禁用配置
	Disable bool `json:"disable"`
	// Portal 启用配置
	PortalEnable bool `json:"portalEnable"`
	// Portal 禁用配置
	PortalDisable bool `json:"portalDisable"`
}

// MutationConfig 突变配置
type MutationConfig struct {
	// 启用
	Enable bool `json:"enable"`
	// 操作记录表 不填则不启用 需手动创建表,会对表进行需要的字段进行检查
	OpLogTable string `json:"opLogTable"`
	// 操作者 获取语句 必须以 operator 开头 例如 operator := meta.GetUserID(ctx)
	OperatorCode string `json:"operatorCode"`
	// TODO 同步到搜索

}

// CustomField 自定义字段配置
// 用于定义业务逻辑字段，这些字段不存在于数据库中，仅用于代码生成
type CustomField struct {
	// Name 字段名称，用于生成结构体字段名和 JSON 标签
	Name string `json:"name"`

	// Comment 字段注释，用于生成代码注释和 proto 文件中的注释
	Comment string `json:"comment"`

	// PbType protobuf 类型定义，用于生成 proto 文件中的字段类型
	PbType string `json:"pbType"`

	// PbIndex protobuf 字段序号，用于生成 proto 文件中的字段编号
	PbIndex int `json:"pbIndex"`

	// GoType Go 语言类型定义，用于生成 Go 结构体中的字段类型
	GoType string `json:"goType"`

	// GoTypeToPb Go 类型转 protobuf 类型的处理代码，如无则直接使用 = 赋值
	// 例如：当需要特殊转换逻辑时，可以指定转换代码
	GoTypeToPb string `json:"goTypeToPb"`

	// DetailReply 是否在详情接口的响应中包含此字段
	DetailReply bool `json:"detailReply"`

	// QueryReply 是否在查询接口的响应中包含此字段
	QueryReply bool `json:"queryReply"`

	// PortalDetailReply 是否在 Portal 层详情接口的响应中包含此字段
	PortalDetailReply bool `json:"portalDetailReply"`

	// PortalQueryReply 是否在 Portal 层查询接口的响应中包含此字段
	PortalQueryReply bool `json:"portalQueryReply"`

	// Biz 是否在业务层（biz）生成该字段，true 表示在 biz 层结构体中包含此字段
	Biz bool `json:"biz"`
	// BizDetailGetCode 业务层详情接口获取值的代码
	BizDetailGetCode string `json:"bizDetailGetCode"`
	// BizQueryGetCode 业务层查询接口获取值的代码
	BizQueryGetCode string `json:"bizQueryGetCode"`
	// DataGetCode 字段在数据层获取值的代码
	// 用于在数据层生成该字段的值，通常用于计算字段或从其他数据源获取数据
	DataGetCode string `json:"dataGetCode"`
}
