package models

type AddPersonRequest struct {
	Name      string `json:"name" binding:"required"`
	Surname   string `json:"surname" binding:"required"`
	Age       int    `json:"age" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Telephone string `json:"telephone" binding:"required"`
}

type UpdatePersonRequest struct {
	Name      string `json:"name,omitempty"`
	Surname   string `json:"surname,omitempty"`
	Age       int    `json:"age,omitempty"`
	Email     string `json:"email,omitempty"`
	Telephone string `json:"telephone,omitempty"`
}
