package schemas

type Usuario struct {
	Cpf       string `gorm:"primaryKey"`
	Nome      string
	Sobrenome string
	Email     string `gorm:"unique"`
	Senha     string 
	Cargo     string
}