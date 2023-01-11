package database

import (
	"database/sql"
	"fmt"
	"github.com/babbageLabs/brinch/bin"
	"github.com/babbageLabs/brinch/bin/core/routing"
	"github.com/babbageLabs/brinch/bin/core/types"
)

const DBNamespace = "DB"

type StoredProcedureParam struct {
	RoutineName      string
	RoutineSchema    string
	SpecificName     string
	ParameterName    string
	ParameterMode    types.ParameterMode
	DataType         string
	UdtName          string
	ParameterDefault sql.NullString
}

type StoredProcedures struct {
	routes    map[string]*routing.Route
	sps       map[string][]StoredProcedureParam
	DB        *sql.DB
	transport *types.ITransport
}

func (sps *StoredProcedures) Register(route *routing.Route) {
	_ = route.SetNameSpace(DBNamespace)

	name, _ := route.GetName()

	if sps.routes == nil {
		sps.routes = map[string]*routing.Route{name: route}
	} else {
		sps.routes[name] = route
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

func (sps *StoredProcedures) SPsToRoutes() error {
	for name, params := range sps.sps {
		bin.Logger.Debug("Creating new route from stored procedure ", name)
		route := &routing.Route{
			NameSpace:        DBNamespace,
			Name:             name,
			Parameters:       nil,
			ValidateRequest:  true,
			ValidateResponse: true,
			Transport:        sps.transport,
		}

		for _, param := range params {
			_, err := route.AddParam(&routing.Param{
				Name:    param.ParameterName,
				Value:   nil,
				Type:    param.UdtName,
				Err:     nil,
				Mode:    param.ParameterMode,
				IsArray: param.DataType == "ARRAY",
			})
			if err != nil {
				return err
			}
		}

		sps.Register(route)
	}

	return nil
}

func (sps *StoredProcedures) SetTransport(transport types.ITransport) error {
	sps.transport = &transport

	return nil
}

func (sps *StoredProcedures) GetRoute(name string) (*routing.Route, error) {
	route, ok := sps.routes[name]
	if ok {
		return route, nil
	}

	return nil, fmt.Errorf(fmt.Sprintf("Route not found : %s", name))
}

func (sps *StoredProcedures) GetRoutes() map[string]*routing.Route {
	return sps.routes
}
