package kernel

type Request interface {
	Category() string

	Product() string

	Function() string

	Url() string
}
