package types

import (
	"github.com/gin-gonic/gin"
)

type Route interface {
	ICallable
	SetNameSpace(name string) error
	GetName() (string, error)
	AddParam(param Param) (bool, error)
	SetContext(context *gin.Context)
	GetMethod() string
	GetResponseCode() int
}

type ParameterMode string

type Param interface {
	GetName() (string, error)
	GetMode() (ParameterMode, error)
	IsInput() bool
	IsOutPut() bool
	IValidatable
}

type Params []Param
