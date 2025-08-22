// Package pg ent postgres driver
package pg

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

	// init postgres driver
	_ "github.com/lib/pq"
)

// GetPostgreSQLDriver _
func GetPostgreSQLDriver(conf *config.Database) dialect.Driver {
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
			"postgres",
			trace.WithAttributes(
				attribute.String("db.statement", fmt.Sprint(i...)),
				attribute.String("db.system", "postgres"),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
}
