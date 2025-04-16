// cSpell: disable
package lang

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// MsgKey message key
// MsgKey 消息键类型
type MsgKey string

// String 将 MsgKey 转换为字符串
func (m MsgKey) String() string {
	return string(m)
}

const (
	validatePrefix MsgKey = "validate_"
	// 根据文档 还有文档内的各个 ***Rules 生成更个验证的多语言内容
	// https://github.com/bufbuild/protovalidate/blob/main/proto/protovalidate/buf/validate/validate.proto
	// 基本数据类型验证规则
	// FloatRules float = 1;
	// DoubleRules double = 2;
	// Int32Rules int32 = 3;
	// Int64Rules int64 = 4;
	// UInt32Rules uint32 = 5;
	// UInt64Rules uint64 = 6;
	// SInt32Rules sint32 = 7;
	// SInt64Rules sint64 = 8;
	// Fixed32Rules fixed32 = 9;
	// Fixed64Rules fixed64 = 10;
	// SFixed32Rules sfixed32 = 11;
	// SFixed64Rules sfixed64 = 12;
	// BoolRules bool = 13;
	// StringRules string = 14;
	// BytesRules bytes = 15;

	// 复杂数据类型验证规则
	// EnumRules enum = 16;
	// MessageRules message = 17;
	// RepeatedRules repeated = 18;
	// MapRules map = 19;

	// 特殊类型验证规则
	// AnyRules any = 20;
	// DurationRules duration = 21;
	// TimestampRules timestamp = 22;

	// 通用错误消息键
	parameterErrorKey            MsgKey = "parameter_error"              // 参数错误
	getMetadataFailKey           MsgKey = "get_metadata_fail"            // 获取元数据失败
	getMetadataConversionFailKey MsgKey = "get_metadata_conversion_fail" // 元数据转换失败
	missingMetadataKey           MsgKey = "missing_metadata"             // 缺少元数据
	encryptFailKey               MsgKey = "encrypt_fail"                 // 加密失败
	decryptFailKey               MsgKey = "decrypt_fail"                 // 解密失败
	paramCanNotEmptyKey          MsgKey = "param_can_not_empty"          // 参数不能为空
	notEditableKey               MsgKey = "not_editable"                 // 数据不可编辑
	requestErrorKey              MsgKey = "request_error"                // 请求错误
	pleaseLoginKey               MsgKey = "please_login"                 // 请先登录
	needReLoginKey               MsgKey = "need_re_login"                // 需要重新登录
	accountDisabledKey           MsgKey = "account_disabled"             // 账号已禁用
	pleaseChangePasswordKey      MsgKey = "please_change_password"       // 请修改密码
	passwordErrorKey             MsgKey = "password_error"               // 密码错误
	noPermissionKey              MsgKey = "no_permission"                // 无权限
	notConfiguredKey             MsgKey = "not_configured"               // 未配置
	canNotEmptyKey               MsgKey = "can_not_empty"                // 不能为空
	expiredKey                   MsgKey = "expired"                      // 已过期

	// 数据相关错误消息键
	dataAbnormalKey        MsgKey = "data_abnormal"         // 数据异常
	dataNotFoundKey        MsgKey = "data_not_found"        // 数据未找到
	dataDuplicateKey       MsgKey = "data_duplicate"        // 数据重复
	dataConstraintKey      MsgKey = "data_constraint"       // 数据约束错误
	dataNotLoadedKey       MsgKey = "data_not_loaded"       // 数据未加载
	dataNotSingularKey     MsgKey = "data_not_singular"     // 数据非唯一
	dataValidationErrorKey MsgKey = "data_validation_error" // 数据验证错误
	dataErrorKey           MsgKey = "data_error"            // 数据错误

	// 缓存相关错误消息键
	cacheNotFoundKey        MsgKey = "cache_not_found"          // 缓存未找到
	cachePreMatchGetFailKey MsgKey = "cache_pre_match_get_fail" // 缓存预匹配获取失败
	cacheSetFailKey         MsgKey = "cache_set_fail"           // 缓存设置失败
	cacheMSetFailKey        MsgKey = "cache_mset_fail"          // 缓存批量设置失败
	cacheDelFailKey         MsgKey = "cache_del_fail"           // 缓存删除失败
	cachePreMatchDelFailKey MsgKey = "cache_pre_match_del_fail" // 缓存预匹配删除失败
	cacheFlushFailKey       MsgKey = "cache_flush_fail"         // 缓存清空失败
	cacheMGetFailKey        MsgKey = "cache_mget_fail"          // 缓存批量获取失败
	cacheMDelFailKey        MsgKey = "cache_mdel_fail"          // 缓存批量删除失败

	// 必填
	validateRequiredKey MsgKey = validatePrefix + "required" // 必填验证
	// 字符串验证消息键
	validateStringConstKey             MsgKey = validatePrefix + "string.const"               // 字符串常量验证
	validateStringLenKey               MsgKey = validatePrefix + "string.len"                 // 字符串长度验证
	validateStringMinLenKey            MsgKey = validatePrefix + "string.min_len"             // 字符串最小长度验证
	validateStringMaxLenKey            MsgKey = validatePrefix + "string.max_len"             // 字符串最大长度验证
	validateStringPatternKey           MsgKey = validatePrefix + "string.pattern"             // 字符串正则表达式验证
	validateStringPrefixKey            MsgKey = validatePrefix + "string.prefix"              // 字符串前缀验证
	validateStringSuffixKey            MsgKey = validatePrefix + "string.suffix"              // 字符串后缀验证
	validateStringContainsKey          MsgKey = validatePrefix + "string.contains"            // 字符串包含验证
	validateStringNotContainsKey       MsgKey = validatePrefix + "string.not_contains"        // 字符串不包含验证
	validateStringInKey                MsgKey = validatePrefix + "string.in"                  // 字符串枚举验证
	validateStringNotInKey             MsgKey = validatePrefix + "string.not_in"              // 字符串非枚举验证
	validateStringEmailKey             MsgKey = validatePrefix + "string.email"               // 邮箱格式验证
	validateStringHostnameKey          MsgKey = validatePrefix + "string.hostname"            // 主机名验证
	validateStringIPKey                MsgKey = validatePrefix + "string.ip"                  // IP地址验证
	validateStringIPv4Key              MsgKey = validatePrefix + "string.ipv4"                // IPv4地址验证
	validateStringIPv6Key              MsgKey = validatePrefix + "string.ipv6"                // IPv6地址验证
	validateStringURIKey               MsgKey = validatePrefix + "string.uri"                 // URI验证
	validateStringURIRefKey            MsgKey = validatePrefix + "string.uri_ref"             // URI引用验证
	validateStringAddressKey           MsgKey = validatePrefix + "string.address"             // 地址验证
	validateStringUUIDKey              MsgKey = validatePrefix + "string.uuid"                // UUID验证
	validateStringWellKnownRegexKey    MsgKey = validatePrefix + "string.well_known_regex"    // 预定义正则表达式验证
	validateStringLenBytesKey          MsgKey = validatePrefix + "string.len_bytes"           // 字节长度验证
	validateStringMinBytesKey          MsgKey = validatePrefix + "string.min_bytes"           // 最小字节长度验证
	validateStringMaxBytesKey          MsgKey = validatePrefix + "string.max_bytes"           // 最大字节长度验证
	validateStringTUUIDKey             MsgKey = validatePrefix + "string.tuuid"               // 时间UUID验证
	validateStringIPPrefixKey          MsgKey = validatePrefix + "string.ip_prefix"           // IP前缀验证
	validateStringIPWithPrefixlenKey   MsgKey = validatePrefix + "string.ip_with_prefixlen"   // IP带前缀长度验证
	validateStringIPv4WithPrefixlenKey MsgKey = validatePrefix + "string.ipv4_with_prefixlen" // IPv4带前缀长度验证
	validateStringIPv6WithPrefixlenKey MsgKey = validatePrefix + "string.ipv6_with_prefixlen" // IPv6带前缀长度验证
	validateStringStrictKey            MsgKey = validatePrefix + "string.strict"              // 严格模式验证
	validateStringExampleKey           MsgKey = validatePrefix + "string.example"             // 示例值验证

	// 布尔值验证消息键
	validateBoolConstKey MsgKey = validatePrefix + "bool.const" // 布尔常量验证

	// 枚举验证消息键
	validateEnumConstKey   MsgKey = validatePrefix + "enum.const"   // 枚举常量验证
	validateEnumDefinedKey MsgKey = validatePrefix + "enum.defined" // 枚举定义验证
	validateEnumInKey      MsgKey = validatePrefix + "enum.in"      // 枚举范围验证
	validateEnumNotInKey   MsgKey = validatePrefix + "enum.not_in"  // 枚举非范围验证

	// 重复字段验证消息键
	validateRepeatedMinItemsKey MsgKey = validatePrefix + "repeated.min_items" // 最小元素数量验证
	validateRepeatedMaxItemsKey MsgKey = validatePrefix + "repeated.max_items" // 最大元素数量验证
	validateRepeatedUniqueKey   MsgKey = validatePrefix + "repeated.unique"    // 元素唯一性验证
	validateRepeatedItemsKey    MsgKey = validatePrefix + "repeated.items"     // 元素验证

	// Map验证消息键
	validateMapMinPairsKey MsgKey = validatePrefix + "map.min_pairs" // 最小键值对数量验证
	validateMapMaxPairsKey MsgKey = validatePrefix + "map.max_pairs" // 最大键值对数量验证
	validateMapNoSparseKey MsgKey = validatePrefix + "map.no_sparse" // 非稀疏验证
	validateMapKeysKey     MsgKey = validatePrefix + "map.keys"      // 键验证
	validateMapValuesKey   MsgKey = validatePrefix + "map.values"    // 值验证

	// Float验证消息键
	validateFloatConstKey  MsgKey = validatePrefix + "float.const"  // 浮点常量验证
	validateFloatLtKey     MsgKey = validatePrefix + "float.lt"     // 小于验证
	validateFloatLteKey    MsgKey = validatePrefix + "float.lte"    // 小于等于验证
	validateFloatGtKey     MsgKey = validatePrefix + "float.gt"     // 大于验证
	validateFloatGteKey    MsgKey = validatePrefix + "float.gte"    // 大于等于验证
	validateFloatInKey     MsgKey = validatePrefix + "float.in"     // 数值枚举验证
	validateFloatNotInKey  MsgKey = validatePrefix + "float.not_in" // 数值非枚举验证
	validateFloatFiniteKey MsgKey = validatePrefix + "float.finite" // 有限数验证

	// Double验证消息键
	validateDoubleConstKey  MsgKey = validatePrefix + "double.const"  // 双精度常量验证
	validateDoubleLtKey     MsgKey = validatePrefix + "double.lt"     // 小于验证
	validateDoubleLteKey    MsgKey = validatePrefix + "double.lte"    // 小于等于验证
	validateDoubleGtKey     MsgKey = validatePrefix + "double.gt"     // 大于验证
	validateDoubleGteKey    MsgKey = validatePrefix + "double.gte"    // 大于等于验证
	validateDoubleInKey     MsgKey = validatePrefix + "double.in"     // 数值枚举验证
	validateDoubleNotInKey  MsgKey = validatePrefix + "double.not_in" // 数值非枚举验证
	validateDoubleFiniteKey MsgKey = validatePrefix + "double.finite" // 有限数验证

	// Int32验证消息键
	validateInt32ConstKey MsgKey = validatePrefix + "int32.const"  // 32位整数常量验证
	validateInt32LtKey    MsgKey = validatePrefix + "int32.lt"     // 小于验证
	validateInt32LteKey   MsgKey = validatePrefix + "int32.lte"    // 小于等于验证
	validateInt32GtKey    MsgKey = validatePrefix + "int32.gt"     // 大于验证
	validateInt32GteKey   MsgKey = validatePrefix + "int32.gte"    // 大于等于验证
	validateInt32InKey    MsgKey = validatePrefix + "int32.in"     // 数值枚举验证
	validateInt32NotInKey MsgKey = validatePrefix + "int32.not_in" // 数值非枚举验证

	// Int64 validation
	validateInt64ConstKey MsgKey = validatePrefix + "int64.const"  // 64位整数常量验证
	validateInt64LtKey    MsgKey = validatePrefix + "int64.lt"     // 小于验证
	validateInt64LteKey   MsgKey = validatePrefix + "int64.lte"    // 小于等于验证
	validateInt64GtKey    MsgKey = validatePrefix + "int64.gt"     // 大于验证
	validateInt64GteKey   MsgKey = validatePrefix + "int64.gte"    // 大于等于验证
	validateInt64InKey    MsgKey = validatePrefix + "int64.in"     // 数值枚举验证
	validateInt64NotInKey MsgKey = validatePrefix + "int64.not_in" // 数值非枚举验证

	// UInt32 validation
	validateUInt32ConstKey MsgKey = validatePrefix + "uint32.const"  // 32位无符号整数常量验证
	validateUInt32LtKey    MsgKey = validatePrefix + "uint32.lt"     // 小于验证
	validateUInt32LteKey   MsgKey = validatePrefix + "uint32.lte"    // 小于等于验证
	validateUInt32GtKey    MsgKey = validatePrefix + "uint32.gt"     // 大于验证
	validateUInt32GteKey   MsgKey = validatePrefix + "uint32.gte"    // 大于等于验证
	validateUInt32InKey    MsgKey = validatePrefix + "uint32.in"     // 数值枚举验证
	validateUInt32NotInKey MsgKey = validatePrefix + "uint32.not_in" // 数值非枚举验证

	// UInt64 validation
	validateUInt64ConstKey MsgKey = validatePrefix + "uint64.const"  // 64位无符号整数常量验证
	validateUInt64LtKey    MsgKey = validatePrefix + "uint64.lt"     // 小于验证
	validateUInt64LteKey   MsgKey = validatePrefix + "uint64.lte"    // 小于等于验证
	validateUInt64GtKey    MsgKey = validatePrefix + "uint64.gt"     // 大于验证
	validateUInt64GteKey   MsgKey = validatePrefix + "uint64.gte"    // 大于等于验证
	validateUInt64InKey    MsgKey = validatePrefix + "uint64.in"     // 数值枚举验证
	validateUInt64NotInKey MsgKey = validatePrefix + "uint64.not_in" // 数值非枚举验证

	// SInt32 validation
	validateSInt32ConstKey MsgKey = validatePrefix + "sint32.const"  // 32位有符号整数常量验证
	validateSInt32LtKey    MsgKey = validatePrefix + "sint32.lt"     // 小于验证
	validateSInt32LteKey   MsgKey = validatePrefix + "sint32.lte"    // 小于等于验证
	validateSInt32GtKey    MsgKey = validatePrefix + "sint32.gt"     // 大于验证
	validateSInt32GteKey   MsgKey = validatePrefix + "sint32.gte"    // 大于等于验证
	validateSInt32InKey    MsgKey = validatePrefix + "sint32.in"     // 数值枚举验证
	validateSInt32NotInKey MsgKey = validatePrefix + "sint32.not_in" // 数值非枚举验证

	// SInt64 validation
	validateSInt64ConstKey MsgKey = validatePrefix + "sint64.const"  // 64位有符号整数常量验证
	validateSInt64LtKey    MsgKey = validatePrefix + "sint64.lt"     // 小于验证
	validateSInt64LteKey   MsgKey = validatePrefix + "sint64.lte"    // 小于等于验证
	validateSInt64GtKey    MsgKey = validatePrefix + "sint64.gt"     // 大于验证
	validateSInt64GteKey   MsgKey = validatePrefix + "sint64.gte"    // 大于等于验证
	validateSInt64InKey    MsgKey = validatePrefix + "sint64.in"     // 数值枚举验证
	validateSInt64NotInKey MsgKey = validatePrefix + "sint64.not_in" // 数值非枚举验证

	// Fixed32 validation
	validateFixed32ConstKey MsgKey = validatePrefix + "fixed32.const"  // 32位定点数常量验证
	validateFixed32LtKey    MsgKey = validatePrefix + "fixed32.lt"     // 小于验证
	validateFixed32LteKey   MsgKey = validatePrefix + "fixed32.lte"    // 小于等于验证
	validateFixed32GtKey    MsgKey = validatePrefix + "fixed32.gt"     // 大于验证
	validateFixed32GteKey   MsgKey = validatePrefix + "fixed32.gte"    // 大于等于验证
	validateFixed32InKey    MsgKey = validatePrefix + "fixed32.in"     // 数值枚举验证
	validateFixed32NotInKey MsgKey = validatePrefix + "fixed32.not_in" // 数值非枚举验证

	// Fixed64 validation
	validateFixed64ConstKey MsgKey = validatePrefix + "fixed64.const"  // 64位定点数常量验证
	validateFixed64LtKey    MsgKey = validatePrefix + "fixed64.lt"     // 小于验证
	validateFixed64LteKey   MsgKey = validatePrefix + "fixed64.lte"    // 小于等于验证
	validateFixed64GtKey    MsgKey = validatePrefix + "fixed64.gt"     // 大于验证
	validateFixed64GteKey   MsgKey = validatePrefix + "fixed64.gte"    // 大于等于验证
	validateFixed64InKey    MsgKey = validatePrefix + "fixed64.in"     // 数值枚举验证
	validateFixed64NotInKey MsgKey = validatePrefix + "fixed64.not_in" // 数值非枚举验证

	// SFixed32 validation
	validateSFixed32ConstKey MsgKey = validatePrefix + "sfixed32.const"  // 32位有符号定点数常量验证
	validateSFixed32LtKey    MsgKey = validatePrefix + "sfixed32.lt"     // 小于验证
	validateSFixed32LteKey   MsgKey = validatePrefix + "sfixed32.lte"    // 小于等于验证
	validateSFixed32GtKey    MsgKey = validatePrefix + "sfixed32.gt"     // 大于验证
	validateSFixed32GteKey   MsgKey = validatePrefix + "sfixed32.gte"    // 大于等于验证
	validateSFixed32InKey    MsgKey = validatePrefix + "sfixed32.in"     // 数值枚举验证
	validateSFixed32NotInKey MsgKey = validatePrefix + "sfixed32.not_in" // 数值非枚举验证

	// SFixed64 validation
	validateSFixed64ConstKey MsgKey = validatePrefix + "sfixed64.const"  // 64位有符号定点数常量验证
	validateSFixed64LtKey    MsgKey = validatePrefix + "sfixed64.lt"     // 小于验证
	validateSFixed64LteKey   MsgKey = validatePrefix + "sfixed64.lte"    // 小于等于验证
	validateSFixed64GtKey    MsgKey = validatePrefix + "sfixed64.gt"     // 大于验证
	validateSFixed64GteKey   MsgKey = validatePrefix + "sfixed64.gte"    // 大于等于验证
	validateSFixed64InKey    MsgKey = validatePrefix + "sfixed64.in"     // 数值枚举验证
	validateSFixed64NotInKey MsgKey = validatePrefix + "sfixed64.not_in" // 数值非枚举验证

	// Bytes validation
	validateBytesConstKey    MsgKey = validatePrefix + "bytes.const"    // 字节常量验证
	validateBytesLenKey      MsgKey = validatePrefix + "bytes.len"      // 字节长度验证
	validateBytesMinLenKey   MsgKey = validatePrefix + "bytes.min_len"  // 最小字节长度验证
	validateBytesMaxLenKey   MsgKey = validatePrefix + "bytes.max_len"  // 最大字节长度验证
	validateBytesPatternKey  MsgKey = validatePrefix + "bytes.pattern"  // 字节模式验证
	validateBytesPrefixKey   MsgKey = validatePrefix + "bytes.prefix"   // 字节前缀验证
	validateBytesSuffixKey   MsgKey = validatePrefix + "bytes.suffix"   // 字节后缀验证
	validateBytesContainsKey MsgKey = validatePrefix + "bytes.contains" // 字节包含验证
	validateBytesInKey       MsgKey = validatePrefix + "bytes.in"       // 字节枚举验证
	validateBytesNotInKey    MsgKey = validatePrefix + "bytes.not_in"   // 字节非枚举验证

	// Message validation
	validateMessageRequiredKey MsgKey = validatePrefix + "message.required" // 消息必填验证
	validateMessageSkipKey     MsgKey = validatePrefix + "message.skip"     // 消息跳过验证

	// Any validation
	validateAnyRequiredKey MsgKey = validatePrefix + "any.required" // Any 必填验证
	validateAnyInKey       MsgKey = validatePrefix + "any.in"       // Any 类型枚举验证
	validateAnyNotInKey    MsgKey = validatePrefix + "any.not_in"   // Any 类型非枚举验证

	// Duration validation
	validateDurationRequiredKey MsgKey = validatePrefix + "duration.required" // Duration 必填验证
	validateDurationConstKey    MsgKey = validatePrefix + "duration.const"    // Duration 常量验证
	validateDurationLtKey       MsgKey = validatePrefix + "duration.lt"       // Duration 小于验证
	validateDurationLteKey      MsgKey = validatePrefix + "duration.lte"      // Duration 小于等于验证
	validateDurationGtKey       MsgKey = validatePrefix + "duration.gt"       // Duration 大于验证
	validateDurationGteKey      MsgKey = validatePrefix + "duration.gte"      // Duration 大于等于验证
	validateDurationInKey       MsgKey = validatePrefix + "duration.in"       // Duration 枚举验证
	validateDurationNotInKey    MsgKey = validatePrefix + "duration.not_in"   // Duration 非枚举验证

	// Timestamp validation
	validateTimestampRequiredKey MsgKey = validatePrefix + "timestamp.required" // Timestamp 必填验证
	validateTimestampConstKey    MsgKey = validatePrefix + "timestamp.const"    // Timestamp 常量验证
	validateTimestampLtKey       MsgKey = validatePrefix + "timestamp.lt"       // Timestamp 小于验证
	validateTimestampLteKey      MsgKey = validatePrefix + "timestamp.lte"      // Timestamp 小于等于验证
	validateTimestampGtKey       MsgKey = validatePrefix + "timestamp.gt"       // Timestamp 大于验证
	validateTimestampGteKey      MsgKey = validatePrefix + "timestamp.gte"      // Timestamp 大于等于验证
	validateTimestampLtNowKey    MsgKey = validatePrefix + "timestamp.lt_now"   // Timestamp 小于当前时间验证
	validateTimestampGtNowKey    MsgKey = validatePrefix + "timestamp.gt_now"   // Timestamp 大于当前时间验证
	validateTimestampWithinKey   MsgKey = validatePrefix + "timestamp.within"   // Timestamp 在时间范围内验证
)

