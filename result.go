package clickhouse

import (
	"fmt"
)

type result struct{}

func (*result) LastInsertId() (int64, error) {
	return 0, nil
}
func (*result) RowsAffected() (int64, error) {
	return 0, nil
}
