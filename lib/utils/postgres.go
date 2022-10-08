package utils

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type QueryHandler func(rows pgx.Rows)

func QueryDB(query *string, handler QueryHandler) {
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

	handler(rows)
}
