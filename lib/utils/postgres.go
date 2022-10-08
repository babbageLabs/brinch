package utils

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type PostgresTypeCategory string

const (
	ArrayType      PostgresTypeCategory = "A"
	BooleanType                         = "B"
	CompositeType                       = "C"
	DateType                            = "D"
	EnumType                            = "E"
	GeometricType                       = "G"
	NetworkType                         = "I"
	NumericType                         = "N"
	PseudoType                          = "P"
	RangeType                           = "R"
	StringType                          = "S"
	TimespanType                        = "T"
	UserDefineType                      = "U"
	BitStringType                       = "V"
	UnknownType                         = "X"
)

func (category PostgresTypeCategory) ToJsonType() SchemaType {
	switch category {
	case ArrayType:
		return Array
	case BooleanType:
		return BooleanType
	case UserDefineType:
	case CompositeType:
		return Object
	case DateType:
	case EnumType:
	case GeometricType:
	case NetworkType:
	case PseudoType:
	case RangeType:
	case BitStringType:
	case TimespanType:
	case StringType:
		return String
	case NumericType:
		return Number
	case UnknownType:
		return Null
	}
	return Null
}

type QueryHandler func(rows pgx.Rows) Schemas

func QueryDB(query *string, handler QueryHandler) Schemas {
	dbUrl := viper.GetString("db.config.url")
	conn, err := pgx.Connect(context.Background(), dbUrl)
	cobra.CheckErr(err)
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			logrus.Error(err)
		}
	}(conn, context.Background())

	rows, err := conn.Query(context.Background(), *query)
	cobra.CheckErr(err)

	// rows.Close is called by rows.Next when all rows are read
	// or an error occurs in Next or Scan. So it may optionally be
	// omitted if nothing in the rows.Next loop can panic. It is
	// safe to close rows multiple times.
	defer rows.Close()

	return handler(rows)
}
