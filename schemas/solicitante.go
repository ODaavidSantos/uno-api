package schemas

import ("gorm.io/gorm")

type Solicitante struct {
	gorm.Model
	Cnpj string
	NomeFantasia string
	Status string
	Cep string
	Rua string
	Numero string
	Cidade string
	Estado string
}