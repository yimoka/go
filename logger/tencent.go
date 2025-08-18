// Package logger tencent.go
package logger

import (
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	cls "github.com/tencentcloud/tencentcloud-cls-sdk-go"
	"github.com/yimoka/go/config"
	"google.golang.org/protobuf/proto"
)

// GetTencentLogger _
func GetTencentLogger(conf *config.Config) log.Logger {
	tencentLogger, err := NewTencentLogger(conf)
	if err != nil {
		panic(fmt.Sprintf("init tencent logger error: %v", err))
	}

	return getLogger(conf.Server, tencentLogger, conf.Logger, false)
}

// TencentLogger _
type TencentLogger interface {
	log.Logger

	GetProducer() *cls.AsyncProducerClient
	Close() error
}

type tencentLog struct {
	producer  *cls.AsyncProducerClient
	opts      *config.Logger
	stdLogger log.Logger
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

	// 发送到腾讯云日志
	err := log.producer.SendLog(log.opts.TopicID, logInst, nil)

	// 如果配置了同时输出到标准输出
	if log.stdLogger != nil {
		_ = log.stdLogger.Log(level, keyValues...)
	}

	return err
}

// NewTencentLogger _
func NewTencentLogger(conf *config.Config) (TencentLogger, error) {
	opts := conf.Logger
	producerConfig := cls.GetDefaultAsyncProducerClientConfig()
	producerConfig.AccessKeyID = opts.AccessKey
	producerConfig.AccessKeySecret = opts.AccessSecret
	producerConfig.Endpoint = opts.Endpoint
	producerInst, err := cls.NewAsyncProducerClient(producerConfig)
	if err != nil {
		return nil, err
	}
	producerInst.Start()

	tencentLog := &tencentLog{
		producer: producerInst,
		opts:     opts,
	}

	// 如果配置了同时输出到标准输出
	if opts.AlsoStd {
		stdLogger := GetStdLogger(conf)
		tencentLog.stdLogger = stdLogger
	}

	return tencentLog, nil
}

func newString(s string) *string {
	return &s
}
