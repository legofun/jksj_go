package main

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
)

func main() {
	log.Println(business())
}

func business() error {
	err := testErr()
	if err == sql.ErrNoRows {
		return err
	} else {
		return errors.Errorf("数据库查询失败, ", err)
	}
}

func testErr() error {
	return sql.ErrNoRows
}
