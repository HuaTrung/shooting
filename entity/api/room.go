package api

type Room struct {
	ID          uint `json:"id"`
	Name      	string 	`json:"name,omitempty"`
	HasPassword bool	`json:"has_password,omitempty"`
	Password 	string   `json:"is_requested"`
}

