package common

func GetVersion() string {
	return "v0.4.0"
}

type Functionality struct{}

func NewFunctionality() Functionality {
	return Functionality{}
}
