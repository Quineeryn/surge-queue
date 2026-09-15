package presenter

type LoginOutput struct {
	Token string `json:"token"`
}

func NewLoginOutput(token string) *LoginOutput {
	return &LoginOutput{
		Token: token,
	}
}
