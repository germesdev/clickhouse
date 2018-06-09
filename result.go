package clickhouse

import (
	"fmt"
)

type result struct{}

func (*result) LastInsertId() (int64, error) {
	fmt.Printf("[NOTSUPPORTED] LastInsertId is not supported")
	return 0, nil
}
func (*result) RowsAffected() (int64, error) {
	fmt.Printf("[NOTSUPPORTED] RowsAffected is not supported")
	return 0, nil
}
