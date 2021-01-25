package dialect

import "reflect"

type Dialect interface {
	DataTypeOf(val reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

var dialectsMap = map[string]Dialect{}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, status bool) {
	dialect, status = dialectsMap[name]
	return
}
