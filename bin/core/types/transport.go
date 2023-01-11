package types

type MetaData struct {
	ResponseCode int
}

type Response struct {
	Data []byte
	Meta *MetaData
}

type ITransport interface {
	Connect() (bool, error)
	Close() (bool, error)
	Exec(subject string, msg []byte, meta *MetaData) (*Response, error)
}
