// Package customsql where
package customsql

import (
	"fmt"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/samber/lo"
)

// GetStrSliceAndQuery 获取字符串切片的并查询
func GetStrSliceAndQuery(field string, values ...string) func(s *sql.Selector) {
	if len(values) == 0 {
		return func(s *sql.Selector) {}
	}
	str := `"` + strings.Join(values, `","`) + `"`
	return func(s *sql.Selector) {
		s.Where(sql.P(func(b *sql.Builder) {
			b.WriteString("JSON_CONTAINS").Wrap(func(b *sql.Builder) {
				b.Ident(field).Comma()
				b.WriteString(`JSON_ARRAY(` + str + `)`)
			})
		}))
	}
}

// GetStrSliceOrQuery 获取字符串切片的或查询
func GetStrSliceOrQuery(field string, values ...string) func(s *sql.Selector) {
	if len(values) == 0 {
		return func(s *sql.Selector) {}
	}
	return func(s *sql.Selector) {
		pArr := lo.Map(values, func(val string, i int) *sql.Predicate { return getJSONStrQuery(field, val) })
		s.Where(sql.Or(pArr...))
	}
}

func getJSONStrQuery(field string, value string) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		b.WriteString("JSON_CONTAINS").Wrap(func(b *sql.Builder) {
			b.Ident(field).Comma()
			b.WriteString(`'"` + value + `"'`)
			b.WriteString(`,"$"`)
		})
	})
}

// GetIntSliceAndQuery 获取整数切片的并查询
func GetIntSliceAndQuery(field string, values ...int) func(s *sql.Selector) {
	fmt.Printf("values %+v \n", values)
	if len(values) == 0 {
		return func(s *sql.Selector) {}
	}
	str := strings.Join(lo.Map(values, func(val int, i int) string { return strconv.Itoa(val) }), `,`)
	return func(s *sql.Selector) {
		s.Where(sql.P(func(b *sql.Builder) {
			b.WriteString("JSON_CONTAINS").Wrap(func(b *sql.Builder) {
				b.Ident(field).Comma()
				b.WriteString(`JSON_ARRAY(` + str + `)`)
			})
		}))
	}
}

// GetIntSliceOrQuery 获取整数切片的或查询
func GetIntSliceOrQuery(field string, values ...int) func(s *sql.Selector) {
	if len(values) == 0 {
		return func(s *sql.Selector) {}
	}
	return func(s *sql.Selector) {
		pArr := lo.Map(values, func(val int, i int) *sql.Predicate { return getJSONIntQuery(field, val) })
		s.Where(sql.Or(pArr...))
	}
}

func getJSONIntQuery(field string, value int) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		b.WriteString("JSON_CONTAINS").Wrap(func(b *sql.Builder) {
			b.Ident(field).Comma()
			b.WriteString(`"` + strconv.Itoa(value) + `"`)
			b.WriteString(`,"$"`)
		})
	})
}
