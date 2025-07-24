package metrics

import (
	"sync"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/yimoka/go/config"
	ymetrics "github.com/yimoka/go/metrics"
	"go.opentelemetry.io/otel/metric"
)

// 单例变量
var (
	once           sync.Once
	metricRequests metric.Int64Counter
	metricSeconds  metric.Float64Histogram
)

func CreateServerMiddleware(conf *config.Metrics) middleware.Middleware {
	requests, seconds := GetDefaultMetrics(conf)
	return metrics.Server(
		metrics.WithSeconds(seconds),
		metrics.WithRequests(requests),
	)
}

func CreateClientMiddleware(conf *config.Metrics) middleware.Middleware {
	requests, seconds := GetDefaultMetrics(conf)
	return metrics.Client(
		metrics.WithSeconds(seconds),
		metrics.WithRequests(requests),
	)
}

// GetDefaultMetrics 使用单例模式获取默认的请求计数器和耗时直方图
func GetDefaultMetrics(conf *config.Metrics) (metric.Int64Counter, metric.Float64Histogram) {
	once.Do(func() {
		meter := ymetrics.GetMeter(conf)
		var err error

		metricRequests, err = metrics.DefaultRequestsCounter(meter, metrics.DefaultServerRequestsCounterName)
		if err != nil {
			panic(err)
		}

		metricSeconds, err = metrics.DefaultSecondsHistogram(meter, metrics.DefaultServerSecondsHistogramName)
		if err != nil {
			panic(err)
		}
	})

	return metricRequests, metricSeconds
}
