package domain_model

type AuthResult struct {
	Token      string `json:"token,omitempty"`
	ExpireTime int    `json:"expireTime,omitempty"`
}
