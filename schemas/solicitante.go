package schemas

type Solicitante struct {
	Cnpj string `gorm:"primaryKey;unique;size:14"`
	NomeFantasia string `gorm:"unique"`
	Ativo bool `gorm:"default:true"`
	Cep string `gorm:"size:8"`
	Rua string 
	Numero string
	Cidade string
	Estado string
}