// 默认语言的 msg 配置 英语
var dfMsgMap = map[MsgKey]*i18n.Message{
	parameterErrorKey:            {ID: parameterErrorKey.String(), Other: "Parameter error, please check your parameters"},
	getMetadataFailKey:           {ID: getMetadataFailKey.String(), Other: "Get metadata failed"},
	getMetadataConversionFailKey: {ID: getMetadataConversionFailKey.String(), Other: "metadata {{.Source}} conversion to {{.Target}} failed"},
	missingMetadataKey: {
		ID:    missingMetadataKey.String(),
		Other: "Missing metadata {{.Name}}, please check whether the transmission link is enabled for metadata transmission and pass the value",
	},
	encryptFailKey:          {ID: encryptFailKey.String(), Other: "Encryption failed"},
	decryptFailKey:          {ID: decryptFailKey.String(), Other: "Decryption failed"},
	paramCanNotEmptyKey:     {ID: paramCanNotEmptyKey.String(), Other: "Parameter {{.Name}} cannot be empty"},
	notEditableKey:          {ID: notEditableKey.String(), Other: "The data is not editable"},
	requestErrorKey:         {ID: requestErrorKey.String(), Other: "Request error"},
	pleaseLoginKey:          {ID: pleaseLoginKey.String(), Other: "Please login first"},
	needReLoginKey:          {ID: needReLoginKey.String(), Other: "Need to re login"},
	accountDisabledKey:      {ID: accountDisabledKey.String(), Other: "Account has been disabled"},
	pleaseChangePasswordKey: {ID: pleaseChangePasswordKey.String(), Other: "Please change your password first"},
	passwordErrorKey:        {ID: passwordErrorKey.String(), Other: "Password error"},
	noPermissionKey:         {ID: noPermissionKey.String(), Other: "No permission"},
	notConfiguredKey:        {ID: notConfiguredKey.String(), Other: "{{.Name}} not configured"},
	canNotEmptyKey:          {ID: canNotEmptyKey.String(), Other: "{{.Name}} can not be empty"},
	expiredKey:              {ID: expiredKey.String(), Other: "{{.Name}} has expired"},

	dataAbnormalKey:        {ID: dataAbnormalKey.String(), Other: "{{.Name}} data abnormal"},
	dataNotFoundKey:        {ID: dataNotFoundKey.String(), Other: "Data not found"},
	dataDuplicateKey:       {ID: dataDuplicateKey.String(), Other: "The data already exists, please do not add it repeatedly"},
	dataConstraintKey:      {ID: dataConstraintKey.String(), Other: "Data constraint check failed, please check your parameters"},
	dataNotLoadedKey:       {ID: dataNotLoadedKey.String(), Other: "Database not loaded, please contact the administrator"},
	dataNotSingularKey:     {ID: dataNotSingularKey.String(), Other: "Data error Not Singular, please contact the administrator"},
	dataValidationErrorKey: {ID: dataValidationErrorKey.String(), Other: "Data validation failed, please check your parameters"},
	dataErrorKey:           {ID: dataErrorKey.String(), Other: "Data layer error, please contact the administrator"},

	cacheNotFoundKey:        {ID: cacheNotFoundKey.String(), Other: "Cache not found"},
	cachePreMatchGetFailKey: {ID: cachePreMatchGetFailKey.String(), Other: "Pre-match cache get failed"},
	cacheSetFailKey:         {ID: cacheSetFailKey.String(), Other: "Set cache failed"},
	cacheMSetFailKey:        {ID: cacheMSetFailKey.String(), Other: "Batch setting cache failed"},
	cacheDelFailKey:         {ID: cacheDelFailKey.String(), Other: "Delete cache failed"},
	cachePreMatchDelFailKey: {ID: cachePreMatchDelFailKey.String(), Other: "Pre-match delete cache failed"},
	cacheFlushFailKey:       {ID: cacheFlushFailKey.String(), Other: "Flush cache failed"},
	cacheMGetFailKey:        {ID: cacheMGetFailKey.String(), Other: "Batch get cache failed"},
	cacheMDelFailKey:        {ID: cacheMDelFailKey.String(), Other: "Batch delete cache failed"},

	validateRequiredKey: {ID: validateRequiredKey.String(), Other: "This field is required"},
	// String validation messages
	validateStringConstKey:             {ID: validateStringConstKey.String(), Other: "Must be exactly '{{.Rule}}'"},
	validateStringLenKey:               {ID: validateStringLenKey.String(), Other: "Length must be exactly {{.Rule}} characters"},
	validateStringMinLenKey:            {ID: validateStringMinLenKey.String(), Other: "Length must be at least {{.Rule}} characters"},
	validateStringMaxLenKey:            {ID: validateStringMaxLenKey.String(), Other: "Length must be at most {{.Rule}} characters"},
	validateStringPatternKey:           {ID: validateStringPatternKey.String(), Other: "Must match the pattern '{{.Rule}}'"},
	validateStringPrefixKey:            {ID: validateStringPrefixKey.String(), Other: "Must start with '{{.Rule}}'"},
	validateStringSuffixKey:            {ID: validateStringSuffixKey.String(), Other: "Must end with '{{.Rule}}'"},
	validateStringContainsKey:          {ID: validateStringContainsKey.String(), Other: "Must contain '{{.Rule}}'"},
	validateStringNotContainsKey:       {ID: validateStringNotContainsKey.String(), Other: "Must not contain '{{.Rule}}'"},
	validateStringInKey:                {ID: validateStringInKey.String(), Other: "Must be one of [{{.Rule}}]"},
	validateStringNotInKey:             {ID: validateStringNotInKey.String(), Other: "Must not be one of [{{.Rule}}]"},
	validateStringEmailKey:             {ID: validateStringEmailKey.String(), Other: "Must be a valid email address"},
	validateStringHostnameKey:          {ID: validateStringHostnameKey.String(), Other: "Must be a valid hostname"},
	validateStringIPKey:                {ID: validateStringIPKey.String(), Other: "Must be a valid IP address"},
	validateStringIPv4Key:              {ID: validateStringIPv4Key.String(), Other: "Must be a valid IPv4 address"},
	validateStringIPv6Key:              {ID: validateStringIPv6Key.String(), Other: "Must be a valid IPv6 address"},
	validateStringURIKey:               {ID: validateStringURIKey.String(), Other: "Must be a valid URI"},
	validateStringURIRefKey:            {ID: validateStringURIRefKey.String(), Other: "Must be a valid URI reference"},
	validateStringAddressKey:           {ID: validateStringAddressKey.String(), Other: "Must be a valid address"},
	validateStringUUIDKey:              {ID: validateStringUUIDKey.String(), Other: "Must be a valid UUID"},
	validateStringWellKnownRegexKey:    {ID: validateStringWellKnownRegexKey.String(), Other: "Must match the format {{.Rule}}"},
	validateStringLenBytesKey:          {ID: validateStringLenBytesKey.String(), Other: "Must be exactly {{.Rule}} bytes"},
	validateStringMinBytesKey:          {ID: validateStringMinBytesKey.String(), Other: "Must be at least {{.Rule}} bytes"},
	validateStringMaxBytesKey:          {ID: validateStringMaxBytesKey.String(), Other: "Must be at most {{.Rule}} bytes"},
	validateStringTUUIDKey:             {ID: validateStringTUUIDKey.String(), Other: "Must be a valid time-based UUID"},
	validateStringIPPrefixKey:          {ID: validateStringIPPrefixKey.String(), Other: "Must be a valid IP network prefix"},
	validateStringIPWithPrefixlenKey:   {ID: validateStringIPWithPrefixlenKey.String(), Other: "Must be a valid IP address with prefix length"},
	validateStringIPv4WithPrefixlenKey: {ID: validateStringIPv4WithPrefixlenKey.String(), Other: "Must be a valid IPv4 address with prefix length"},
	validateStringIPv6WithPrefixlenKey: {ID: validateStringIPv6WithPrefixlenKey.String(), Other: "Must be a valid IPv6 address with prefix length"},
	validateStringStrictKey:            {ID: validateStringStrictKey.String(), Other: "Must match strict validation rules"},
	validateStringExampleKey:           {ID: validateStringExampleKey.String(), Other: "Must match the example format: {{.Rule}}"},

	// Boolean validation messages
	validateBoolConstKey: {ID: validateBoolConstKey.String(), Other: "Must be {{.Rule}}"},

	// Enum validation messages
	validateEnumConstKey:   {ID: validateEnumConstKey.String(), Other: "Must be exactly {{.Rule}}"},
	validateEnumDefinedKey: {ID: validateEnumDefinedKey.String(), Other: "Must be a defined enum value"},
	validateEnumInKey:      {ID: validateEnumInKey.String(), Other: "Must be one of {{.Rule}}"},
	validateEnumNotInKey:   {ID: validateEnumNotInKey.String(), Other: "Must not be one of {{.Rule}}"},

	// Repeated field validation messages
	validateRepeatedMinItemsKey: {ID: validateRepeatedMinItemsKey.String(), Other: "Must have at least {{.Rule}} items"},
	validateRepeatedMaxItemsKey: {ID: validateRepeatedMaxItemsKey.String(), Other: "Must have at most {{.Rule}} items"},
	validateRepeatedUniqueKey:   {ID: validateRepeatedUniqueKey.String(), Other: "All items must be unique"},
	validateRepeatedItemsKey:    {ID: validateRepeatedItemsKey.String(), Other: "Items must satisfy the rules"},

	// Map validation messages
	validateMapMinPairsKey: {ID: validateMapMinPairsKey.String(), Other: "Must have at least {{.Rule}} key-value pairs"},
	validateMapMaxPairsKey: {ID: validateMapMaxPairsKey.String(), Other: "Must have at most {{.Rule}} key-value pairs"},
	validateMapNoSparseKey: {ID: validateMapNoSparseKey.String(), Other: "Must not contain null values"},
	validateMapKeysKey:     {ID: validateMapKeysKey.String(), Other: "Keys must satisfy the rules"},
	validateMapValuesKey:   {ID: validateMapValuesKey.String(), Other: "Values must satisfy the rules"},

	// Float validation messages
	validateFloatConstKey:  {ID: validateFloatConstKey.String(), Other: "Must be {{.Rule}}"},
	validateFloatLtKey:     {ID: validateFloatLtKey.String(), Other: "Must be less than {{.Rule}}"},
	validateFloatLteKey:    {ID: validateFloatLteKey.String(), Other: "Must be less than or equal to {{.Rule}}"},
	validateFloatGtKey:     {ID: validateFloatGtKey.String(), Other: "Must be greater than {{.Rule}}"},
	validateFloatGteKey:    {ID: validateFloatGteKey.String(), Other: "Must be greater than or equal to {{.Rule}}"},
	validateFloatInKey:     {ID: validateFloatInKey.String(), Other: "Must be one of {{.Rule}}"},
	validateFloatNotInKey:  {ID: validateFloatNotInKey.String(), Other: "Must not be one of {{.Rule}}"},
	validateFloatFiniteKey: {ID: validateFloatFiniteKey.String(), Other: "Must be finite"},

	// Double validation messages
	validateDoubleConstKey:  {ID: validateDoubleConstKey.String(), Other: "Must be {{.Rule}}"},
	validateDoubleLtKey:     {ID: validateDoubleLtKey.String(), Other: "Must be less than {{.Rule}}"},
	validateDoubleLteKey:    {ID: validateDoubleLteKey.String(), Other: "Must be less than or equal to {{.Rule}}"},
	validateDoubleGtKey:     {ID: validateDoubleGtKey.String(), Other: "Must be greater than {{.Rule}}"},
	validateDoubleGteKey:    {ID: validateDoubleGteKey.String(), Other: "Must be greater than or equal to {{.Rule}}"},
	validateDoubleInKey:     {ID: validateDoubleInKey.String(), Other: "Must be one of {{.Rule}}"},
	validateDoubleNotInKey:  {ID: validateDoubleNotInKey.String(), Other: "Must not be one of {{.Rule}}"},
	validateDoubleFiniteKey: {ID: validateDoubleFiniteKey.String(), Other: "Must be finite"},

	// Int32 validation messages
	validateInt32ConstKey: {ID: validateInt32ConstKey.String(), Other: "Must be {{.Rule}}"},
	validateInt32LtKey:    {ID: validateInt32LtKey.String(), Other: "Must be less than {{.Rule}}"},
	validateInt32LteKey:   {ID: validateInt32LteKey.String(), Other: "Must be less than or equal to {{.Rule}}"},
	validateInt32GtKey:    {ID: validateInt32GtKey.String(), Other: "Must be greater than {{.Rule}}"},
	validateInt32GteKey:   {ID: validateInt32GteKey.String(), Other: "Must be greater than or equal to {{.Rule}}"},
	validateInt32InKey:    {ID: validateInt32InKey.String(), Other: "Must be one of {{.Rule}}"},
	validateInt32NotInKey: {ID: validateInt32NotInKey.String(), Other: "Must not be one of {{.Rule}}"},

	// Int64 validation messages
	validateInt64ConstKey: {ID: validateInt64ConstKey.String(), Other: "Must be {{.Rule}}"},
	validateInt64LtKey:    {ID: validateInt64LtKey.String(), Other: "Must be less than {{.Rule}}"},
	validateInt64LteKey:   {ID: validateInt64LteKey.String(), Other: "Must be less than or equal to {{.Rule}}"},
	validateInt64GtKey:    {ID: validateInt64GtKey.String(), Other: "Must be greater than {{.Rule}}"},
	validateInt64GteKey:   {ID: validateInt64GteKey.String(), Other: "Must be greater than or equal to {{.Rule}}"},
	validateInt64InKey:    {ID: validateInt64InKey.String(), Other: "Must be one of {{.Rule}}"},
	validateInt64NotInKey: {ID: validateInt64NotInKey.String(), Other: "Must not be one of {{.Rule}}"},

	// UInt32 validation messages
	validateUInt32ConstKey: {ID: validateUInt32ConstKey.String(), Other: "Must be {{.Rule}}"},
	validateUInt32LtKey:    {ID: validateUInt32LtKey.String(), Other: "Must be less than {{.Rule}}"},
	validateUInt32LteKey:   {ID: validateUInt32LteKey.String(), Other: "Must be less than or equal to {{.Rule}}"},
	validateUInt32GtKey:    {ID: validateUInt32GtKey.String(), Other: "Must be greater than {{.Rule}}"},
	validateUInt32GteKey:   {ID: validateUInt32GteKey.String(), Other: "Must be greater than or equal to {{.Rule}}"},
	validateUInt32InKey:    {ID: validateUInt32InKey.String(), Other: "Must be one of {{.Rule}}"},
	validateUInt32NotInKey: {ID: validateUInt32NotInKey.String(), Other: "Must not be one of {{.Rule}}"},

	// UInt64 validation messages
	validateUInt64ConstKey: {ID: validateUInt64ConstKey.String(), Other: "Must be {{.Rule}}"},
	validateUInt64LtKey:    {ID: validateUInt64LtKey.String(), Other: "Must be less than {{.Rule}}"},
	validateUInt64LteKey:   {ID: validateUInt64LteKey.String(), Other: "Must be less than or equal to {{.Rule}}"},
	validateUInt64GtKey:    {ID: validateUInt64GtKey.String(), Other: "Must be greater than {{.Rule}}"},
	validateUInt64GteKey:   {ID: validateUInt64GteKey.String(), Other: "Must be greater than or equal to {{.Rule}}"},
	validateUInt64InKey:    {ID: validateUInt64InKey.String(), Other: "Must be one of {{.Rule}}"},
	validateUInt64NotInKey: {ID: validateUInt64NotInKey.String(), Other: "Must not be one of {{.Rule}}"},

	// SInt32 validation messages
	validateSInt32ConstKey: {ID: validateSInt32ConstKey.String(), Other: "Must be {{.Rule}}"},
	validateSInt32LtKey:    {ID: validateSInt32LtKey.String(), Other: "Must be less than {{.Rule}}"},
	validateSInt32LteKey:   {ID: validateSInt32LteKey.String(), Other: "Must be less than or equal to {{.Rule}}"},
	validateSInt32GtKey:    {ID: validateSInt32GtKey.String(), Other: "Must be greater than {{.Rule}}"},
	validateSInt32GteKey:   {ID: validateSInt32GteKey.String(), Other: "Must be greater than or equal to {{.Rule}}"},
	validateSInt32InKey:    {ID: validateSInt32InKey.String(), Other: "Must be one of {{.Rule}}"},
	validateSInt32NotInKey: {ID: validateSInt32NotInKey.String(), Other: "Must not be one of {{.Rule}}"},

	// SInt64 validation messages
	validateSInt64ConstKey: {ID: validateSInt64ConstKey.String(), Other: "Must be {{.Rule}}"},
	validateSInt64LtKey:    {ID: validateSInt64LtKey.String(), Other: "Must be less than {{.Rule}}"},
	validateSInt64LteKey:   {ID: validateSInt64LteKey.String(), Other: "Must be less than or equal to {{.Rule}}"},
	validateSInt64GtKey:    {ID: validateSInt64GtKey.String(), Other: "Must be greater than {{.Rule}}"},
	validateSInt64GteKey:   {ID: validateSInt64GteKey.String(), Other: "Must be greater than or equal to {{.Rule}}"},
	validateSInt64InKey:    {ID: validateSInt64InKey.String(), Other: "Must be one of {{.Rule}}"},
	validateSInt64NotInKey: {ID: validateSInt64NotInKey.String(), Other: "Must not be one of {{.Rule}}"},

	// Fixed32 validation messages
	validateFixed32ConstKey: {ID: validateFixed32ConstKey.String(), Other: "Must be {{.Rule}}"},
	validateFixed32LtKey:    {ID: validateFixed32LtKey.String(), Other: "Must be less than {{.Rule}}"},
	validateFixed32LteKey:   {ID: validateFixed32LteKey.String(), Other: "Must be less than or equal to {{.Rule}}"},
	validateFixed32GtKey:    {ID: validateFixed32GtKey.String(), Other: "Must be greater than {{.Rule}}"},
	validateFixed32GteKey:   {ID: validateFixed32GteKey.String(), Other: "Must be greater than or equal to {{.Rule}}"},
	validateFixed32InKey:    {ID: validateFixed32InKey.String(), Other: "Must be one of {{.Rule}}"},
	validateFixed32NotInKey: {ID: validateFixed32NotInKey.String(), Other: "Must not be one of {{.Rule}}"},

	// Fixed64 validation messages
	validateFixed64ConstKey: {ID: validateFixed64ConstKey.String(), Other: "Must be {{.Rule}}"},
	validateFixed64LtKey:    {ID: validateFixed64LtKey.String(), Other: "Must be less than {{.Rule}}"},
	validateFixed64LteKey:   {ID: validateFixed64LteKey.String(), Other: "Must be less than or equal to {{.Rule}}"},
	validateFixed64GtKey:    {ID: validateFixed64GtKey.String(), Other: "Must be greater than {{.Rule}}"},
	validateFixed64GteKey:   {ID: validateFixed64GteKey.String(), Other: "Must be greater than or equal to {{.Rule}}"},
	validateFixed64InKey:    {ID: validateFixed64InKey.String(), Other: "Must be one of {{.Rule}}"},
	validateFixed64NotInKey: {ID: validateFixed64NotInKey.String(), Other: "Must not be one of {{.Rule}}"},

	// SFixed32 validation messages
	validateSFixed32ConstKey: {ID: validateSFixed32ConstKey.String(), Other: "Must be {{.Rule}}"},
	validateSFixed32LtKey:    {ID: validateSFixed32LtKey.String(), Other: "Must be less than {{.Rule}}"},
	validateSFixed32LteKey:   {ID: validateSFixed32LteKey.String(), Other: "Must be less than or equal to {{.Rule}}"},
	validateSFixed32GtKey:    {ID: validateSFixed32GtKey.String(), Other: "Must be greater than {{.Rule}}"},
	validateSFixed32GteKey:   {ID: validateSFixed32GteKey.String(), Other: "Must be greater than or equal to {{.Rule}}"},
	validateSFixed32InKey:    {ID: validateSFixed32InKey.String(), Other: "Must be one of {{.Rule}}"},
	validateSFixed32NotInKey: {ID: validateSFixed32NotInKey.String(), Other: "Must not be one of {{.Rule}}"},

	// SFixed64 validation messages
	validateSFixed64ConstKey: {ID: validateSFixed64ConstKey.String(), Other: "Must be {{.Rule}}"},
	validateSFixed64LtKey:    {ID: validateSFixed64LtKey.String(), Other: "Must be less than {{.Rule}}"},
	validateSFixed64LteKey:   {ID: validateSFixed64LteKey.String(), Other: "Must be less than or equal to {{.Rule}}"},
	validateSFixed64GtKey:    {ID: validateSFixed64GtKey.String(), Other: "Must be greater than {{.Rule}}"},
	validateSFixed64GteKey:   {ID: validateSFixed64GteKey.String(), Other: "Must be greater than or equal to {{.Rule}}"},
	validateSFixed64InKey:    {ID: validateSFixed64InKey.String(), Other: "Must be one of {{.Rule}}"},
	validateSFixed64NotInKey: {ID: validateSFixed64NotInKey.String(), Other: "Must not be one of {{.Rule}}"},

	// Bytes validation messages
	validateBytesConstKey:    {ID: validateBytesConstKey.String(), Other: "Must be exactly '{{.Rule}}'"},
	validateBytesLenKey:      {ID: validateBytesLenKey.String(), Other: "Length must be exactly {{.Rule}} bytes"},
	validateBytesMinLenKey:   {ID: validateBytesMinLenKey.String(), Other: "Length must be at least {{.Rule}} bytes"},
	validateBytesMaxLenKey:   {ID: validateBytesMaxLenKey.String(), Other: "Length must be at most {{.Rule}} bytes"},
	validateBytesPatternKey:  {ID: validateBytesPatternKey.String(), Other: "Must match the pattern '{{.Rule}}'"},
	validateBytesPrefixKey:   {ID: validateBytesPrefixKey.String(), Other: "Must start with '{{.Rule}}'"},
	validateBytesSuffixKey:   {ID: validateBytesSuffixKey.String(), Other: "Must end with '{{.Rule}}'"},
	validateBytesContainsKey: {ID: validateBytesContainsKey.String(), Other: "Must contain '{{.Rule}}'"},
	validateBytesInKey:       {ID: validateBytesInKey.String(), Other: "Must be one of [{{.Rule}}]"},
	validateBytesNotInKey:    {ID: validateBytesNotInKey.String(), Other: "Must not be one of [{{.Rule}}]"},

	// Message validation messages
	validateMessageRequiredKey: {ID: validateMessageRequiredKey.String(), Other: "Message is required"},
	validateMessageSkipKey:     {ID: validateMessageSkipKey.String(), Other: "Message is skipped"},

	// Any validation messages
	validateAnyRequiredKey: {ID: validateAnyRequiredKey.String(), Other: "Any is required"},
	validateAnyInKey:       {ID: validateAnyInKey.String(), Other: "Any must be one of [{{.Rule}}]"},
	validateAnyNotInKey:    {ID: validateAnyNotInKey.String(), Other: "Any must not be one of [{{.Rule}}]"},

	// Duration validation messages
	validateDurationRequiredKey: {ID: validateDurationRequiredKey.String(), Other: "Duration is required"},
	validateDurationConstKey:    {ID: validateDurationConstKey.String(), Other: "Duration must be '{{.Rule}}'"},
	validateDurationLtKey:       {ID: validateDurationLtKey.String(), Other: "Duration must be less than {{.Rule}}"},
	validateDurationLteKey:      {ID: validateDurationLteKey.String(), Other: "Duration must be less than or equal to {{.Rule}}"},
	validateDurationGtKey:       {ID: validateDurationGtKey.String(), Other: "Duration must be greater than {{.Rule}}"},
	validateDurationGteKey:      {ID: validateDurationGteKey.String(), Other: "Duration must be greater than or equal to {{.Rule}}"},
	validateDurationInKey:       {ID: validateDurationInKey.String(), Other: "Duration must be one of [{{.Rule}}]"},
	validateDurationNotInKey:    {ID: validateDurationNotInKey.String(), Other: "Duration must not be one of [{{.Rule}}]"},

	// Timestamp validation messages
	validateTimestampRequiredKey: {ID: validateTimestampRequiredKey.String(), Other: "Timestamp is required"},
	validateTimestampConstKey:    {ID: validateTimestampConstKey.String(), Other: "Timestamp must be '{{.Rule}}'"},
	validateTimestampLtKey:       {ID: validateTimestampLtKey.String(), Other: "Timestamp must be less than {{.Rule}}"},
	validateTimestampLteKey:      {ID: validateTimestampLteKey.String(), Other: "Timestamp must be less than or equal to {{.Rule}}"},
	validateTimestampGtKey:       {ID: validateTimestampGtKey.String(), Other: "Timestamp must be greater than {{.Rule}}"},
	validateTimestampGteKey:      {ID: validateTimestampGteKey.String(), Other: "Timestamp must be greater than or equal to {{.Rule}}"},
	validateTimestampLtNowKey:    {ID: validateTimestampLtNowKey.String(), Other: "Timestamp must be less than the current time"},
	validateTimestampGtNowKey:    {ID: validateTimestampGtNowKey.String(), Other: "Timestamp must be greater than the current time"},
	validateTimestampWithinKey:   {ID: validateTimestampWithinKey.String(), Other: "Timestamp must be within the time range"},
}

