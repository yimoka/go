package lang

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
)

func (c *CommonLang) GetValidateErrorMsg(ctx context.Context, violation *protovalidate.Violation, langs ...string) (string, bool) {
	constraintID := violation.Proto.GetConstraintId()
	if constraintID == "" {
		return "", false
	}
	msgKey := validatePrefix + MsgKey(constraintID)
	// 判断是否存在
	if _, ok := dfMsgMap[msgKey]; !ok {
		return "", false
	}
	templateData := map[string]interface{}{
		"Field":     violation.FieldValue.String(),
		"FieldDesc": violation.FieldDescriptor.Name(),
		"Rule":      violation.RuleValue.String(),
		"RuleDesc":  violation.RuleDescriptor.Name(),
		"Message":   violation.Proto.GetMessage(),
	}
	return c.getMsg(ctx, msgKey, templateData, langs...), true
}
