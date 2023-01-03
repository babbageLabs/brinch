package core

import "database/sql"

const DbNamespace = "DB"

type StoredProcedureParam struct {
	RoutineName      string
	RoutineSchema    string
	SpecificName     string
	ParameterName    string
	ParameterMode    ParameterMode
	DataType         string
	UdtName          string
	ParameterDefault sql.NullString
}

type StoredProcedures struct {
	routes map[string]*Route
	sps    map[string][]StoredProcedureParam
	DB     *sql.DB
}

func (sps *StoredProcedures) Register(route *Route) bool {
	_, ok := sps.routes[route.Name]
	if ok {
		return false
	}

	route.NameSpace = DbNamespace
	sps.routes[route.Name] = route
	return true
}

func (sps *StoredProcedures) Initialize(query string) error {
	rows, err := sps.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		sp := StoredProcedureParam{}
		err := rows.Scan(sp.RoutineSchema,
			sp.SpecificName,
			sp.RoutineName,
			sp.ParameterName,
			sp.ParameterMode,
			sp.DataType,
			sp.UdtName,
			sp.ParameterDefault,
		)
		if err != nil {
			return err
		}

		params, ok := sps.sps[sp.RoutineName]
		if ok {
			sps.sps[sp.RoutineName] = append(params, sp)
		} else {
			var sl []StoredProcedureParam
			sps.sps[sp.RoutineName] = append(sl, sp)
		}
	}

	return nil
}

func (sps *StoredProcedures) SPsToRoutes() {

}
