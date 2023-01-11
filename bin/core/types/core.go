package types

type ICallable interface {
	Validate(params *IValidatable) (bool, error)
	MustValidateRequest() bool
	MustValidateResponse() bool

	GetReqParams() Params
	GetResParams() Params

	GetSubject() (string, error) // ge the name of the topic in Nats
	GetTransport() (ITransport, error)
}
