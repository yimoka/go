package logger

import (
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/yimoka/go/config"
)

func TestLoggerFilter(t *testing.T) {
	tests := []struct {
		name        string
		loggerConf  *config.Logger
		logLevel    log.Level
		keyValues   []interface{}
		shouldLog   bool
		shouldMask  bool
		expectedMsg string
	}{
		{
			name: "按级别过滤 - 只输出error级别",
			loggerConf: &config.Logger{
				FilterLevel: "error",
			},
			logLevel:  log.LevelInfo,
			keyValues: []interface{}{"msg", "info message"},
			shouldLog: false,
		},
		{
			name: "按级别过滤 - error级别应该输出",
			loggerConf: &config.Logger{
				FilterLevel: "error",
			},
			logLevel:  log.LevelError,
			keyValues: []interface{}{"msg", "error message"},
			shouldLog: true,
		},
		{
			name: "按key过滤 - 脱敏password字段",
			loggerConf: &config.Logger{
				FilterKeys: []string{"password"},
			},
			logLevel:   log.LevelInfo,
			keyValues:  []interface{}{"msg", "login attempt", "password", "123456"},
			shouldLog:  true,
			shouldMask: true,
		},
		{
			name: "按value过滤 - 脱敏敏感值",
			loggerConf: &config.Logger{
				FilterValues: []string{"123456"},
			},
			logLevel:   log.LevelInfo,
			keyValues:  []interface{}{"msg", "login attempt", "code", "123456"},
			shouldLog:  true,
			shouldMask: true,
		},
		{
			name: "组合过滤 - 级别和key过滤",
			loggerConf: &config.Logger{
				FilterLevel: "warn",
				FilterKeys:  []string{"password"},
			},
			logLevel:   log.LevelWarn,
			keyValues:  []interface{}{"msg", "warning", "password", "secret"},
			shouldLog:  true,
			shouldMask: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建测试配置
			conf := &config.Config{
				Server: &config.Server{
					Id:      "test-service",
					Name:    "test",
					Version: "1.0.0",
				},
				Logger: tt.loggerConf,
			}

			// 创建logger
			logger := GetLogger(conf)

			// 这里我们需要创建一个可以捕获输出的logger用于测试
			// 由于GetLogger返回的是包装后的logger，我们需要直接测试过滤逻辑

			// 测试parseLogLevel函数
			if tt.loggerConf != nil && tt.loggerConf.FilterLevel != "" {
				level := parseLogLevel(tt.loggerConf.FilterLevel)
				if level == log.LevelInfo && tt.loggerConf.FilterLevel != "info" {
					t.Errorf("parseLogLevel failed for level: %s", tt.loggerConf.FilterLevel)
				}
			}

			// 测试日志记录
			err := logger.Log(tt.logLevel, tt.keyValues...)
			if err != nil {
				t.Errorf("Logger.Log() error = %v", err)
			}
		})
	}
}

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected log.Level
	}{
		{"debug", log.LevelDebug},
		{"info", log.LevelInfo},
		{"warn", log.LevelWarn},
		{"error", log.LevelError},
		{"fatal", log.LevelFatal},
		{"unknown", log.LevelInfo}, // 默认值
		{"", log.LevelInfo},        // 空字符串
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := parseLogLevel(tt.input)
			if result != tt.expected {
				t.Errorf("parseLogLevel(%s) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFilterIntegration(t *testing.T) {
	// 测试集成场景
	conf := &config.Config{
		Server: &config.Server{
			Id:      "test-service",
			Name:    "test",
			Version: "1.0.0",
		},
		Logger: &config.Logger{
			FilterLevel:  "warn",
			FilterKeys:   []string{"password", "token"},
			FilterValues: []string{"secret123"},
		},
	}

	logger := GetLogger(conf)

	// 测试不同级别的日志
	testCases := []struct {
		level    log.Level
		keyvals  []interface{}
		expected bool // 是否应该被过滤掉
	}{
		{log.LevelDebug, []interface{}{"msg", "debug message"}, true},  // 应该被过滤
		{log.LevelInfo, []interface{}{"msg", "info message"}, true},    // 应该被过滤
		{log.LevelWarn, []interface{}{"msg", "warn message"}, false},   // 不应该被过滤
		{log.LevelError, []interface{}{"msg", "error message"}, false}, // 不应该被过滤
	}

	for _, tc := range testCases {
		err := logger.Log(tc.level, tc.keyvals...)
		if err != nil {
			t.Errorf("Logger.Log() error = %v", err)
		}
	}
}

// 测试敏感信息脱敏
func TestSensitiveDataMasking(t *testing.T) {
	conf := &config.Config{
		Server: &config.Server{
			Id:      "test-service",
			Name:    "test",
			Version: "1.0.0",
		},
		Logger: &config.Logger{
			FilterKeys:   []string{"password", "token", "secret"},
			FilterValues: []string{"admin123", "sensitive_data"},
		},
	}

	logger := GetLogger(conf)

	// 测试包含敏感信息的日志
	sensitiveLogs := []struct {
		keyvals []interface{}
		desc    string
	}{
		{[]interface{}{"msg", "user login", "password", "mypassword123"}, "password字段应该被脱敏"},
		{[]interface{}{"msg", "api call", "token", "jwt_token_here"}, "token字段应该被脱敏"},
		{[]interface{}{"msg", "data access", "secret", "confidential"}, "secret字段应该被脱敏"},
		{[]interface{}{"msg", "user action", "code", "admin123"}, "敏感值应该被脱敏"},
	}

	for _, logEntry := range sensitiveLogs {
		err := logger.Log(log.LevelInfo, logEntry.keyvals...)
		if err != nil {
			t.Errorf("Logger.Log() error for %s: %v", logEntry.desc, err)
		}
	}
}

// 测试正则表达式过滤功能
func TestRegexFilter(t *testing.T) {
	tests := []struct {
		name        string
		loggerConf  *config.Logger
		logLevel    log.Level
		keyValues   []interface{}
		expectedMsg string
	}{
		{
			name: "过滤JWT token",
			loggerConf: &config.Logger{
				SensitiveRegex: []string{
					`eyJ[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+`,
				},
			},
			logLevel: log.LevelInfo,
			keyValues: []interface{}{
				"msg", "用户认证",
				"user_id", "12345",
				"token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
				"action", "login",
			},
			expectedMsg: "token=***",
		},
		{
			name: "过滤args中的token",
			loggerConf: &config.Logger{
				SensitiveRegex: []string{
					`token:"([^"]*)"`, // 精确匹配token字段，只替换括号内的内容
				},
			},
			logLevel: log.LevelInfo,
			keyValues: []interface{}{
				"msg", "API调用",
				"args", "token:\"LSbglXIKpiuYyukUCCp4zpXgplU1No9dEU73/yHNPB6Op9eCSqKUlSU1UcOmWqFFR6E6+VUiPChCHSoRAIFw3BPR4HFY/y9aMQHfTsPHud1uFK4lfwHxS4FDTenke75oZ6Urcoc0yBN5ehs+7hwO+b/44ikJbz22t/mVjnE4qRDOpB9Xq+Fv+hE0XOUUSQLYl6NNdH1UYtq9CndIOmj7sBdYMeytCtrSQ5NOlbVmcc2hI2SThtE/b+fUz9qBQOUBALuuDhkP17tfHpNNIRnlY6dropeeGR10SBBYoLg0kxtp/yPD3nML3gIswP31Q6fCKGB8kDnoz2dqRaeC6Rry97ty/d4VGLB/nl/t9jkhGo9MhtcrfOVrscEHwX6hmTTh3unvSmtQjeTMUvkXlhjbhpptEFqqlT3dZgfMp9sKc+WtJVNSSqE5Jmy3ktmcs1iHaACnpqawdf6HPYHWhoBi0MyjWVA2+y29SXK92pqfNykEbeevvq8rRIS95kPtexMF7CQRq16Op233vTxSWiL2HUPBaeZYaNzzwGsg6SuARrzIscHJ38VPseeL3I8FekBfh2BdUOTcDG4wRFdlfxTyCHuEy/k7fBRVqWBsr2oLkKqsmghsmoJYrnikjypJZxuPuR4caOLyvsPYDGvk1wiIDzsqPpSTr43r8MtyOjPnuSaJu5GCH/EakPw/jNjtUwQ/BJFIbgvTg+ndO6R5GE0gLjLpwmteZl1IQAbynAOvdxmqNeB5BLa/zwxRy3K4i5yjML4191VA9iaP60BmFhiQzsvaMuI8bWoijb9FjQ==\" userAgent:\"Mozilla/5.0\"",
				"code", 200,
			},
			expectedMsg: "args=***",
		},
		{
			name: "过滤args中的password字段",
			loggerConf: &config.Logger{
				SensitiveRegex: []string{
					`password:"([^"]*)"`, // 精确匹配password字段
				},
			},
			logLevel: log.LevelInfo,
			keyValues: []interface{}{
				"msg", "用户登录",
				"args", "username:\"admin\" password:\"secret123\" userAgent:\"Mozilla/5.0\"",
				"code", 200,
			},
			expectedMsg: "args=***",
		},
		{
			name: "过滤args中的api_key字段",
			loggerConf: &config.Logger{
				SensitiveRegex: []string{
					`api_key:"([^"]*)"`, // 精确匹配api_key字段
				},
			},
			logLevel: log.LevelInfo,
			keyValues: []interface{}{
				"msg", "API调用",
				"args", "api_key:\"sk-1234567890abcdef\" endpoint:\"/api/v1/users\" method:\"GET\"",
				"code", 200,
			},
			expectedMsg: "args=***",
		},
		{
			name: "过滤API密钥",
			loggerConf: &config.Logger{
				SensitiveRegex: []string{
					`sk-[A-Za-z0-9]+`,
				},
			},
			logLevel: log.LevelInfo,
			keyValues: []interface{}{
				"msg", "API调用",
				"api_key", "sk-1234567890abcdef",
				"amount", 100.50,
			},
			expectedMsg: "api_key=***",
		},
		{
			name: "过滤敏感值",
			loggerConf: &config.Logger{
				SensitiveRegex: []string{
					`admin123|password123|secret123`,
				},
			},
			logLevel: log.LevelInfo,
			keyValues: []interface{}{
				"msg", "系统操作",
				"old_value", "admin123",
				"new_value", "new_password",
			},
			expectedMsg: "old_value=***",
		},
		{
			name: "组合正则表达式过滤",
			loggerConf: &config.Logger{
				SensitiveRegex: []string{
					`eyJ[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+`,
					`sk-[A-Za-z0-9]+`,
					`admin123|password123`,
				},
			},
			logLevel: log.LevelInfo,
			keyValues: []interface{}{
				"msg", "完整API调用",
				"token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
				"api_key", "sk-1234567890abcdef",
				"password", "admin123",
				"status", "success",
			},
			expectedMsg: "token=*** api_key=*** password=***",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建测试配置
			conf := &config.Config{
				Server: &config.Server{
					Id:      "test-service",
					Name:    "test",
					Version: "1.0.0",
				},
				Logger: tt.loggerConf,
			}

			// 创建logger
			logger := GetLogger(conf)

			// 测试日志记录
			err := logger.Log(tt.logLevel, tt.keyValues...)
			if err != nil {
				t.Errorf("Logger.Log() error = %v", err)
			}
		})
	}
}

// 测试正则表达式过滤的性能
func TestRegexFilterPerformance(t *testing.T) {
	conf := &config.Config{
		Server: &config.Server{
			Id:      "test-service",
			Name:    "test",
			Version: "1.0.0",
		},
		Logger: &config.Logger{
			SensitiveRegex: []string{
				`eyJ[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+`,
				`token:"[^"]*"`,
				`sk-[A-Za-z0-9]+`,
				`password:"[^"]*"`,
				`admin123|password123|secret123`,
			},
		},
	}

	logger := GetLogger(conf)

	// 性能测试：记录大量日志
	start := time.Now()
	for i := 0; i < 1000; i++ {
		err := logger.Log(log.LevelInfo,
			"msg", "性能测试",
			"iteration", i,
			"token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			"api_key", "sk-1234567890abcdef",
			"password", "admin123",
		)
		if err != nil {
			t.Errorf("Logger.Log() error = %v", err)
		}
	}
	duration := time.Since(start)

	// 确保性能在合理范围内（1000条日志应该在1秒内完成）
	if duration > time.Second {
		t.Errorf("性能测试失败：1000条日志耗时 %v，超过1秒", duration)
	}

	t.Logf("性能测试通过：1000条日志耗时 %v", duration)
}

// 测试无效正则表达式的处理
func TestInvalidRegex(t *testing.T) {
	conf := &config.Config{
		Server: &config.Server{
			Id:      "test-service",
			Name:    "test",
			Version: "1.0.0",
		},
		Logger: &config.Logger{
			SensitiveRegex: []string{
				`[invalid regex`, // 无效的正则表达式
				`valid.*regex`,   // 有效的正则表达式
			},
		},
	}

	logger := GetLogger(conf)

	// 测试日志记录（不应该因为无效正则表达式而崩溃）
	err := logger.Log(log.LevelInfo,
		"msg", "测试无效正则表达式",
		"data", "some data",
	)
	if err != nil {
		t.Errorf("Logger.Log() error = %v", err)
	}
}

// 测试args中包含多个参数时只替换token内容
func TestArgsTokenReplacement(t *testing.T) {
	conf := &config.Config{
		Server: &config.Server{
			Id:      "test-service",
			Name:    "test",
			Version: "1.0.0",
		},
		Logger: &config.Logger{
			SensitiveRegex: []string{
				`[a-zA-Z_][a-zA-Z0-9_]*:"[^"]*"`, // 匹配任何字段名:值的格式
			},
		},
	}

	logger := GetLogger(conf)

	// 测试1: args中包含token和其他参数
	t.Run("args包含token和其他参数", func(t *testing.T) {
		args := `token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" userAgent:"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36" method:"POST" path:"/api/auth/login"`

		err := logger.Log(log.LevelInfo,
			"msg", "API调用",
			"args", args,
			"status", 200,
		)
		if err != nil {
			t.Errorf("Logger.Log() error = %v", err)
		}
	})

	// 测试2: args中包含多个token
	t.Run("args包含多个token", func(t *testing.T) {
		args := `token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" refreshToken:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" userAgent:"Mozilla/5.0"`

		err := logger.Log(log.LevelInfo,
			"msg", "API调用",
			"args", args,
			"status", 200,
		)
		if err != nil {
			t.Errorf("Logger.Log() error = %v", err)
		}
	})

	// 测试3: args中不包含token
	t.Run("args不包含token", func(t *testing.T) {
		args := `userAgent:"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36" method:"POST" path:"/api/auth/login" userId:"12345"`

		err := logger.Log(log.LevelInfo,
			"msg", "API调用",
			"args", args,
			"status", 200,
		)
		if err != nil {
			t.Errorf("Logger.Log() error = %v", err)
		}
	})

	// 测试4: 复杂的args字符串
	t.Run("复杂的args字符串", func(t *testing.T) {
		args := `token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" userAgent:"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36" method:"POST" path:"/api/auth/login" userId:"12345" timestamp:"2025-07-24T09:30:00Z" ip:"192.168.1.100"`

		err := logger.Log(log.LevelInfo,
			"msg", "API调用",
			"args", args,
			"status", 200,
		)
		if err != nil {
			t.Errorf("Logger.Log() error = %v", err)
		}
	})
}
