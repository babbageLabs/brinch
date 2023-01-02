package transports

type MetaData struct {
}

type Response struct {
	data []byte
	meta *MetaData
}

type Transport interface {
	Connect() (bool, error)
	Close() (bool, error)
	Exec(subject string, msg []byte, meta *MetaData) (*Response, error)
}
