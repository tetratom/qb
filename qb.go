package qb

import (
	"fmt"

	"github.com/tetratom/qb/internal"
)

type Dialect int

const (
	DialectDefault Dialect = iota
	DialectPq
	DialectGoracle
	DialectMssql
)

func Lit(s string, args ...interface{}) interface{} {
	return internal.Literal(fmt.Sprintf(s, args...))
}

func WithDialectPQ() Query {
	return DialectOption(DialectPq)
}

func WithDialectDefault() Query {
	return DialectOption(DialectDefault)
}

func WithDialectGoracle() Query {
	return DialectOption(DialectGoracle)
}

func WithDialectMssql() Query {
	return DialectOption(DialectMssql)
}
