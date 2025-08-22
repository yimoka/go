// Package my ent mysql driver
package my

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect/sql"

	"entgo.io/ent/dialect"
	"github.com/yimoka/go/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// GetMySQLDriver _
func GetMySQLDriver(conf *config.Database) dialect.Driver {
	drv, err := sql.Open(conf.Driver, conf.Source)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	if !conf.IsTrace {
		return dialect.Driver(drv)
	}
	return dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		tracer := otel.Tracer("ent")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"mysql",
			trace.WithAttributes(
				attribute.String("db.statement", fmt.Sprint(i...)),
				attribute.String("db.system", "mysql"),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
}
