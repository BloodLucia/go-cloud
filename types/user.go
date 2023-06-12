package types

type UserServiceLoginReq struct {
	Identify string `json:"identify"`
	Password string `json:"password"`
}
