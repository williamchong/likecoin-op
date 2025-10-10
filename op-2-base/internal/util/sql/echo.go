package sql

import "fmt"

func Echo(sql string) string {
	return fmt.Sprintf("SELECT $$%s$$ as echo", sql)
}

func Stmt(sql string) string {
	return fmt.Sprintf("%s;", sql)
}
