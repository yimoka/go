// Package logger  logger.go
package logger

import (
	"os"
	"regexp"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/yimoka/go/config"
)

// GetLogger _
func GetLogger(conf *config.Config) log.Logger {
	logger := conf.Logger
	if logger == nil {
		return GetStdLogger(conf)
	}
	if logger.Provider == "tencent" {
		return GetTencentLogger(conf)
	}
	if logger.Provider == "otel" {
		return GetOtelLogger(conf)
	}
	return GetStdLogger(conf)
}

// GetStdLogger _
func GetStdLogger(conf *config.Config) log.Logger {
	logger := log.NewStdLogger(os.Stdout)
	return getLogger(conf.Server, logger, conf.Logger)
}

// GetLogger _
func getLogger(service *config.Server, logger log.Logger, loggerConf *config.Logger) log.Logger {
	// 应用日志过滤
	if loggerConf != nil {
		var filterOptions []log.FilterOption

		// 按日志级别过滤
		if loggerConf.FilterLevel != "" {
			level := parseLogLevel(loggerConf.FilterLevel)
			filterOptions = append(filterOptions, log.FilterLevel(level))
		}

		// 按key过滤（脱敏）
		if len(loggerConf.FilterKeys) > 0 {
			filterOptions = append(filterOptions, log.FilterKey(loggerConf.FilterKeys...))
		}

		// 按value过滤（脱敏）
		if len(loggerConf.FilterValues) > 0 {
			filterOptions = append(filterOptions, log.FilterValue(loggerConf.FilterValues...))
		}

		// 按敏感信息正则表达式脱敏
		if len(loggerConf.SensitiveRegex) > 0 {
			filterOptions = append(filterOptions, log.FilterFunc(func(_ log.Level, keyvals ...any) bool {
				for _, regex := range loggerConf.SensitiveRegex {
					re, err := regexp.Compile(regex)
					if err != nil {
						continue // 跳过无效的正则表达式
					}
					for i := 1; i < len(keyvals); i += 2 {
						valueStr := toString(keyvals[i])
						if strings.Contains(regex, ":") && re.MatchString(valueStr) {
							// 使用正则表达式替换，保持字段名不变
							keyvals[i] = re.ReplaceAllStringFunc(valueStr, func(match string) string {
								// 提取字段名（冒号前的部分）
								if colonIndex := strings.Index(match, ":"); colonIndex > 0 {
									fieldName := match[:colonIndex]
									return fieldName + ":\"***\""
								}
								return "***"
							})
						}
					}
				}
				return false
			}))
		}

		// 应用所有过滤选项
		if len(filterOptions) > 0 {
			logger = log.NewFilter(logger, filterOptions...)
		}
	}

	return log.With(logger,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", service.Id,
		"service.name", service.Name,
		"service.version", service.Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}

// parseLogLevel 解析日志级别字符串为log.Level
func parseLogLevel(level string) log.Level {
	switch level {
	case "debug":
		return log.LevelDebug
	case "info":
		return log.LevelInfo
	case "warn":
		return log.LevelWarn
	case "error":
		return log.LevelError
	case "fatal":
		return log.LevelFatal
	default:
		return log.LevelInfo
	}
}
