package databases

type DbMeta struct {
	SourceType SourceType
}

type SourceType string

const (
	Postgres SourceType = "postgres"
)
