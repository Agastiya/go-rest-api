package Services

import "go-rest-api/Repository/auth"

type authService struct{}

var (
	Auth     AuthInterface    = &authService{}
	Account  AccountInterface = &authService{}
	repoAuth                  = &auth.AuthRepository{}
)
