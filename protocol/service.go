package protocol

type Service interface {
	Start() error
	Stop() error
}
