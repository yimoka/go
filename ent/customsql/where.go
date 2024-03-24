// Package customsql where
package customsql

import (
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/samber/lo"
)

// GetStrSliceAndQuery 获取字符串切片的并查询
func GetStrSliceAndQuery(field string, values ...string) func(s *sql.Selector) {
	if len(values) == 0 {
		return func(_ *sql.Selector) {}
	}

	return func(s *sql.Selector) {
		if isPostgres(s) {
			//nolint:goconst
			str := `["` + strings.Join(values, `","`) + `"]`
			s.Where(sql.P(func(b *sql.Builder) {
				b.Ident(field).WriteString(" @> ").Arg(str)
			}))
		} else {
			str := `"` + strings.Join(values, `","`) + `"`
			s.Where(sql.P(func(b *sql.Builder) {
				b.WriteString("JSON_CONTAINS").Wrap(func(b *sql.Builder) {
					b.Ident(field).Comma()
					//nolint:goconst
					b.WriteString(`JSON_ARRAY(` + str + `)`)
				})
			}))
		}
	}
}

// GetStrSliceOrQuery 获取字符串切片的或查询
func GetStrSliceOrQuery(field string, values ...string) func(s *sql.Selector) {
	if len(values) == 0 {
		return func(_ *sql.Selector) {}
	}
	return func(s *sql.Selector) {
		if isPostgres(s) {
			pArr := lo.Map(values, func(val string, _ int) *sql.Predicate { return getPgJSONStrQuery(field, val) })
			s.Where(sql.Or(pArr...))
			return
		}
		pArr := lo.Map(values, func(val string, _ int) *sql.Predicate { return getMysqlJSONStrQuery(field, val) })
		s.Where(sql.Or(pArr...))
	}
}

func getMysqlJSONStrQuery(field string, value string) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		b.WriteString("JSON_CONTAINS").Wrap(func(b *sql.Builder) {
			b.Ident(field).Comma()
			b.WriteString(`'"` + value + `"'`)
			b.WriteString(`,"$"`)
		})
	})
}

func getPgJSONStrQuery(field string, value string) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		b.Ident(field).WriteString(" @> ").Arg(`["` + value + `"]`)
	})
}

// GetIntSliceAndQuery 获取整数切片的并查询
func GetIntSliceAndQuery(field string, values ...int) func(s *sql.Selector) {
	if len(values) == 0 {
		return func(_ *sql.Selector) {}
	}
	return func(s *sql.Selector) {
		if isPostgres(s) {
			str := "[" + strings.Join(lo.Map(values, func(val int, _ int) string { return strconv.Itoa(val) }), ",") + "]"
			s.Where(sql.P(func(b *sql.Builder) {
				b.Ident(field).WriteString(" @> ").Arg(str)
			}))
		} else {
			str := strings.Join(lo.Map(values, func(val int, _ int) string { return strconv.Itoa(val) }), `,`)
			s.Where(sql.P(func(b *sql.Builder) {
				b.WriteString("JSON_CONTAINS").Wrap(func(b *sql.Builder) {
					b.Ident(field).Comma()
					b.WriteString(`JSON_ARRAY(` + str + `)`)
				})
			}))
		}
	}
}

// GetIntSliceOrQuery 获取整数切片的或查询
func GetIntSliceOrQuery(field string, values ...int) func(s *sql.Selector) {
	if len(values) == 0 {
		return func(_ *sql.Selector) {}
	}
	return func(s *sql.Selector) {
		if isPostgres(s) {
			pArr := lo.Map(values, func(val int, _ int) *sql.Predicate { return getPgJSONIntQuery(field, val) })
			s.Where(sql.Or(pArr...))
			return
		}
		pArr := lo.Map(values, func(val int, _ int) *sql.Predicate { return getMyJSONIntQuery(field, val) })
		s.Where(sql.Or(pArr...))
	}
}

func getMyJSONIntQuery(field string, value int) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		b.WriteString("JSON_CONTAINS").Wrap(func(b *sql.Builder) {
			b.Ident(field).Comma()
			b.WriteString(`"` + strconv.Itoa(value) + `"`)
			b.WriteString(`,"$"`)
		})
	})
}

func getPgJSONIntQuery(field string, value int) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		b.Ident(field).WriteString(" @> ").Arg(`[` + strconv.Itoa(value) + `]`)
	})
}

func isPostgres(s *sql.Selector) bool {
	return s.Dialect() == "postgres"
}