var zhMsgMap = map[MsgKey]*i18n.Message{
	parameterErrorKey:            {ID: parameterErrorKey.String(), Other: "参数错误,请检查您的参数"},
	getMetadataFailKey:           {ID: getMetadataFailKey.String(), Other: "获取元数据失败"},
	getMetadataConversionFailKey: {ID: getMetadataConversionFailKey.String(), Other: "元数据 {{.Source}} 转换为 {{.Target}} 失败"},
	missingMetadataKey:           {ID: missingMetadataKey.String(), Other: "缺少元数据 {{.Name}},请检查传输链路是否启用元数据传递,并传值。"},
	encryptFailKey:               {ID: encryptFailKey.String(), Other: "加密失败"},
	decryptFailKey:               {ID: decryptFailKey.String(), Other: "解密失败"},
	paramCanNotEmptyKey:          {ID: paramCanNotEmptyKey.String(), Other: "参数 {{.Name}} 不能为空"},
	notEditableKey:               {ID: notEditableKey.String(), Other: "数据不可编辑"},
	requestErrorKey:              {ID: requestErrorKey.String(), Other: "请求出错了"},
	pleaseLoginKey:               {ID: pleaseLoginKey.String(), Other: "请先登录"},
	needReLoginKey:               {ID: needReLoginKey.String(), Other: "需要重新登录"},
	accountDisabledKey:           {ID: accountDisabledKey.String(), Other: "账号已被禁用"},
	pleaseChangePasswordKey:      {ID: pleaseChangePasswordKey.String(), Other: "请先修改密码"},
	passwordErrorKey:             {ID: passwordErrorKey.String(), Other: "密码错误"},
	noPermissionKey:              {ID: noPermissionKey.String(), Other: "没有权限"},
	notConfiguredKey:             {ID: notConfiguredKey.String(), Other: "{{.Name}} 未配置"},
	canNotEmptyKey:               {ID: canNotEmptyKey.String(), Other: "{{.Name}} 不能为空"},
	expiredKey:                   {ID: expiredKey.String(), Other: "{{.Name}} 已过期"},

	dataAbnormalKey:        {ID: dataAbnormalKey.String(), Other: "{{.Name}} 数据异常"},
	dataNotFoundKey:        {ID: dataNotFoundKey.String(), Other: "找不到数据"},
	dataDuplicateKey:       {ID: dataDuplicateKey.String(), Other: "该数据已存在,请勿重复添加"},
	dataConstraintKey:      {ID: dataConstraintKey.String(), Other: "数据约束检查失败，请检查您的参数"},
	dataNotLoadedKey:       {ID: dataNotLoadedKey.String(), Other: "数据库未加载，请联系管理员"},
	dataNotSingularKey:     {ID: dataNotSingularKey.String(), Other: "数据出错了 Not Singular,请联系管理员"},
	dataValidationErrorKey: {ID: dataValidationErrorKey.String(), Other: "数据校验失败，请检查您的参数"},
	dataErrorKey:           {ID: dataErrorKey.String(), Other: "数据层出错了,请联系管理员"},

	cacheNotFoundKey:        {ID: cacheNotFoundKey.String(), Other: "缓存不存在"},
	cachePreMatchGetFailKey: {ID: cachePreMatchGetFailKey.String(), Other: "前置匹配获取缓存失败"},
	cacheSetFailKey:         {ID: cacheSetFailKey.String(), Other: "设置缓存失败"},
	cacheMSetFailKey:        {ID: cacheMSetFailKey.String(), Other: "批量设置缓存失败"},
	cacheDelFailKey:         {ID: cacheDelFailKey.String(), Other: "删除缓存失败"},
	cachePreMatchDelFailKey: {ID: cachePreMatchDelFailKey.String(), Other: "前置匹配删除缓存失败"},
	cacheFlushFailKey:       {ID: cacheFlushFailKey.String(), Other: "清空缓存失败"},
	cacheMGetFailKey:        {ID: cacheMGetFailKey.String(), Other: "批量获取缓存失败"},
	cacheMDelFailKey:        {ID: cacheMDelFailKey.String(), Other: "批量删除缓存失败"},

	validateRequiredKey: {ID: validateRequiredKey.String(), Other: "该字段是必填的"},
	// String validation messages
	validateStringConstKey:             {ID: validateStringConstKey.String(), Other: "必须等于 '{{.Rule}}'"},
	validateStringLenKey:               {ID: validateStringLenKey.String(), Other: "长度必须等于 {{.Rule}} 个字符"},
	validateStringMinLenKey:            {ID: validateStringMinLenKey.String(), Other: "长度必须至少为 {{.Rule}} 个字符"},
	validateStringMaxLenKey:            {ID: validateStringMaxLenKey.String(), Other: "长度必须最多为 {{.Rule}} 个字符"},
	validateStringPatternKey:           {ID: validateStringPatternKey.String(), Other: "必须匹配模式 '{{.Rule}}'"},
	validateStringPrefixKey:            {ID: validateStringPrefixKey.String(), Other: "必须以 '{{.Rule}}' 开头"},
	validateStringSuffixKey:            {ID: validateStringSuffixKey.String(), Other: "必须以 '{{.Rule}}' 结尾"},
	validateStringContainsKey:          {ID: validateStringContainsKey.String(), Other: "必须包含 '{{.Rule}}'"},
	validateStringNotContainsKey:       {ID: validateStringNotContainsKey.String(), Other: "不能包含 '{{.Rule}}'"},
	validateStringInKey:                {ID: validateStringInKey.String(), Other: "必须是 [{{.Rule}}] 其中之一"},
	validateStringNotInKey:             {ID: validateStringNotInKey.String(), Other: "不能是 [{{.Rule}}] 其中之一"},
	validateStringEmailKey:             {ID: validateStringEmailKey.String(), Other: "必须是有效的电子邮件地址"},
	validateStringHostnameKey:          {ID: validateStringHostnameKey.String(), Other: "必须是有效的主机名"},
	validateStringIPKey:                {ID: validateStringIPKey.String(), Other: "必须是有效的IP地址"},
	validateStringIPv4Key:              {ID: validateStringIPv4Key.String(), Other: "必须是有效的IPv4地址"},
	validateStringIPv6Key:              {ID: validateStringIPv6Key.String(), Other: "必须是有效的IPv6地址"},
	validateStringURIKey:               {ID: validateStringURIKey.String(), Other: "必须是有效的URI"},
	validateStringURIRefKey:            {ID: validateStringURIRefKey.String(), Other: "必须是有效的URI引用"},
	validateStringAddressKey:           {ID: validateStringAddressKey.String(), Other: "必须是有效的地址"},
	validateStringUUIDKey:              {ID: validateStringUUIDKey.String(), Other: "必须是有效的UUID"},
	validateStringWellKnownRegexKey:    {ID: validateStringWellKnownRegexKey.String(), Other: "必须匹配格式 {{.Rule}}"},
	validateStringLenBytesKey:          {ID: validateStringLenBytesKey.String(), Other: "必须是 {{.Rule}} 字节"},
	validateStringMinBytesKey:          {ID: validateStringMinBytesKey.String(), Other: "必须至少 {{.Rule}} 字节"},
	validateStringMaxBytesKey:          {ID: validateStringMaxBytesKey.String(), Other: "必须最多 {{.Rule}} 字节"},
	validateStringTUUIDKey:             {ID: validateStringTUUIDKey.String(), Other: "必须是有效的基于时间的UUID"},
	validateStringIPPrefixKey:          {ID: validateStringIPPrefixKey.String(), Other: "必须是有效的IP网络前缀"},
	validateStringIPWithPrefixlenKey:   {ID: validateStringIPWithPrefixlenKey.String(), Other: "必须是带有前缀长度的有效IP地址"},
	validateStringIPv4WithPrefixlenKey: {ID: validateStringIPv4WithPrefixlenKey.String(), Other: "必须是带有前缀长度的有效IPv4地址"},
	validateStringIPv6WithPrefixlenKey: {ID: validateStringIPv6WithPrefixlenKey.String(), Other: "必须是带有前缀长度的有效IPv6地址"},
	validateStringStrictKey:            {ID: validateStringStrictKey.String(), Other: "必须符合严格验证规则"},
	validateStringExampleKey:           {ID: validateStringExampleKey.String(), Other: "必须匹配示例格式: {{.Rule}}"},

	// Boolean validation messages
	validateBoolConstKey: {ID: validateBoolConstKey.String(), Other: "必须为 {{.Rule}}"},

	// Enum validation messages
	validateEnumConstKey:   {ID: validateEnumConstKey.String(), Other: "必须等于 {{.Rule}}"},
	validateEnumDefinedKey: {ID: validateEnumDefinedKey.String(), Other: "必须是已定义的枚举值"},
	validateEnumInKey:      {ID: validateEnumInKey.String(), Other: "必须是 {{.Rule}} 其中之一"},
	validateEnumNotInKey:   {ID: validateEnumNotInKey.String(), Other: "不能是 {{.Rule}} 其中之一"},

	// Repeated field validation messages
	validateRepeatedMinItemsKey: {ID: validateRepeatedMinItemsKey.String(), Other: "至少需要 {{.Rule}} 个元素"},
	validateRepeatedMaxItemsKey: {ID: validateRepeatedMaxItemsKey.String(), Other: "最多只能有 {{.Rule}} 个元素"},
	validateRepeatedUniqueKey:   {ID: validateRepeatedUniqueKey.String(), Other: "所有元素必须唯一"},
	validateRepeatedItemsKey:    {ID: validateRepeatedItemsKey.String(), Other: "元素必须满足规则"},

	// Map validation messages
	validateMapMinPairsKey: {ID: validateMapMinPairsKey.String(), Other: "至少需要 {{.Rule}} 个键值对"},
	validateMapMaxPairsKey: {ID: validateMapMaxPairsKey.String(), Other: "最多只能有 {{.Rule}} 个键值对"},
	validateMapNoSparseKey: {ID: validateMapNoSparseKey.String(), Other: "不能包含空值"},
	validateMapKeysKey:     {ID: validateMapKeysKey.String(), Other: "键必须满足规则"},
	validateMapValuesKey:   {ID: validateMapValuesKey.String(), Other: "值必须满足规则"},

	// Float validation messages
	validateFloatConstKey:  {ID: validateFloatConstKey.String(), Other: "必须等于 {{.Rule}}"},
	validateFloatLtKey:     {ID: validateFloatLtKey.String(), Other: "必须小于 {{.Rule}}"},
	validateFloatLteKey:    {ID: validateFloatLteKey.String(), Other: "必须小于或等于 {{.Rule}}"},
	validateFloatGtKey:     {ID: validateFloatGtKey.String(), Other: "必须大于 {{.Rule}}"},
	validateFloatGteKey:    {ID: validateFloatGteKey.String(), Other: "必须大于或等于 {{.Rule}}"},
	validateFloatInKey:     {ID: validateFloatInKey.String(), Other: "必须是 {{.Rule}} 其中之一"},
	validateFloatNotInKey:  {ID: validateFloatNotInKey.String(), Other: "不能是 {{.Rule}} 其中之一"},
	validateFloatFiniteKey: {ID: validateFloatFiniteKey.String(), Other: "必须是有限数"},

	// Double validation messages
	validateDoubleConstKey:  {ID: validateDoubleConstKey.String(), Other: "必须等于 {{.Rule}}"},
	validateDoubleLtKey:     {ID: validateDoubleLtKey.String(), Other: "必须小于 {{.Rule}}"},
	validateDoubleLteKey:    {ID: validateDoubleLteKey.String(), Other: "必须小于或等于 {{.Rule}}"},
	validateDoubleGtKey:     {ID: validateDoubleGtKey.String(), Other: "必须大于 {{.Rule}}"},
	validateDoubleGteKey:    {ID: validateDoubleGteKey.String(), Other: "必须大于或等于 {{.Rule}}"},
	validateDoubleInKey:     {ID: validateDoubleInKey.String(), Other: "必须是 {{.Rule}} 其中之一"},
	validateDoubleNotInKey:  {ID: validateDoubleNotInKey.String(), Other: "不能是 {{.Rule}} 其中之一"},
	validateDoubleFiniteKey: {ID: validateDoubleFiniteKey.String(), Other: "必须是有限数"},

	// Int32 validation messages
	validateInt32ConstKey: {ID: validateInt32ConstKey.String(), Other: "必须等于 {{.Rule}}"},
	validateInt32LtKey:    {ID: validateInt32LtKey.String(), Other: "必须小于 {{.Rule}}"},
	validateInt32LteKey:   {ID: validateInt32LteKey.String(), Other: "必须小于或等于 {{.Rule}}"},
	validateInt32GtKey:    {ID: validateInt32GtKey.String(), Other: "必须大于 {{.Rule}}"},
	validateInt32GteKey:   {ID: validateInt32GteKey.String(), Other: "必须大于或等于 {{.Rule}}"},
	validateInt32InKey:    {ID: validateInt32InKey.String(), Other: "必须是 {{.Rule}} 其中之一"},
	validateInt32NotInKey: {ID: validateInt32NotInKey.String(), Other: "不能是 {{.Rule}} 其中之一"},

	// Int64 validation messages
	validateInt64ConstKey: {ID: validateInt64ConstKey.String(), Other: "必须等于 {{.Rule}}"},
	validateInt64LtKey:    {ID: validateInt64LtKey.String(), Other: "必须小于 {{.Rule}}"},
	validateInt64LteKey:   {ID: validateInt64LteKey.String(), Other: "必须小于或等于 {{.Rule}}"},
	validateInt64GtKey:    {ID: validateInt64GtKey.String(), Other: "必须大于 {{.Rule}}"},
	validateInt64GteKey:   {ID: validateInt64GteKey.String(), Other: "必须大于或等于 {{.Rule}}"},
	validateInt64InKey:    {ID: validateInt64InKey.String(), Other: "必须是 {{.Rule}} 其中之一"},
	validateInt64NotInKey: {ID: validateInt64NotInKey.String(), Other: "不能是 {{.Rule}} 其中之一"},

	// UInt32 validation messages
	validateUInt32ConstKey: {ID: validateUInt32ConstKey.String(), Other: "必须等于 {{.Rule}}"},
	validateUInt32LtKey:    {ID: validateUInt32LtKey.String(), Other: "必须小于 {{.Rule}}"},
	validateUInt32LteKey:   {ID: validateUInt32LteKey.String(), Other: "必须小于或等于 {{.Rule}}"},
	validateUInt32GtKey:    {ID: validateUInt32GtKey.String(), Other: "必须大于 {{.Rule}}"},
	validateUInt32GteKey:   {ID: validateUInt32GteKey.String(), Other: "必须大于或等于 {{.Rule}}"},
	validateUInt32InKey:    {ID: validateUInt32InKey.String(), Other: "必须是 {{.Rule}} 其中之一"},
	validateUInt32NotInKey: {ID: validateUInt32NotInKey.String(), Other: "不能是 {{.Rule}} 其中之一"},

	// UInt64 validation messages
	validateUInt64ConstKey: {ID: validateUInt64ConstKey.String(), Other: "必须等于 {{.Rule}}"},
	validateUInt64LtKey:    {ID: validateUInt64LtKey.String(), Other: "必须小于 {{.Rule}}"},
	validateUInt64LteKey:   {ID: validateUInt64LteKey.String(), Other: "必须小于或等于 {{.Rule}}"},
	validateUInt64GtKey:    {ID: validateUInt64GtKey.String(), Other: "必须大于 {{.Rule}}"},
	validateUInt64GteKey:   {ID: validateUInt64GteKey.String(), Other: "必须大于或等于 {{.Rule}}"},
	validateUInt64InKey:    {ID: validateUInt64InKey.String(), Other: "必须是 {{.Rule}} 其中之一"},
	validateUInt64NotInKey: {ID: validateUInt64NotInKey.String(), Other: "不能是 {{.Rule}} 其中之一"},

	// SInt32 validation messages
	validateSInt32ConstKey: {ID: validateSInt32ConstKey.String(), Other: "必须等于 {{.Rule}}"},
	validateSInt32LtKey:    {ID: validateSInt32LtKey.String(), Other: "必须小于 {{.Rule}}"},
	validateSInt32LteKey:   {ID: validateSInt32LteKey.String(), Other: "必须小于或等于 {{.Rule}}"},
	validateSInt32GtKey:    {ID: validateSInt32GtKey.String(), Other: "必须大于 {{.Rule}}"},
	validateSInt32GteKey:   {ID: validateSInt32GteKey.String(), Other: "必须大于或等于 {{.Rule}}"},
	validateSInt32InKey:    {ID: validateSInt32InKey.String(), Other: "必须是 {{.Rule}} 其中之一"},
	validateSInt32NotInKey: {ID: validateSInt32NotInKey.String(), Other: "不能是 {{.Rule}} 其中之一"},

	// SInt64 validation messages
	validateSInt64ConstKey: {ID: validateSInt64ConstKey.String(), Other: "必须等于 {{.Rule}}"},
	validateSInt64LtKey:    {ID: validateSInt64LtKey.String(), Other: "必须小于 {{.Rule}}"},
	validateSInt64LteKey:   {ID: validateSInt64LteKey.String(), Other: "必须小于或等于 {{.Rule}}"},
	validateSInt64GtKey:    {ID: validateSInt64GtKey.String(), Other: "必须大于 {{.Rule}}"},
	validateSInt64GteKey:   {ID: validateSInt64GteKey.String(), Other: "必须大于或等于 {{.Rule}}"},
	validateSInt64InKey:    {ID: validateSInt64InKey.String(), Other: "必须是 {{.Rule}} 其中之一"},
	validateSInt64NotInKey: {ID: validateSInt64NotInKey.String(), Other: "不能是 {{.Rule}} 其中之一"},

	// Fixed32 validation messages
	validateFixed32ConstKey: {ID: validateFixed32ConstKey.String(), Other: "必须等于 {{.Rule}}"},
	validateFixed32LtKey:    {ID: validateFixed32LtKey.String(), Other: "必须小于 {{.Rule}}"},
	validateFixed32LteKey:   {ID: validateFixed32LteKey.String(), Other: "必须小于或等于 {{.Rule}}"},
	validateFixed32GtKey:    {ID: validateFixed32GtKey.String(), Other: "必须大于 {{.Rule}}"},
	validateFixed32GteKey:   {ID: validateFixed32GteKey.String(), Other: "必须大于或等于 {{.Rule}}"},
	validateFixed32InKey:    {ID: validateFixed32InKey.String(), Other: "必须是 {{.Rule}} 其中之一"},
	validateFixed32NotInKey: {ID: validateFixed32NotInKey.String(), Other: "不能是 {{.Rule}} 其中之一"},

	// Fixed64 validation messages
	validateFixed64ConstKey: {ID: validateFixed64ConstKey.String(), Other: "必须等于 {{.Rule}}"},
	validateFixed64LtKey:    {ID: validateFixed64LtKey.String(), Other: "必须小于 {{.Rule}}"},
	validateFixed64LteKey:   {ID: validateFixed64LteKey.String(), Other: "必须小于或等于 {{.Rule}}"},
	validateFixed64GtKey:    {ID: validateFixed64GtKey.String(), Other: "必须大于 {{.Rule}}"},
	validateFixed64GteKey:   {ID: validateFixed64GteKey.String(), Other: "必须大于或等于 {{.Rule}}"},
	validateFixed64InKey:    {ID: validateFixed64InKey.String(), Other: "必须是 {{.Rule}} 其中之一"},
	validateFixed64NotInKey: {ID: validateFixed64NotInKey.String(), Other: "不能是 {{.Rule}} 其中之一"},

	// SFixed32 validation messages
	validateSFixed32ConstKey: {ID: validateSFixed32ConstKey.String(), Other: "必须等于 {{.Rule}}"},
	validateSFixed32LtKey:    {ID: validateSFixed32LtKey.String(), Other: "必须小于 {{.Rule}}"},
	validateSFixed32LteKey:   {ID: validateSFixed32LteKey.String(), Other: "必须小于或等于 {{.Rule}}"},
	validateSFixed32GtKey:    {ID: validateSFixed32GtKey.String(), Other: "必须大于 {{.Rule}}"},
	validateSFixed32GteKey:   {ID: validateSFixed32GteKey.String(), Other: "必须大于或等于 {{.Rule}}"},
	validateSFixed32InKey:    {ID: validateSFixed32InKey.String(), Other: "必须是 {{.Rule}} 其中之一"},
	validateSFixed32NotInKey: {ID: validateSFixed32NotInKey.String(), Other: "不能是 {{.Rule}} 其中之一"},

	// SFixed64 validation messages
	validateSFixed64ConstKey: {ID: validateSFixed64ConstKey.String(), Other: "必须等于 {{.Rule}}"},
	validateSFixed64LtKey:    {ID: validateSFixed64LtKey.String(), Other: "必须小于 {{.Rule}}"},
	validateSFixed64LteKey:   {ID: validateSFixed64LteKey.String(), Other: "必须小于或等于 {{.Rule}}"},
	validateSFixed64GtKey:    {ID: validateSFixed64GtKey.String(), Other: "必须大于 {{.Rule}}"},
	validateSFixed64GteKey:   {ID: validateSFixed64GteKey.String(), Other: "必须大于或等于 {{.Rule}}"},
	validateSFixed64InKey:    {ID: validateSFixed64InKey.String(), Other: "必须是 {{.Rule}} 其中之一"},
	validateSFixed64NotInKey: {ID: validateSFixed64NotInKey.String(), Other: "不能是 {{.Rule}} 其中之一"},

	// Bytes validation messages
	validateBytesConstKey:    {ID: validateBytesConstKey.String(), Other: "必须等于 '{{.Rule}}'"},
	validateBytesLenKey:      {ID: validateBytesLenKey.String(), Other: "长度必须等于 {{.Rule}} 字节"},
	validateBytesMinLenKey:   {ID: validateBytesMinLenKey.String(), Other: "长度必须至少为 {{.Rule}} 字节"},
	validateBytesMaxLenKey:   {ID: validateBytesMaxLenKey.String(), Other: "长度必须最多为 {{.Rule}} 字节"},
	validateBytesPatternKey:  {ID: validateBytesPatternKey.String(), Other: "必须匹配模式 '{{.Rule}}'"},
	validateBytesPrefixKey:   {ID: validateBytesPrefixKey.String(), Other: "必须以 '{{.Rule}}' 开头"},
	validateBytesSuffixKey:   {ID: validateBytesSuffixKey.String(), Other: "必须以 '{{.Rule}}' 结尾"},
	validateBytesContainsKey: {ID: validateBytesContainsKey.String(), Other: "必须包含 '{{.Rule}}'"},
	validateBytesInKey:       {ID: validateBytesInKey.String(), Other: "必须是 [{{.Rule}}] 其中之一"},
	validateBytesNotInKey:    {ID: validateBytesNotInKey.String(), Other: "不能是 [{{.Rule}}] 其中之一"},

	// Message validation messages
	validateMessageRequiredKey: {ID: validateMessageRequiredKey.String(), Other: "消息是必填的"},
	validateMessageSkipKey:     {ID: validateMessageSkipKey.String(), Other: "消息被跳过"},

	// Any validation messages
	validateAnyRequiredKey: {ID: validateAnyRequiredKey.String(), Other: "Any 是必填的"},
	validateAnyInKey:       {ID: validateAnyInKey.String(), Other: "Any 必须是 [{{.Rule}}] 其中之一"},
	validateAnyNotInKey:    {ID: validateAnyNotInKey.String(), Other: "Any 不能是 [{{.Rule}}] 其中之一"},

	// Duration validation messages
	validateDurationRequiredKey: {ID: validateDurationRequiredKey.String(), Other: "Duration 是必填的"},
	validateDurationConstKey:    {ID: validateDurationConstKey.String(), Other: "Duration 必须是 '{{.Rule}}'"},
	validateDurationLtKey:       {ID: validateDurationLtKey.String(), Other: "Duration 必须小于 {{.Rule}}"},
	validateDurationLteKey:      {ID: validateDurationLteKey.String(), Other: "Duration 必须小于或等于 {{.Rule}}"},
	validateDurationGtKey:       {ID: validateDurationGtKey.String(), Other: "Duration 必须大于 {{.Rule}}"},
	validateDurationGteKey:      {ID: validateDurationGteKey.String(), Other: "Duration 必须大于或等于 {{.Rule}}"},
	validateDurationInKey:       {ID: validateDurationInKey.String(), Other: "Duration 必须是 [{{.Rule}}] 其中之一"},
	validateDurationNotInKey:    {ID: validateDurationNotInKey.String(), Other: "Duration 不能是 [{{.Rule}}] 其中之一"},

	// Timestamp validation messages
	validateTimestampRequiredKey: {ID: validateTimestampRequiredKey.String(), Other: "Timestamp 是必填的"},
	validateTimestampConstKey:    {ID: validateTimestampConstKey.String(), Other: "Timestamp 必须是 '{{.Rule}}'"},
	validateTimestampLtKey:       {ID: validateTimestampLtKey.String(), Other: "Timestamp 必须小于 {{.Rule}}"},
	validateTimestampLteKey:      {ID: validateTimestampLteKey.String(), Other: "Timestamp 必须小于或等于 {{.Rule}}"},
	validateTimestampGtKey:       {ID: validateTimestampGtKey.String(), Other: "Timestamp 必须大于 {{.Rule}}"},
	validateTimestampGteKey:      {ID: validateTimestampGteKey.String(), Other: "Timestamp 必须大于或等于 {{.Rule}}"},
	validateTimestampLtNowKey:    {ID: validateTimestampLtNowKey.String(), Other: "Timestamp 必须小于当前时间"},
	validateTimestampGtNowKey:    {ID: validateTimestampGtNowKey.String(), Other: "Timestamp 必须大于当前时间"},
	validateTimestampWithinKey:   {ID: validateTimestampWithinKey.String(), Other: "Timestamp 必须在时间范围内"},
}

