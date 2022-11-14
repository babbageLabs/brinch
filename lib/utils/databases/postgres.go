package databases

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
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

//func (category PostgresTypeCategory) ToJsonType() JsonSchema.SchemaType {
//	switch category {
//	case ArrayType:
//		return JsonSchema.Array
//	case BooleanType:
//		return JsonSchema.Boolean
//	case UserDefineType:
//	case CompositeType:
//		return JsonSchema.Object
//	case DateType:
//	case EnumType:
//	case GeometricType:
//	case NetworkType:
//	case PseudoType:
//	case RangeType:
//	case BitStringType:
//	case TimespanType:
//	case StringType:
//		return JsonSchema.String
//	case NumericType:
//		return JsonSchema.Number
//	case UnknownType:
//		return JsonSchema.Null
//	}
//	return JsonSchema.Null
//}

type QueryHandler func(rows pgx.Rows, meta *DbMeta) (bool, error)

func QueryDB(query *string, handler QueryHandler) (bool, error) {
	dbUrl := viper.GetString("db.config.url")
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return false, err
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			logrus.Error(err)
		}
	}(conn, context.Background())

	rows, err := conn.Query(context.Background(), *query)
	if err != nil {
		return false, err
	}

	// rows.Close is called by rows.Next when all rows are read
	// or an error occurs in Next or Scan. So it may optionally be
	// omitted if nothing in the rows.Next loop can panic. It is
	// safe to close rows multiple times.
	defer rows.Close()

	return handler(rows, &DbMeta{
		SourceType: Postgres,
	})
}
