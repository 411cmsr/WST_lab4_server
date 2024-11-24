package models

type Person struct {
	ID        uint   `gorm:"primaryKey; not null" xml:"id" yaml:"id"`
	Name      string `gorm:"type:varchar(200)" xml:"name" yaml:"name"`
	Surname   string `gorm:"type:varchar(200)" xml:"surname" yaml:"surname"`
	Age       int    `gorm:"age" xml:"age" yaml:"age"`
	Email     string `gorm:"type:varchar(200)" xml:"email" yaml:"email"`
	Telephone string `gorm:"type:varchar(200)" xml:"telephone" yaml:"telephone"`
}