var ruMsgMap = map[MsgKey]*i18n.Message{
	parameterErrorKey:            {ID: parameterErrorKey.String(), Other: "Ошибка параметра, пожалуйста, проверьте ваши параметры"},
	getMetadataFailKey:           {ID: getMetadataFailKey.String(), Other: "Ошибка получения метаданных"},
	getMetadataConversionFailKey: {ID: getMetadataConversionFailKey.String(), Other: "Ошибка преобразования метаданных {{.Source}} в {{.Target}}"},
	missingMetadataKey: {
		ID:    missingMetadataKey.String(),
		Other: "Отсутствует метаданные {{.Name}}, пожалуйста, проверьте, включена ли передача метаданных в цепи передачи и передайте значение",
	},
	encryptFailKey:          {ID: encryptFailKey.String(), Other: "Ошибка шифрования"},
	decryptFailKey:          {ID: decryptFailKey.String(), Other: "Ошибка дешифрования"},
	paramCanNotEmptyKey:     {ID: paramCanNotEmptyKey.String(), Other: "Параметр {{.Name}} не может быть пустым"},
	notEditableKey:          {ID: notEditableKey.String(), Other: "Данные нельзя редактировать"},
	requestErrorKey:         {ID: requestErrorKey.String(), Other: "Ошибка запроса"},
	pleaseLoginKey:          {ID: pleaseLoginKey.String(), Other: "Пожалуйста, сначала войдите"},
	needReLoginKey:          {ID: needReLoginKey.String(), Other: "Нужно перелогиниться"},
	accountDisabledKey:      {ID: accountDisabledKey.String(), Other: "Учетная запись отключена"},
	pleaseChangePasswordKey: {ID: pleaseChangePasswordKey.String(), Other: "Пожалуйста, сначала измените пароль"},
	passwordErrorKey:        {ID: passwordErrorKey.String(), Other: "Ошибка пар оля"},
	noPermissionKey:         {ID: noPermissionKey.String(), Other: "Нет разрешения"},
	notConfiguredKey:        {ID: notConfiguredKey.String(), Other: "{{.Name}} не настроено"},
	canNotEmptyKey:          {ID: canNotEmptyKey.String(), Other: "{{.Name}} не может быть пустым"},
	expiredKey:              {ID: expiredKey.String(), Other: "{{.Name}} истек срок действия"},

	dataAbnormalKey:         {ID: dataAbnormalKey.String(), Other: "{{.Name}} данные аномальные"},
	dataNotFoundKey:         {ID: dataNotFoundKey.String(), Other: "Данные не найдены"},
	dataDuplicateKey:        {ID: dataDuplicateKey.String(), Other: "Эти данные уже существуют, пожалуйста, не добавляйте их повторно"},
	dataConstraintKey:       {ID: dataConstraintKey.String(), Other: "Ошибка проверки ограничений данных, пожалуйста, проверьте ваши параметры"},
	dataNotLoadedKey:        {ID: dataNotLoadedKey.String(), Other: "База данных не загружена, пожалуйста, свяжитесь с администратором"},
	dataNotSingularKey:      {ID: dataNotSingularKey.String(), Other: "Ошибка данных Not Singular, пожалуйста, свяжитесь с администратором"},
	dataValidationErrorKey:  {ID: dataValidationErrorKey.String(), Other: "Ошибка проверки данных, пожалуйста, проверьте ваши параметры"},
	dataErrorKey:            {ID: dataErrorKey.String(), Other: "Ошибка слоя данных, пожалуйста, свяжитесь с администратором"},
	cacheNotFoundKey:        {ID: cacheNotFoundKey.String(), Other: "Кэш не найден"},
	cachePreMatchGetFailKey: {ID: cachePreMatchGetFailKey.String(), Other: "Предварительное сопоставление получения кэша не удалось"},
	cacheSetFailKey:         {ID: cacheSetFailKey.String(), Other: "Ошибка установки кэша"},
	cacheMSetFailKey:        {ID: cacheMSetFailKey.String(), Other: "Ошибка установки кэша"},
	cacheDelFailKey:         {ID: cacheDelFailKey.String(), Other: "Ошибка удаления кэша"},
	cachePreMatchDelFailKey: {ID: cachePreMatchDelFailKey.String(), Other: "Предварительное сопоставление удаления кэша не удалось"},
	cacheFlushFailKey:       {ID: cacheFlushFailKey.String(), Other: "Ошибка очистки кэша"},
	cacheMGetFailKey:        {ID: cacheMGetFailKey.String(), Other: "Ошибка получения кэша"},
	cacheMDelFailKey:        {ID: cacheMDelFailKey.String(), Other: "Ошибка удаления кэша"},

	validateRequiredKey: {ID: validateRequiredKey.String(), Other: "Это поле обязательное"},
	// String validation messages
	validateStringConstKey:             {ID: validateStringConstKey.String(), Other: "Должно быть равно '{{.Rule}}'"},
	validateStringLenKey:               {ID: validateStringLenKey.String(), Other: "Длина должна быть равна {{.Rule}} символам"},
	validateStringMinLenKey:            {ID: validateStringMinLenKey.String(), Other: "Длина должна быть не менее {{.Rule}} символов"},
	validateStringMaxLenKey:            {ID: validateStringMaxLenKey.String(), Other: "Длина должна быть не более {{.Rule}} символов"},
	validateStringPatternKey:           {ID: validateStringPatternKey.String(), Other: "Должно соответствовать шаблону '{{.Rule}}'"},
	validateStringPrefixKey:            {ID: validateStringPrefixKey.String(), Other: "Должно начинаться с '{{.Rule}}'"},
	validateStringSuffixKey:            {ID: validateStringSuffixKey.String(), Other: "Должно заканчиваться на '{{.Rule}}'"},
	validateStringContainsKey:          {ID: validateStringContainsKey.String(), Other: "Должно содержать '{{.Rule}}'"},
	validateStringNotContainsKey:       {ID: validateStringNotContainsKey.String(), Other: "Не должно содержать '{{.Rule}}'"},
	validateStringInKey:                {ID: validateStringInKey.String(), Other: "Должно быть одним из [{{.Rule}}]"},
	validateStringNotInKey:             {ID: validateStringNotInKey.String(), Other: "Не должно быть одним из [{{.Rule}}]"},
	validateStringEmailKey:             {ID: validateStringEmailKey.String(), Other: "Должно быть действительным email адресом"},
	validateStringHostnameKey:          {ID: validateStringHostnameKey.String(), Other: "Должно быть действительным именем хоста"},
	validateStringIPKey:                {ID: validateStringIPKey.String(), Other: "Должно быть действительным IP адресом"},
	validateStringIPv4Key:              {ID: validateStringIPv4Key.String(), Other: "Должно быть действительным IPv4 адресом"},
	validateStringIPv6Key:              {ID: validateStringIPv6Key.String(), Other: "Должно быть действительным IPv6 адресом"},
	validateStringURIKey:               {ID: validateStringURIKey.String(), Other: "Должно быть действительным URI"},
	validateStringURIRefKey:            {ID: validateStringURIRefKey.String(), Other: "Должно быть действительной ссылкой URI"},
	validateStringAddressKey:           {ID: validateStringAddressKey.String(), Other: "Должно быть действительным адресом"},
	validateStringUUIDKey:              {ID: validateStringUUIDKey.String(), Other: "Должно быть действительным UUID"},
	validateStringWellKnownRegexKey:    {ID: validateStringWellKnownRegexKey.String(), Other: "Должно соответствовать формату {{.Rule}}"},
	validateStringLenBytesKey:          {ID: validateStringLenBytesKey.String(), Other: "Должно быть ровно {{.Rule}} байт"},
	validateStringMinBytesKey:          {ID: validateStringMinBytesKey.String(), Other: "Должно быть не менее {{.Rule}} байт"},
	validateStringMaxBytesKey:          {ID: validateStringMaxBytesKey.String(), Other: "Должно быть не более {{.Rule}} байт"},
	validateStringTUUIDKey:             {ID: validateStringTUUIDKey.String(), Other: "Должно быть действительным UUID на основе времени"},
	validateStringIPPrefixKey:          {ID: validateStringIPPrefixKey.String(), Other: "Должно быть действительным префиксом сети IP"},
	validateStringIPWithPrefixlenKey:   {ID: validateStringIPWithPrefixlenKey.String(), Other: "Должно быть действительным IP адресом с длиной префикса"},
	validateStringIPv4WithPrefixlenKey: {ID: validateStringIPv4WithPrefixlenKey.String(), Other: "Должно быть действительным IPv4 адресом с длиной префикса"},
	validateStringIPv6WithPrefixlenKey: {ID: validateStringIPv6WithPrefixlenKey.String(), Other: "Должно быть действительным IPv6 адресом с длиной префикса"},
	validateStringStrictKey:            {ID: validateStringStrictKey.String(), Other: "Должно соответствовать строгим правилам проверки"},
	validateStringExampleKey:           {ID: validateStringExampleKey.String(), Other: "Должно соответствовать формату примера: {{.Rule}}"},

	// Boolean validation messages
	validateBoolConstKey: {ID: validateBoolConstKey.String(), Other: "Должно быть {{.Rule}}"},

	// Enum validation messages
	validateEnumConstKey:   {ID: validateEnumConstKey.String(), Other: "Должно быть равно {{.Rule}}"},
	validateEnumDefinedKey: {ID: validateEnumDefinedKey.String(), Other: "Должно быть определенным значением перечисления"},
	validateEnumInKey:      {ID: validateEnumInKey.String(), Other: "Должно быть одним из {{.Rule}}"},
	validateEnumNotInKey:   {ID: validateEnumNotInKey.String(), Other: "Не должно быть одним из {{.Rule}}"},

	// Repeated field validation messages
	validateRepeatedMinItemsKey: {ID: validateRepeatedMinItemsKey.String(), Other: "Должно содержать не менее {{.Rule}} элементов"},
	validateRepeatedMaxItemsKey: {ID: validateRepeatedMaxItemsKey.String(), Other: "Должно содержать не более {{.Rule}} элементов"},
	validateRepeatedUniqueKey:   {ID: validateRepeatedUniqueKey.String(), Other: "Все элементы должны быть уникальными"},
	validateRepeatedItemsKey:    {ID: validateRepeatedItemsKey.String(), Other: "Элементы должны соответствовать правилам"},

	// Map validation messages
	validateMapMinPairsKey: {ID: validateMapMinPairsKey.String(), Other: "Должно содержать не менее {{.Rule}} пар ключ-значение"},
	validateMapMaxPairsKey: {ID: validateMapMaxPairsKey.String(), Other: "Должно содержать не более {{.Rule}} пар ключ-значение"},
	validateMapNoSparseKey: {ID: validateMapNoSparseKey.String(), Other: "Не должно содержать пустых значений"},
	validateMapKeysKey:     {ID: validateMapKeysKey.String(), Other: "Ключи должны соответствовать правилам"},
	validateMapValuesKey:   {ID: validateMapValuesKey.String(), Other: "Значения должны соответствовать правилам"},

	// Float validation messages
	validateFloatConstKey:  {ID: validateFloatConstKey.String(), Other: "Должно быть равно {{.Rule}}"},
	validateFloatLtKey:     {ID: validateFloatLtKey.String(), Other: "Должно быть меньше {{.Rule}}"},
	validateFloatLteKey:    {ID: validateFloatLteKey.String(), Other: "Должно быть меньше или равно {{.Rule}}"},
	validateFloatGtKey:     {ID: validateFloatGtKey.String(), Other: "Должно быть больше {{.Rule}}"},
	validateFloatGteKey:    {ID: validateFloatGteKey.String(), Other: "Должно быть больше или равно {{.Rule}}"},
	validateFloatInKey:     {ID: validateFloatInKey.String(), Other: "Должно быть одним из {{.Rule}}"},
	validateFloatNotInKey:  {ID: validateFloatNotInKey.String(), Other: "Не должно быть одним из {{.Rule}}"},
	validateFloatFiniteKey: {ID: validateFloatFiniteKey.String(), Other: "Должно быть конечным числом"},

	// Double validation messages
	validateDoubleConstKey:  {ID: validateDoubleConstKey.String(), Other: "Должно быть равно {{.Rule}}"},
	validateDoubleLtKey:     {ID: validateDoubleLtKey.String(), Other: "Должно быть меньше {{.Rule}}"},
	validateDoubleLteKey:    {ID: validateDoubleLteKey.String(), Other: "Должно быть меньше или равно {{.Rule}}"},
	validateDoubleGtKey:     {ID: validateDoubleGtKey.String(), Other: "Должно быть больше {{.Rule}}"},
	validateDoubleGteKey:    {ID: validateDoubleGteKey.String(), Other: "Должно быть больше или равно {{.Rule}}"},
	validateDoubleInKey:     {ID: validateDoubleInKey.String(), Other: "Должно быть одним из {{.Rule}}"},
	validateDoubleNotInKey:  {ID: validateDoubleNotInKey.String(), Other: "Не должно быть одним из {{.Rule}}"},
	validateDoubleFiniteKey: {ID: validateDoubleFiniteKey.String(), Other: "Должно быть конечным числом"},

	// Int32 validation messages
	validateInt32ConstKey: {ID: validateInt32ConstKey.String(), Other: "Должно быть равно {{.Rule}}"},
	validateInt32LtKey:    {ID: validateInt32LtKey.String(), Other: "Должно быть меньше {{.Rule}}"},
	validateInt32LteKey:   {ID: validateInt32LteKey.String(), Other: "Должно быть меньше или равно {{.Rule}}"},
	validateInt32GtKey:    {ID: validateInt32GtKey.String(), Other: "Должно быть больше {{.Rule}}"},
	validateInt32GteKey:   {ID: validateInt32GteKey.String(), Other: "Должно быть больше или равно {{.Rule}}"},
	validateInt32InKey:    {ID: validateInt32InKey.String(), Other: "Должно быть одним из {{.Rule}}"},
	validateInt32NotInKey: {ID: validateInt32NotInKey.String(), Other: "Не должно быть одним из {{.Rule}}"},

	// Int64 validation messages
	validateInt64ConstKey: {ID: validateInt64ConstKey.String(), Other: "Должно быть равно {{.Rule}}"},
	validateInt64LtKey:    {ID: validateInt64LtKey.String(), Other: "Должно быть меньше {{.Rule}}"},
	validateInt64LteKey:   {ID: validateInt64LteKey.String(), Other: "Должно быть меньше или равно {{.Rule}}"},
	validateInt64GtKey:    {ID: validateInt64GtKey.String(), Other: "Должно быть больше {{.Rule}}"},
	validateInt64GteKey:   {ID: validateInt64GteKey.String(), Other: "Должно быть больше или равно {{.Rule}}"},
	validateInt64InKey:    {ID: validateInt64InKey.String(), Other: "Должно быть одним из {{.Rule}}"},
	validateInt64NotInKey: {ID: validateInt64NotInKey.String(), Other: "Не должно быть одним из {{.Rule}}"},

	// UInt32 validation messages
	validateUInt32ConstKey: {ID: validateUInt32ConstKey.String(), Other: "Должно быть равно {{.Rule}}"},
	validateUInt32LtKey:    {ID: validateUInt32LtKey.String(), Other: "Должно быть меньше {{.Rule}}"},
	validateUInt32LteKey:   {ID: validateUInt32LteKey.String(), Other: "Должно быть меньше или равно {{.Rule}}"},
	validateUInt32GtKey:    {ID: validateUInt32GtKey.String(), Other: "Должно быть больше {{.Rule}}"},
	validateUInt32GteKey:   {ID: validateUInt32GteKey.String(), Other: "Должно быть больше или равно {{.Rule}}"},
	validateUInt32InKey:    {ID: validateUInt32InKey.String(), Other: "Должно быть одним из {{.Rule}}"},
	validateUInt32NotInKey: {ID: validateUInt32NotInKey.String(), Other: "Не должно быть одним из {{.Rule}}"},

	// UInt64 validation messages
	validateUInt64ConstKey: {ID: validateUInt64ConstKey.String(), Other: "Должно быть равно {{.Rule}}"},
	validateUInt64LtKey:    {ID: validateUInt64LtKey.String(), Other: "Должно быть меньше {{.Rule}}"},
	validateUInt64LteKey:   {ID: validateUInt64LteKey.String(), Other: "Должно быть меньше или равно {{.Rule}}"},
	validateUInt64GtKey:    {ID: validateUInt64GtKey.String(), Other: "Должно быть больше {{.Rule}}"},
	validateUInt64GteKey:   {ID: validateUInt64GteKey.String(), Other: "Должно быть больше или равно {{.Rule}}"},
	validateUInt64InKey:    {ID: validateUInt64InKey.String(), Other: "Должно быть одним из {{.Rule}}"},
	validateUInt64NotInKey: {ID: validateUInt64NotInKey.String(), Other: "Не должно быть одним из {{.Rule}}"},

	// SInt32 validation messages
	validateSInt32ConstKey: {ID: validateSInt32ConstKey.String(), Other: "Должно быть равно {{.Rule}}"},
	validateSInt32LtKey:    {ID: validateSInt32LtKey.String(), Other: "Должно быть меньше {{.Rule}}"},
	validateSInt32LteKey:   {ID: validateSInt32LteKey.String(), Other: "Должно быть меньше или равно {{.Rule}}"},
	validateSInt32GtKey:    {ID: validateSInt32GtKey.String(), Other: "Должно быть больше {{.Rule}}"},
	validateSInt32GteKey:   {ID: validateSInt32GteKey.String(), Other: "Должно быть больше или равно {{.Rule}}"},
	validateSInt32InKey:    {ID: validateSInt32InKey.String(), Other: "Должно быть одним из {{.Rule}}"},
	validateSInt32NotInKey: {ID: validateSInt32NotInKey.String(), Other: "Не должно быть одним из {{.Rule}}"},

	// SInt64 validation messages
	validateSInt64ConstKey: {ID: validateSInt64ConstKey.String(), Other: "Должно быть равно {{.Rule}}"},
	validateSInt64LtKey:    {ID: validateSInt64LtKey.String(), Other: "Должно быть меньше {{.Rule}}"},
	validateSInt64LteKey:   {ID: validateSInt64LteKey.String(), Other: "Должно быть меньше или равно {{.Rule}}"},
	validateSInt64GtKey:    {ID: validateSInt64GtKey.String(), Other: "Должно быть больше {{.Rule}}"},
	validateSInt64GteKey:   {ID: validateSInt64GteKey.String(), Other: "Должно быть больше или равно {{.Rule}}"},
	validateSInt64InKey:    {ID: validateSInt64InKey.String(), Other: "Должно быть одним из {{.Rule}}"},
	validateSInt64NotInKey: {ID: validateSInt64NotInKey.String(), Other: "Не должно быть одним из {{.Rule}}"},

	// Fixed32 validation messages
	validateFixed32ConstKey: {ID: validateFixed32ConstKey.String(), Other: "Должно быть равно {{.Rule}}"},
	validateFixed32LtKey:    {ID: validateFixed32LtKey.String(), Other: "Должно быть меньше {{.Rule}}"},
	validateFixed32LteKey:   {ID: validateFixed32LteKey.String(), Other: "Должно быть меньше или равно {{.Rule}}"},
	validateFixed32GtKey:    {ID: validateFixed32GtKey.String(), Other: "Должно быть больше {{.Rule}}"},
	validateFixed32GteKey:   {ID: validateFixed32GteKey.String(), Other: "Должно быть больше или равно {{.Rule}}"},
	validateFixed32InKey:    {ID: validateFixed32InKey.String(), Other: "Должно быть одним из {{.Rule}}"},
	validateFixed32NotInKey: {ID: validateFixed32NotInKey.String(), Other: "Не должно быть одним из {{.Rule}}"},

	// Fixed64 validation messages
	validateFixed64ConstKey: {ID: validateFixed64ConstKey.String(), Other: "Должно быть равно {{.Rule}}"},
	validateFixed64LtKey:    {ID: validateFixed64LtKey.String(), Other: "Должно быть меньше {{.Rule}}"},
	validateFixed64LteKey:   {ID: validateFixed64LteKey.String(), Other: "Должно быть меньше или равно {{.Rule}}"},
	validateFixed64GtKey:    {ID: validateFixed64GtKey.String(), Other: "Должно быть больше {{.Rule}}"},
	validateFixed64GteKey:   {ID: validateFixed64GteKey.String(), Other: "Должно быть больше или равно {{.Rule}}"},
	validateFixed64InKey:    {ID: validateFixed64InKey.String(), Other: "Должно быть одним из {{.Rule}}"},
	validateFixed64NotInKey: {ID: validateFixed64NotInKey.String(), Other: "Не должно быть одним из {{.Rule}}"},

	// SFixed32 validation messages
	validateSFixed32ConstKey: {ID: validateSFixed32ConstKey.String(), Other: "Должно быть равно {{.Rule}}"},
	validateSFixed32LtKey:    {ID: validateSFixed32LtKey.String(), Other: "Должно быть меньше {{.Rule}}"},
	validateSFixed32LteKey:   {ID: validateSFixed32LteKey.String(), Other: "Должно быть меньше или равно {{.Rule}}"},
	validateSFixed32GtKey:    {ID: validateSFixed32GtKey.String(), Other: "Должно быть больше {{.Rule}}"},
	validateSFixed32GteKey:   {ID: validateSFixed32GteKey.String(), Other: "Должно быть больше или равно {{.Rule}}"},
	validateSFixed32InKey:    {ID: validateSFixed32InKey.String(), Other: "Должно быть одним из {{.Rule}}"},
	validateSFixed32NotInKey: {ID: validateSFixed32NotInKey.String(), Other: "Не должно быть одним из {{.Rule}}"},

	// SFixed64 validation messages
	validateSFixed64ConstKey: {ID: validateSFixed64ConstKey.String(), Other: "Должно быть равно {{.Rule}}"},
	validateSFixed64LtKey:    {ID: validateSFixed64LtKey.String(), Other: "Должно быть меньше {{.Rule}}"},
	validateSFixed64LteKey:   {ID: validateSFixed64LteKey.String(), Other: "Должно быть меньше или равно {{.Rule}}"},
	validateSFixed64GtKey:    {ID: validateSFixed64GtKey.String(), Other: "Должно быть больше {{.Rule}}"},
	validateSFixed64GteKey:   {ID: validateSFixed64GteKey.String(), Other: "Должно быть больше или равно {{.Rule}}"},
	validateSFixed64InKey:    {ID: validateSFixed64InKey.String(), Other: "Должно быть одним из {{.Rule}}"},
	validateSFixed64NotInKey: {ID: validateSFixed64NotInKey.String(), Other: "Не должно быть одним из {{.Rule}}"},

	// Bytes validation messages
	validateBytesConstKey:    {ID: validateBytesConstKey.String(), Other: "Must be exactly '{{.Rule}}'"},
	validateBytesLenKey:      {ID: validateBytesLenKey.String(), Other: "Length must be exactly {{.Rule}} bytes"},
	validateBytesMinLenKey:   {ID: validateBytesMinLenKey.String(), Other: "Length must be at least {{.Rule}} bytes"},
	validateBytesMaxLenKey:   {ID: validateBytesMaxLenKey.String(), Other: "Length must be at most {{.Rule}} bytes"},
	validateBytesPatternKey:  {ID: validateBytesPatternKey.String(), Other: "Must match the pattern '{{.Rule}}'"},
	validateBytesPrefixKey:   {ID: validateBytesPrefixKey.String(), Other: "Must start with '{{.Rule}}'"},
	validateBytesSuffixKey:   {ID: validateBytesSuffixKey.String(), Other: "Must end with '{{.Rule}}'"},
	validateBytesContainsKey: {ID: validateBytesContainsKey.String(), Other: "Must contain '{{.Rule}}'"},
	validateBytesInKey:       {ID: validateBytesInKey.String(), Other: "Must be one of [{{.Rule}}]"},
	validateBytesNotInKey:    {ID: validateBytesNotInKey.String(), Other: "Must not be one of [{{.Rule}}]"},

	// Message validation messages
	validateMessageRequiredKey: {ID: validateMessageRequiredKey.String(), Other: "Message is required"},
	validateMessageSkipKey:     {ID: validateMessageSkipKey.String(), Other: "Message is skipped"},

	// Any validation messages
	validateAnyRequiredKey: {ID: validateAnyRequiredKey.String(), Other: "Any is required"},
	validateAnyInKey:       {ID: validateAnyInKey.String(), Other: "Any must be one of [{{.Rule}}]"},
	validateAnyNotInKey:    {ID: validateAnyNotInKey.String(), Other: "Any must not be one of [{{.Rule}}]"},

	// Duration validation messages
	validateDurationRequiredKey: {ID: validateDurationRequiredKey.String(), Other: "Duration is required"},
	validateDurationConstKey:    {ID: validateDurationConstKey.String(), Other: "Duration must be '{{.Rule}}'"},
	validateDurationLtKey:       {ID: validateDurationLtKey.String(), Other: "Duration must be less than {{.Rule}}"},
	validateDurationLteKey:      {ID: validateDurationLteKey.String(), Other: "Duration must be less than or equal to {{.Rule}}"},
	validateDurationGtKey:       {ID: validateDurationGtKey.String(), Other: "Duration must be greater than {{.Rule}}"},
	validateDurationGteKey:      {ID: validateDurationGteKey.String(), Other: "Duration must be greater than or equal to {{.Rule}}"},
	validateDurationInKey:       {ID: validateDurationInKey.String(), Other: "Duration must be one of [{{.Rule}}]"},
	validateDurationNotInKey:    {ID: validateDurationNotInKey.String(), Other: "Duration must not be one of [{{.Rule}}]"},

	// Timestamp validation messages
	validateTimestampRequiredKey: {ID: validateTimestampRequiredKey.String(), Other: "Timestamp is required"},
	validateTimestampConstKey:    {ID: validateTimestampConstKey.String(), Other: "Timestamp must be '{{.Rule}}'"},
	validateTimestampLtKey:       {ID: validateTimestampLtKey.String(), Other: "Timestamp must be less than {{.Rule}}"},
	validateTimestampLteKey:      {ID: validateTimestampLteKey.String(), Other: "Timestamp must be less than or equal to {{.Rule}}"},
	validateTimestampGtKey:       {ID: validateTimestampGtKey.String(), Other: "Timestamp must be greater than {{.Rule}}"},
	validateTimestampGteKey:      {ID: validateTimestampGteKey.String(), Other: "Timestamp must be greater than or equal to {{.Rule}}"},
	validateTimestampLtNowKey:    {ID: validateTimestampLtNowKey.String(), Other: "Timestamp must be less than the current time"},
	validateTimestampGtNowKey:    {ID: validateTimestampGtNowKey.String(), Other: "Timestamp must be greater than the current time"},
	validateTimestampWithinKey:   {ID: validateTimestampWithinKey.String(), Other: "Timestamp must be within the time range"},
}

