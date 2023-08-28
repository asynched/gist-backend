package schemas

type SignInSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (schema *SignInSchema) IsValid() (bool, []string) {
	return true, []string{}
}
