package kernel

type ExceptionHandler interface {
	Handle(int, int, any) error
}
