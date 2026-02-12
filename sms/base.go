package sms

type request struct{}

func (*request) Category() string {
	return "message"
}

func (*request) Product() string {
	return "push"
}

func (*request) Function() string {
	return "sms"
}
