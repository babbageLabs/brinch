package core

import (
	"database/sql"
	"github.com/babbageLabs/brinch/bin"
	"github.com/babbageLabs/brinch/bin/transports"
)

const DBNamespace = "DB"

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
	routes    map[string]*Route
	sps       map[string][]StoredProcedureParam
	DB        *sql.DB
	transport transports.ITransport
}

func (sps *StoredProcedures) Register(route *Route) {
	route.NameSpace = DBNamespace

	if sps.routes == nil {
		sps.routes = map[string]*Route{route.Name: route}
	} else {
		sps.routes[route.Name] = route
	}
}

func (sps *StoredProcedures) Initialize(query string) error {
	rows, err := sps.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		sp := StoredProcedureParam{}
		err := rows.Scan(&sp.RoutineSchema,
			&sp.SpecificName,
			&sp.RoutineName,
			&sp.ParameterName,
			&sp.ParameterMode,
			&sp.DataType,
			&sp.UdtName,
			&sp.ParameterDefault,
		)
		if err != nil {
			return err
		}

		params, ok := sps.sps[sp.RoutineName]
		if ok {
			bin.Logger.Debug("Update routine ", sp.RoutineName, " with a new param ", sp.ParameterName)
			sps.sps[sp.RoutineName] = append(params, sp)
		} else {
			bin.Logger.Debug("Add new sp ", sp.RoutineName)
			if sps.sps == nil {
				sps.sps = map[string][]StoredProcedureParam{sp.RoutineName: {sp}}
			} else {
				sps.sps[sp.RoutineName] = append([]StoredProcedureParam{}, sp)
			}
		}
	}

	sps.SPsToRoutes()

	return nil
}

func (sps *StoredProcedures) SPsToRoutes() {
	for name, params := range sps.sps {
		bin.Logger.Debug("Creating new route from stored procedure ", name)
		route := &Route{
			NameSpace:        DBNamespace,
			Name:             name,
			Parameters:       nil,
			validateRequest:  true,
			validateResponse: true,
			transport:        sps.transport,
		}

		for _, param := range params {
			route.AddParam(&Param{
				Name:    param.ParameterName,
				Value:   nil,
				Type:    param.UdtName,
				err:     nil,
				Mode:    param.ParameterMode,
				isArray: param.DataType == "ARRAY",
			})
		}

		sps.Register(route)
	}
}
