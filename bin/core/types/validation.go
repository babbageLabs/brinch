package types

type IValidatable interface {
	IsValid() bool
	Validate() (bool, error)
	Marshal() ([]byte, error)
}

type Validatables []IValidatable
