package JsonSchema

import (
	"brinch/lib/utils"
	"github.com/jackc/pgx/v4"
)

type ToJsonSchema interface {
	GetQuery() string
	QueryHandler(rows pgx.Rows) utils.Schemas
}

func Export(exportable ToJsonSchema) {
	query := exportable.GetQuery()
	utils.QueryDB(&query, exportable.QueryHandler).Export()
}
