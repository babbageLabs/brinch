package core

import "testing"

func TestStoredProcedure_Resolve(t *testing.T) {
	sp := &StoredProcedureParam{
		Engine: "postgres",
		Route: Route{
			NameSpace:        DbNamespace,
			Name:             "route1",
			Parameters:       nil,
			validateRequest:  false,
			validateResponse: false,
			ITransport:       nil,
		},
	}
}
