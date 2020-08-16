package main

import (
	"database/sql"
	"errors"
)

func main() {
	CoverAndCope(ErrorGenerator(1))
	CoverAndCope(ErrorGenerator(-1))
}

func ErrorGenerator(x int) error {
	if x < 0 {
		return sql.ErrNoRows
	} else {
		return errors.New("an error i created")
	}
}

func CoverAndCope(err error) {
	defer func() {
		if eil := recover(); eil != nil {
			switch eil.(type) {

			}
		}
	}()
	panic(err)
}
