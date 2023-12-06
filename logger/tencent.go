// Package logger tencent.go
package logger

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	cls "github.com/tencentcloud/tencentcloud-cls-sdk-go"
	"github.com/yimoka/go/config"
	"google.golang.org/protobuf/proto"
)

// GetTencentLogger _
func GetTencentLogger(conf *config.Config) log.Logger {
	logger, err := NewTencentLogger(conf.Logger)
	if err != nil {
		panic(fmt.Sprintf("init tencent logger error: %v", err))
	}
	return getLogger(conf.Server, logger)
}

// TencentLogger _
type TencentLogger interface {
	log.Logger

	GetProducer() *cls.AsyncProducerClient
	Close() error
}

type tencentLog struct {
	producer *cls.AsyncProducerClient
	opts     *config.Logger
}

func (log *tencentLog) GetProducer() *cls.AsyncProducerClient {
	return log.producer
}

func (log *tencentLog) Close() error {
	return log.producer.Close(5000)
}

func (log *tencentLog) Log(level log.Level, keyValues ...interface{}) error {
	contents := make([]*cls.Log_Content, 0, len(keyValues)/2+1)

	contents = append(contents, &cls.Log_Content{
		Key:   newString(level.Key()),
		Value: newString(level.String()),
	})
	for i := 0; i < len(keyValues); i += 2 {
		contents = append(contents, &cls.Log_Content{
			Key:   newString(toString(keyValues[i])),
			Value: newString(toString(keyValues[i+1])),
		})
	}

	logInst := &cls.Log{
		Time:     proto.Int64(time.Now().Unix()),
		Contents: contents,
	}
	return log.producer.SendLog(log.opts.TopicID, logInst, nil)
}

// NewTencentLogger _
func NewTencentLogger(opts *config.Logger) (TencentLogger, error) {
	producerConfig := cls.GetDefaultAsyncProducerClientConfig()
	producerConfig.AccessKeyID = opts.AccessKey
	producerConfig.AccessKeySecret = opts.AccessSecret
	producerConfig.Endpoint = opts.Endpoint
	producerInst, err := cls.NewAsyncProducerClient(producerConfig)
	if err != nil {
		return nil, err
	}
	producerInst.Start()
	return &tencentLog{
		producer: producerInst,
		opts:     opts,
	}, nil
}

func newString(s string) *string {
	return &s
}

// toString convert any type to string
func toString(v interface{}) string {
	var key string
	if v == nil {
		return key
	}
	switch v := v.(type) {
	case float64:
		key = strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		key = strconv.FormatFloat(float64(v), 'f', -1, 32)
	case int:
		key = strconv.Itoa(v)
	case uint:
		key = strconv.FormatUint(uint64(v), 10)
	case int8:
		key = strconv.Itoa(int(v))
	case uint8:
		key = strconv.FormatUint(uint64(v), 10)
	case int16:
		key = strconv.Itoa(int(v))
	case uint16:
		key = strconv.FormatUint(uint64(v), 10)
	case int32:
		key = strconv.Itoa(int(v))
	case uint32:
		key = strconv.FormatUint(uint64(v), 10)
	case int64:
		key = strconv.FormatInt(v, 10)
	case uint64:
		key = strconv.FormatUint(v, 10)
	case string:
		key = v
	case bool:
		key = strconv.FormatBool(v)
	case []byte:
		key = string(v)
	case fmt.Stringer:
		key = v.String()
	default:
		newValue, _ := json.Marshal(v)
		key = string(newValue)
	}
	return key
}
