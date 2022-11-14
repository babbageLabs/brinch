package JsonSchema

import (
	"brinch/lib/utils/databases"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ToJsonSchema interface {
	GetQuery() string
	QueryHandler(rows pgx.Rows, meta *databases.DbMeta) (bool, error)
	ToJsonSchema() (Schemas, error)
}

func Export(exportable ToJsonSchema) {
	dirPath := viper.GetString("jsonSchema.targetPath")
	jsonServerHost := viper.GetString("jsonSchema.server.jsonServerHost")

	query := exportable.GetQuery()
	_, err := databases.QueryDB(&query, exportable.QueryHandler)
	cobra.CheckErr(err)

	schemas, err := exportable.ToJsonSchema()
	cobra.CheckErr(err)

	schemas.Export(dirPath, jsonServerHost)
}
