package models

type Person struct {
	ID        uint   `gorm:"primaryKey; not null" json:"id,omitempty" yaml:"id,omitempty"`
	Name      string `gorm:"type:varchar(200)" json:"name" yaml:"name"`
	Surname   string `gorm:"type:varchar(200)" json:"surname" yaml:"surname"`
	Age       int    `gorm:"age"json:"age" yaml:"age"`
	Email     string `gorm:"type:varchar(200)" json:"email" yaml:"email"`
	Telephone string `gorm:"type:varchar(200)" json:"telephone" yaml:"telephone"`
}
