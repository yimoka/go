// cSpell: disable
package lang

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// MsgKey message key
type MsgKey string

func (m MsgKey) String() string {
	return string(m)
}

// 公共语言的消息配置
const (
	// 公共
	parameterErrorKey            MsgKey = "parameter_error"
	getMetadataFailKey           MsgKey = "get_metadata_fail"
	getMetadataConversionFailKey MsgKey = "get_metadata_int_fail"
	missingMetadataKey           MsgKey = "missing_metadata"
	encryptFailKey               MsgKey = "encrypt_fail"
	decryptFailKey               MsgKey = "decrypt_fail"
	paramCanNotEmptyKey          MsgKey = "param_can_not_empty"
	notEditableKey               MsgKey = "not_editable"
	requestErrorKey              MsgKey = "request_error"
	pleaseLoginKey               MsgKey = "please_login"
	needReLoginKey               MsgKey = "need_re_login"
	accountDisabledKey           MsgKey = "account_disabled"
	pleaseChangePasswordKey      MsgKey = "please_change_password"
	passwordErrorKey             MsgKey = "password_error"
	noPermissionKey              MsgKey = "no_permission"
	notConfiguredKey             MsgKey = "not_configured"
	canNotEmptyKey               MsgKey = "can_not_empty"
	expiredKey                   MsgKey = "expired"

	// 数据库
	dataAbnormalKey        MsgKey = "data_abnormal"
	dataNotFoundKey        MsgKey = "data_not_found"
	dataDuplicateKey       MsgKey = "data_duplicate_key"
	dataConstraintKey      MsgKey = "data_constraint_error"
	dataNotLoadedKey       MsgKey = "data_not_loaded"
	dataNotSingularKey     MsgKey = "data_not_singular"
	dataValidationErrorKey MsgKey = "data_validation_error"
	dataErrorKey           MsgKey = "data_error"
	// 缓存
	cacheNotFoundKey        MsgKey = "cache_not_found"
	cachePreMatchGetFailKey MsgKey = "cache_pre_match_get_fail"
	cacheSetFailKey         MsgKey = "cache_set_fail"
	cacheMSetFailKey        MsgKey = "cache_m_set_fail"
	cacheDelFailKey         MsgKey = "cache_del_fail"
	cachePreMatchDelFailKey MsgKey = "cache_pre_match_del_fail"
	cacheFlushFailKey       MsgKey = "cache_flush_fail"
	cacheMGetFailKey        MsgKey = "cache_m_get_fail"
	cacheMDelFailKey        MsgKey = "cache_m_del_fail"
)

// 默认语言的 msg 配置 英语
var dfMsgMap = map[MsgKey]*i18n.Message{
	parameterErrorKey:            {ID: parameterErrorKey.String(), Other: "Parameter error, please check your parameters"},
	getMetadataFailKey:           {ID: getMetadataFailKey.String(), Other: "Get metadata failed"},
	getMetadataConversionFailKey: {ID: getMetadataConversionFailKey.String(), Other: "metadata {{.Source}} conversion to {{.Target}} failed"},
	missingMetadataKey: {
		ID:    missingMetadataKey.String(),
		Other: "Missing metadata {{.Name}}, please check whether the transmission link is enabled for metadata transmission and pass the value"},
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

	dataErrorKey:            {ID: dataErrorKey.String(), Other: "数据层出错了,请联系管理员"},
	cacheNotFoundKey:        {ID: cacheNotFoundKey.String(), Other: "缓存不存在"},
	cachePreMatchGetFailKey: {ID: cachePreMatchGetFailKey.String(), Other: "前置匹配获取缓存失败"},
	cacheSetFailKey:         {ID: cacheSetFailKey.String(), Other: "设置缓存失败"},
	cacheMSetFailKey:        {ID: cacheMSetFailKey.String(), Other: "批量设置缓存失败"},
	cacheDelFailKey:         {ID: cacheDelFailKey.String(), Other: "删除缓存失败"},
	cachePreMatchDelFailKey: {ID: cachePreMatchDelFailKey.String(), Other: "前置匹配删除缓存失败"},
	cacheFlushFailKey:       {ID: cacheFlushFailKey.String(), Other: "清空缓存失败"},
	cacheMGetFailKey:        {ID: cacheMGetFailKey.String(), Other: "批量获取缓存失败"},
	cacheMDelFailKey:        {ID: cacheMDelFailKey.String(), Other: "批量删除缓存失败"},
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
}

// 法语
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
}

var langMap = map[language.Tag]map[MsgKey]*i18n.Message{
	language.English: dfMsgMap,
	language.Chinese: zhMsgMap,
	language.Russian: ruMsgMap,
	language.French:  frMsgMap,
}
