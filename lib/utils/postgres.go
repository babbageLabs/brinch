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
	BooleanType    PostgresTypeCategory = "B"
	CompositeType  PostgresTypeCategory = "C"
	DateType       PostgresTypeCategory = "D"
	EnumType       PostgresTypeCategory = "E"
	GeometricType  PostgresTypeCategory = "G"
	NetworkType    PostgresTypeCategory = "I"
	NumericType    PostgresTypeCategory = "N"
	PseudoType     PostgresTypeCategory = "P"
	RangeType      PostgresTypeCategory = "R"
	StringType     PostgresTypeCategory = "S"
	TimespanType   PostgresTypeCategory = "T"
	UserDefineType PostgresTypeCategory = "U"
	BitStringType  PostgresTypeCategory = "V"
	UnknownType    PostgresTypeCategory = "X"
)

func (category PostgresTypeCategory) ToJsonType() SchemaType {
	switch category {
	case ArrayType:
		return Array
	case BooleanType:
		return Boolean
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

type QueryHandler func(rows pgx.Rows) (bool, error)

func QueryDB(query *string, handler QueryHandler) (bool, error) {
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
