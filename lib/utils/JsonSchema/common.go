package JsonSchema

import (
	"brinch/lib/utils"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/cobra"
)

type ToJsonSchema interface {
	GetQuery() string
	QueryHandler(rows pgx.Rows, meta *utils.DbMeta) (bool, error)
	ToJsonSchema() (utils.Schemas, error)
}

func Export(exportable ToJsonSchema) {
	query := exportable.GetQuery()
	_, err := utils.QueryDB(&query, exportable.QueryHandler)
	cobra.CheckErr(err)

	schemas, err := exportable.ToJsonSchema()
	cobra.CheckErr(err)

	schemas.Export()
}
