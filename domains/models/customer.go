package domain_model

type CustomerLogged struct {
	Id    string   `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Roles []string `json:"roles,omitempty"`
}