var frMsgMap = map[MsgKey]*i18n.Message{
	parameterErrorKey:            {ID: parameterErrorKey.String(), Other: "Erreur de paramètre, veuillez vérifier vos paramètres"},
	getMetadataFailKey:           {ID: getMetadataFailKey.String(), Other: "Échec de récupération des métadonnées"},
	getMetadataConversionFailKey: {ID: getMetadataConversionFailKey.String(), Other: "Échec de la conversion des métadonnées {{.Source}} en {{.Target}}"},
	missingMetadataKey: {
		ID:    missingMetadataKey.String(),
		Other: "Métadonnées manquantes {{.Name}}, veuillez vérifier si la transmission de métadonnées est activée dans la chaîne de transmission et transmettre la valeur",
	},
	encryptFailKey:          {ID: encryptFailKey.String(), Other: "Échec de chiffrement"},
	decryptFailKey:          {ID: decryptFailKey.String(), Other: "Échec de déchiffrement"},
	paramCanNotEmptyKey:     {ID: paramCanNotEmptyKey.String(), Other: "Le paramètre {{.Name}} ne peut pas être vide"},
	notEditableKey:          {ID: notEditableKey.String(), Other: "Les données ne sont pas modifiables"},
	requestErrorKey:         {ID: requestErrorKey.String(), Other: "Erreur de requête"},
	pleaseLoginKey:          {ID: pleaseLoginKey.String(), Other: "Veuillez vous connecter d'abord"},
	needReLoginKey:          {ID: needReLoginKey.String(), Other: "Besoin de se reconnecter"},
	accountDisabledKey:      {ID: accountDisabledKey.String(), Other: "Le compte a été désactivé"},
	pleaseChangePasswordKey: {ID: pleaseChangePasswordKey.String(), Other: "Veuillez d'abord changer votre mot de passe"},
	passwordErrorKey:        {ID: passwordErrorKey.String(), Other: "Erreur de mot de passe"},
	noPermissionKey:         {ID: noPermissionKey.String(), Other: "Pas de permission"},
	notConfiguredKey:        {ID: notConfiguredKey.String(), Other: "{{.Name}} non configuré"},
	canNotEmptyKey:          {ID: canNotEmptyKey.String(), Other: "{{.Name}} ne peut pas être vide"},
	expiredKey:              {ID: expiredKey.String(), Other: "{{.Name}} a expiré"},

	dataAbnormalKey:         {ID: dataAbnormalKey.String(), Other: "{{.Name}} données anormales"},
	dataNotFoundKey:         {ID: dataNotFoundKey.String(), Other: "Données non trouvées"},
	dataDuplicateKey:        {ID: dataDuplicateKey.String(), Other: "Les données existent déjà, veuillez ne pas les ajouter à nouveau"},
	dataConstraintKey:       {ID: dataConstraintKey.String(), Other: "Échec de la vérification des contraintes de données, veuillez vérifier vos paramètres"},
	dataNotLoadedKey:        {ID: dataNotLoadedKey.String(), Other: "Base de données non chargée, veuillez contacter l'administrateur"},
	dataNotSingularKey:      {ID: dataNotSingularKey.String(), Other: "Erreur de données Not Singular, veuillez contacter l'administrateur"},
	dataValidationErrorKey:  {ID: dataValidationErrorKey.String(), Other: "Échec de la validation des données, veuillez vérifier vos paramètres"},
	dataErrorKey:            {ID: dataErrorKey.String(), Other: "Erreur de couche de données, veuillez contacter l'administrateur"},
	cacheNotFoundKey:        {ID: cacheNotFoundKey.String(), Other: "Cache introuvable"},
	cachePreMatchGetFailKey: {ID: cachePreMatchGetFailKey.String(), Other: "Échec de la récupération de la pré-correspondance du cache"},
	cacheSetFailKey:         {ID: cacheSetFailKey.String(), Other: "Échec de la définition du cache"},
	cacheMSetFailKey:        {ID: cacheMSetFailKey.String(), Other: "Échec de la définition du cache en masse"},
	cacheDelFailKey:         {ID: cacheDelFailKey.String(), Other: "Échec de la suppression du cache"},
	cachePreMatchDelFailKey: {ID: cachePreMatchDelFailKey.String(), Other: "Échec de la suppression de la pré-correspondance du cache"},
	cacheFlushFailKey:       {ID: cacheFlushFailKey.String(), Other: "Échec du vidage du cache"},
	cacheMGetFailKey:        {ID: cacheMGetFailKey.String(), Other: "Échec de la récupération en masse du cache"},
	cacheMDelFailKey:        {ID: cacheMDelFailKey.String(), Other: "Échec de la suppression en masse du cache"},

	// String validation messages
	validateStringConstKey:             {ID: validateStringConstKey.String(), Other: "Doit être exactement '{{.Rule}}'"},
	validateStringLenKey:               {ID: validateStringLenKey.String(), Other: "La longueur doit être exactement de {{.Rule}} caractères"},
	validateStringMinLenKey:            {ID: validateStringMinLenKey.String(), Other: "La longueur doit être d'au moins {{.Rule}} caractères"},
	validateStringMaxLenKey:            {ID: validateStringMaxLenKey.String(), Other: "La longueur doit être d'au plus {{.Rule}} caractères"},
	validateStringPatternKey:           {ID: validateStringPatternKey.String(), Other: "Doit correspondre au modèle '{{.Rule}}'"},
	validateStringPrefixKey:            {ID: validateStringPrefixKey.String(), Other: "Doit commencer par '{{.Rule}}'"},
	validateStringSuffixKey:            {ID: validateStringSuffixKey.String(), Other: "Doit se terminer par '{{.Rule}}'"},
	validateStringContainsKey:          {ID: validateStringContainsKey.String(), Other: "Doit contenir '{{.Rule}}'"},
	validateStringNotContainsKey:       {ID: validateStringNotContainsKey.String(), Other: "Ne doit pas contenir '{{.Rule}}'"},
	validateStringInKey:                {ID: validateStringInKey.String(), Other: "Doit être l'un de [{{.Rule}}]"},
	validateStringNotInKey:             {ID: validateStringNotInKey.String(), Other: "Ne doit pas être l'un de [{{.Rule}}]"},
	validateStringEmailKey:             {ID: validateStringEmailKey.String(), Other: "Doit être une adresse email valide"},
	validateStringHostnameKey:          {ID: validateStringHostnameKey.String(), Other: "Doit être un nom d'hôte valide"},
	validateStringIPKey:                {ID: validateStringIPKey.String(), Other: "Doit être une adresse IP valide"},
	validateStringIPv4Key:              {ID: validateStringIPv4Key.String(), Other: "Doit être une adresse IPv4 valide"},
	validateStringIPv6Key:              {ID: validateStringIPv6Key.String(), Other: "Doit être une adresse IPv6 valide"},
	validateStringURIKey:               {ID: validateStringURIKey.String(), Other: "Doit être un URI valide"},
	validateStringURIRefKey:            {ID: validateStringURIRefKey.String(), Other: "Doit être une référence URI valide"},
	validateStringAddressKey:           {ID: validateStringAddressKey.String(), Other: "Doit être une adresse valide"},
	validateStringUUIDKey:              {ID: validateStringUUIDKey.String(), Other: "Doit être un UUID valide"},
	validateStringWellKnownRegexKey:    {ID: validateStringWellKnownRegexKey.String(), Other: "Doit correspondre au format {{.Rule}}"},
	validateStringLenBytesKey:          {ID: validateStringLenBytesKey.String(), Other: "Doit être exactement {{.Rule}} octets"},
	validateStringMinBytesKey:          {ID: validateStringMinBytesKey.String(), Other: "Doit être d'au moins {{.Rule}} octets"},
	validateStringMaxBytesKey:          {ID: validateStringMaxBytesKey.String(), Other: "Doit être d'au plus {{.Rule}} octets"},
	validateStringTUUIDKey:             {ID: validateStringTUUIDKey.String(), Other: "Doit être un UUID basé sur le temps valide"},
	validateStringIPPrefixKey:          {ID: validateStringIPPrefixKey.String(), Other: "Doit être un préfixe de réseau IP valide"},
	validateStringIPWithPrefixlenKey:   {ID: validateStringIPWithPrefixlenKey.String(), Other: "Doit être une adresse IP valide avec une longueur de préfixe"},
	validateStringIPv4WithPrefixlenKey: {ID: validateStringIPv4WithPrefixlenKey.String(), Other: "Doit être une adresse IPv4 valide avec une longueur de préfixe"},
	validateStringIPv6WithPrefixlenKey: {ID: validateStringIPv6WithPrefixlenKey.String(), Other: "Doit être une adresse IPv6 valide avec une longueur de préfixe"},
	validateStringStrictKey:            {ID: validateStringStrictKey.String(), Other: "Doit respecter les règles de validation strictes"},
	validateStringExampleKey:           {ID: validateStringExampleKey.String(), Other: "Doit correspondre au format d'exemple: {{.Rule}}"},

	// Boolean validation messages
	validateBoolConstKey: {ID: validateBoolConstKey.String(), Other: "Doit être {{.Rule}}"},

	// Enum validation messages
	validateEnumConstKey:   {ID: validateEnumConstKey.String(), Other: "Doit être exactement {{.Rule}}"},
	validateEnumDefinedKey: {ID: validateEnumDefinedKey.String(), Other: "Doit être une valeur d'énumération définie"},
	validateEnumInKey:      {ID: validateEnumInKey.String(), Other: "Doit être l'un de {{.Rule}}"},
	validateEnumNotInKey:   {ID: validateEnumNotInKey.String(), Other: "Ne doit pas être l'un de {{.Rule}}"},

	// Repeated field validation messages
	validateRepeatedMinItemsKey: {ID: validateRepeatedMinItemsKey.String(), Other: "Doit avoir au moins {{.Rule}} éléments"},
	validateRepeatedMaxItemsKey: {ID: validateRepeatedMaxItemsKey.String(), Other: "Doit avoir au plus {{.Rule}} éléments"},
	validateRepeatedUniqueKey:   {ID: validateRepeatedUniqueKey.String(), Other: "Tous les éléments doivent être uniques"},
	validateRepeatedItemsKey:    {ID: validateRepeatedItemsKey.String(), Other: "Les éléments doivent satisfaire les règles"},

	// Map validation messages
	validateMapMinPairsKey: {ID: validateMapMinPairsKey.String(), Other: "Doit avoir au moins {{.Rule}} paires clé-valeur"},
	validateMapMaxPairsKey: {ID: validateMapMaxPairsKey.String(), Other: "Doit avoir au plus {{.Rule}} paires clé-valeur"},
	validateMapNoSparseKey: {ID: validateMapNoSparseKey.String(), Other: "Ne doit pas contenir de valeurs nulles"},
	validateMapKeysKey:     {ID: validateMapKeysKey.String(), Other: "Les clés doivent satisfaire les règles"},
	validateMapValuesKey:   {ID: validateMapValuesKey.String(), Other: "Les valeurs doivent satisfaire les règles"},

	// Float validation messages
	validateFloatConstKey:  {ID: validateFloatConstKey.String(), Other: "Doit être exactement {{.Rule}}"},
	validateFloatLtKey:     {ID: validateFloatLtKey.String(), Other: "Doit être inférieur à {{.Rule}}"},
	validateFloatLteKey:    {ID: validateFloatLteKey.String(), Other: "Doit être inférieur ou égal à {{.Rule}}"},
	validateFloatGtKey:     {ID: validateFloatGtKey.String(), Other: "Doit être supérieur à {{.Rule}}"},
	validateFloatGteKey:    {ID: validateFloatGteKey.String(), Other: "Doit être supérieur ou égal à {{.Rule}}"},
	validateFloatInKey:     {ID: validateFloatInKey.String(), Other: "Doit être l'un de {{.Rule}}"},
	validateFloatNotInKey:  {ID: validateFloatNotInKey.String(), Other: "Ne doit pas être l'un de {{.Rule}}"},
	validateFloatFiniteKey: {ID: validateFloatFiniteKey.String(), Other: "Doit être un nombre fini"},

	// Double validation messages
	validateDoubleConstKey:  {ID: validateDoubleConstKey.String(), Other: "Doit être exactement {{.Rule}}"},
	validateDoubleLtKey:     {ID: validateDoubleLtKey.String(), Other: "Doit être inférieur à {{.Rule}}"},
	validateDoubleLteKey:    {ID: validateDoubleLteKey.String(), Other: "Doit être inférieur ou égal à {{.Rule}}"},
	validateDoubleGtKey:     {ID: validateDoubleGtKey.String(), Other: "Doit être supérieur à {{.Rule}}"},
	validateDoubleGteKey:    {ID: validateDoubleGteKey.String(), Other: "Doit être supérieur ou égal à {{.Rule}}"},
	validateDoubleInKey:     {ID: validateDoubleInKey.String(), Other: "Doit être l'un de {{.Rule}}"},
	validateDoubleNotInKey:  {ID: validateDoubleNotInKey.String(), Other: "Ne doit pas être l'un de {{.Rule}}"},
	validateDoubleFiniteKey: {ID: validateDoubleFiniteKey.String(), Other: "Doit être un nombre fini"},

	// Int32 validation messages
	validateInt32ConstKey: {ID: validateInt32ConstKey.String(), Other: "Doit être exactement {{.Rule}}"},
	validateInt32LtKey:    {ID: validateInt32LtKey.String(), Other: "Doit être inférieur à {{.Rule}}"},
	validateInt32LteKey:   {ID: validateInt32LteKey.String(), Other: "Doit être inférieur ou égal à {{.Rule}}"},
	validateInt32GtKey:    {ID: validateInt32GtKey.String(), Other: "Doit être supérieur à {{.Rule}}"},
	validateInt32GteKey:   {ID: validateInt32GteKey.String(), Other: "Doit être supérieur ou égal à {{.Rule}}"},
	validateInt32InKey:    {ID: validateInt32InKey.String(), Other: "Doit être l'un de {{.Rule}}"},
	validateInt32NotInKey: {ID: validateInt32NotInKey.String(), Other: "Ne doit pas être l'un de {{.Rule}}"},

	// Int64 validation messages
	validateInt64ConstKey: {ID: validateInt64ConstKey.String(), Other: "Doit être exactement {{.Rule}}"},
	validateInt64LtKey:    {ID: validateInt64LtKey.String(), Other: "Doit être inférieur à {{.Rule}}"},
	validateInt64LteKey:   {ID: validateInt64LteKey.String(), Other: "Doit être inférieur ou égal à {{.Rule}}"},
	validateInt64GtKey:    {ID: validateInt64GtKey.String(), Other: "Doit être supérieur à {{.Rule}}"},
	validateInt64GteKey:   {ID: validateInt64GteKey.String(), Other: "Doit être supérieur ou égal à {{.Rule}}"},
	validateInt64InKey:    {ID: validateInt64InKey.String(), Other: "Doit être l'un de {{.Rule}}"},
	validateInt64NotInKey: {ID: validateInt64NotInKey.String(), Other: "Ne doit pas être l'un de {{.Rule}}"},

	// UInt32 validation messages
	validateUInt32ConstKey: {ID: validateUInt32ConstKey.String(), Other: "Doit être exactement {{.Rule}}"},
	validateUInt32LtKey:    {ID: validateUInt32LtKey.String(), Other: "Doit être inférieur à {{.Rule}}"},
	validateUInt32LteKey:   {ID: validateUInt32LteKey.String(), Other: "Doit être inférieur ou égal à {{.Rule}}"},
	validateUInt32GtKey:    {ID: validateUInt32GtKey.String(), Other: "Doit être supérieur à {{.Rule}}"},
	validateUInt32GteKey:   {ID: validateUInt32GteKey.String(), Other: "Doit être supérieur ou égal à {{.Rule}}"},
	validateUInt32InKey:    {ID: validateUInt32InKey.String(), Other: "Doit être l'un de {{.Rule}}"},
	validateUInt32NotInKey: {ID: validateUInt32NotInKey.String(), Other: "Ne doit pas être l'un de {{.Rule}}"},

	// UInt64 validation messages
	validateUInt64ConstKey: {ID: validateUInt64ConstKey.String(), Other: "Doit être exactement {{.Rule}}"},
	validateUInt64LtKey:    {ID: validateUInt64LtKey.String(), Other: "Doit être inférieur à {{.Rule}}"},
	validateUInt64LteKey:   {ID: validateUInt64LteKey.String(), Other: "Doit être inférieur ou égal à {{.Rule}}"},
	validateUInt64GtKey:    {ID: validateUInt64GtKey.String(), Other: "Doit être supérieur à {{.Rule}}"},
	validateUInt64GteKey:   {ID: validateUInt64GteKey.String(), Other: "Doit être supérieur ou égal à {{.Rule}}"},
	validateUInt64InKey:    {ID: validateUInt64InKey.String(), Other: "Doit être l'un de {{.Rule}}"},
	validateUInt64NotInKey: {ID: validateUInt64NotInKey.String(), Other: "Ne doit pas être l'un de {{.Rule}}"},

	// SInt32 validation messages
	validateSInt32ConstKey: {ID: validateSInt32ConstKey.String(), Other: "Doit être exactement {{.Rule}}"},
	validateSInt32LtKey:    {ID: validateSInt32LtKey.String(), Other: "Doit être inférieur à {{.Rule}}"},
	validateSInt32LteKey:   {ID: validateSInt32LteKey.String(), Other: "Doit être inférieur ou égal à {{.Rule}}"},
	validateSInt32GtKey:    {ID: validateSInt32GtKey.String(), Other: "Doit être supérieur à {{.Rule}}"},
	validateSInt32GteKey:   {ID: validateSInt32GteKey.String(), Other: "Doit être supérieur ou égal à {{.Rule}}"},
	validateSInt32InKey:    {ID: validateSInt32InKey.String(), Other: "Doit être l'un de {{.Rule}}"},
	validateSInt32NotInKey: {ID: validateSInt32NotInKey.String(), Other: "Ne doit pas être l'un de {{.Rule}}"},

	// SInt64 validation messages
	validateSInt64ConstKey: {ID: validateSInt64ConstKey.String(), Other: "Doit être exactement {{.Rule}}"},
	validateSInt64LtKey:    {ID: validateSInt64LtKey.String(), Other: "Doit être inférieur à {{.Rule}}"},
	validateSInt64LteKey:   {ID: validateSInt64LteKey.String(), Other: "Doit être inférieur ou égal à {{.Rule}}"},
	validateSInt64GtKey:    {ID: validateSInt64GtKey.String(), Other: "Doit être supérieur à {{.Rule}}"},
	validateSInt64GteKey:   {ID: validateSInt64GteKey.String(), Other: "Doit être supérieur ou égal à {{.Rule}}"},
	validateSInt64InKey:    {ID: validateSInt64InKey.String(), Other: "Doit être l'un de {{.Rule}}"},
	validateSInt64NotInKey: {ID: validateSInt64NotInKey.String(), Other: "Ne doit pas être l'un de {{.Rule}}"},

	// Fixed32 validation messages
	validateFixed32ConstKey: {ID: validateFixed32ConstKey.String(), Other: "Doit être exactement {{.Rule}}"},
	validateFixed32LtKey:    {ID: validateFixed32LtKey.String(), Other: "Doit être inférieur à {{.Rule}}"},
	validateFixed32LteKey:   {ID: validateFixed32LteKey.String(), Other: "Doit être inférieur ou égal à {{.Rule}}"},
	validateFixed32GtKey:    {ID: validateFixed32GtKey.String(), Other: "Doit être supérieur à {{.Rule}}"},
	validateFixed32GteKey:   {ID: validateFixed32GteKey.String(), Other: "Doit être supérieur ou égal à {{.Rule}}"},
	validateFixed32InKey:    {ID: validateFixed32InKey.String(), Other: "Doit être l'un de {{.Rule}}"},
	validateFixed32NotInKey: {ID: validateFixed32NotInKey.String(), Other: "Ne doit pas être l'un de {{.Rule}}"},

	// Fixed64 validation messages
	validateFixed64ConstKey: {ID: validateFixed64ConstKey.String(), Other: "Doit être exactement {{.Rule}}"},
	validateFixed64LtKey:    {ID: validateFixed64LtKey.String(), Other: "Doit être inférieur à {{.Rule}}"},
	validateFixed64LteKey:   {ID: validateFixed64LteKey.String(), Other: "Doit être inférieur ou égal à {{.Rule}}"},
	validateFixed64GtKey:    {ID: validateFixed64GtKey.String(), Other: "Doit être supérieur à {{.Rule}}"},
	validateFixed64GteKey:   {ID: validateFixed64GteKey.String(), Other: "Doit être supérieur ou égal à {{.Rule}}"},
	validateFixed64InKey:    {ID: validateFixed64InKey.String(), Other: "Doit être l'un de {{.Rule}}"},
	validateFixed64NotInKey: {ID: validateFixed64NotInKey.String(), Other: "Ne doit pas être l'un de {{.Rule}}"},

	// SFixed32 validation messages
	validateSFixed32ConstKey: {ID: validateSFixed32ConstKey.String(), Other: "Doit être exactement {{.Rule}}"},
	validateSFixed32LtKey:    {ID: validateSFixed32LtKey.String(), Other: "Doit être inférieur à {{.Rule}}"},
	validateSFixed32LteKey:   {ID: validateSFixed32LteKey.String(), Other: "Doit être inférieur ou égal à {{.Rule}}"},
	validateSFixed32GtKey:    {ID: validateSFixed32GtKey.String(), Other: "Doit être supérieur à {{.Rule}}"},
	validateSFixed32GteKey:   {ID: validateSFixed32GteKey.String(), Other: "Doit être supérieur ou égal à {{.Rule}}"},
	validateSFixed32InKey:    {ID: validateSFixed32InKey.String(), Other: "Doit être l'un de {{.Rule}}"},
	validateSFixed32NotInKey: {ID: validateSFixed32NotInKey.String(), Other: "Ne doit pas être l'un de {{.Rule}}"},

	// SFixed64 validation messages
	validateSFixed64ConstKey: {ID: validateSFixed64ConstKey.String(), Other: "Doit être exactement {{.Rule}}"},
	validateSFixed64LtKey:    {ID: validateSFixed64LtKey.String(), Other: "Doit être inférieur à {{.Rule}}"},
	validateSFixed64LteKey:   {ID: validateSFixed64LteKey.String(), Other: "Doit être inférieur ou égal à {{.Rule}}"},
	validateSFixed64GtKey:    {ID: validateSFixed64GtKey.String(), Other: "Doit être supérieur à {{.Rule}}"},
	validateSFixed64GteKey:   {ID: validateSFixed64GteKey.String(), Other: "Doit être supérieur ou égal à {{.Rule}}"},
	validateSFixed64InKey:    {ID: validateSFixed64InKey.String(), Other: "Doit être l'un de {{.Rule}}"},
	validateSFixed64NotInKey: {ID: validateSFixed64NotInKey.String(), Other: "Ne doit pas être l'un de {{.Rule}}"},

	// Bytes validation messages
	validateBytesConstKey:    {ID: validateBytesConstKey.String(), Other: "Must be exactly '{{.Rule}}'"},
	validateBytesLenKey:      {ID: validateBytesLenKey.String(), Other: "Length must be exactly {{.Rule}} bytes"},
	validateBytesMinLenKey:   {ID: validateBytesMinLenKey.String(), Other: "Length must be at least {{.Rule}} bytes"},
	validateBytesMaxLenKey:   {ID: validateBytesMaxLenKey.String(), Other: "Length must be at most {{.Rule}} bytes"},
	validateBytesPatternKey:  {ID: validateBytesPatternKey.String(), Other: "Must match the pattern '{{.Rule}}'"},
	validateBytesPrefixKey:   {ID: validateBytesPrefixKey.String(), Other: "Must start with '{{.Rule}}'"},
	validateBytesSuffixKey:   {ID: validateBytesSuffixKey.String(), Other: "Must end with '{{.Rule}}'"},
	validateBytesContainsKey: {ID: validateBytesContainsKey.String(), Other: "Must contain '{{.Rule}}'"},
	validateBytesInKey:       {ID: validateBytesInKey.String(), Other: "Must be one of [{{.Rule}}]"},
	validateBytesNotInKey:    {ID: validateBytesNotInKey.String(), Other: "Must not be one of [{{.Rule}}]"},

	// Message validation messages
	validateMessageRequiredKey: {ID: validateMessageRequiredKey.String(), Other: "Message is required"},
	validateMessageSkipKey:     {ID: validateMessageSkipKey.String(), Other: "Message is skipped"},

	// Any validation messages
	validateAnyRequiredKey: {ID: validateAnyRequiredKey.String(), Other: "Any is required"},
	validateAnyInKey:       {ID: validateAnyInKey.String(), Other: "Any must be one of [{{.Rule}}]"},
	validateAnyNotInKey:    {ID: validateAnyNotInKey.String(), Other: "Any must not be one of [{{.Rule}}]"},

	// Duration validation messages
	validateDurationRequiredKey: {ID: validateDurationRequiredKey.String(), Other: "Duration is required"},
	validateDurationConstKey:    {ID: validateDurationConstKey.String(), Other: "Duration must be '{{.Rule}}'"},
	validateDurationLtKey:       {ID: validateDurationLtKey.String(), Other: "Duration must be less than {{.Rule}}"},
	validateDurationLteKey:      {ID: validateDurationLteKey.String(), Other: "Duration must be less than or equal to {{.Rule}}"},
	validateDurationGtKey:       {ID: validateDurationGtKey.String(), Other: "Duration must be greater than {{.Rule}}"},
	validateDurationGteKey:      {ID: validateDurationGteKey.String(), Other: "Duration must be greater than or equal to {{.Rule}}"},
	validateDurationInKey:       {ID: validateDurationInKey.String(), Other: "Duration must be one of [{{.Rule}}]"},
	validateDurationNotInKey:    {ID: validateDurationNotInKey.String(), Other: "Duration must not be one of [{{.Rule}}]"},

	// Timestamp validation messages
	validateTimestampRequiredKey: {ID: validateTimestampRequiredKey.String(), Other: "Timestamp is required"},
	validateTimestampConstKey:    {ID: validateTimestampConstKey.String(), Other: "Timestamp must be '{{.Rule}}'"},
	validateTimestampLtKey:       {ID: validateTimestampLtKey.String(), Other: "Timestamp must be less than {{.Rule}}"},
	validateTimestampLteKey:      {ID: validateTimestampLteKey.String(), Other: "Timestamp must be less than or equal to {{.Rule}}"},
	validateTimestampGtKey:       {ID: validateTimestampGtKey.String(), Other: "Timestamp must be greater than {{.Rule}}"},
	validateTimestampGteKey:      {ID: validateTimestampGteKey.String(), Other: "Timestamp must be greater than or equal to {{.Rule}}"},
	validateTimestampLtNowKey:    {ID: validateTimestampLtNowKey.String(), Other: "Timestamp must be less than the current time"},
	validateTimestampGtNowKey:    {ID: validateTimestampGtNowKey.String(), Other: "Timestamp must be greater than the current time"},
	validateTimestampWithinKey:   {ID: validateTimestampWithinKey.String(), Other: "Timestamp must be within the time range"},
}

var langMap = map[language.Tag]map[MsgKey]*i18n.Message{
	language.English: dfMsgMap,
	language.Chinese: zhMsgMap,
	language.Russian: ruMsgMap,
	language.French:  frMsgMap,
}
