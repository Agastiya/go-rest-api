package Controller

type Controller struct {
	ControllerInterface
}

type ControllerInterface interface {
	ControllerService
	Account
	Auth
}
