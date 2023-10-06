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

type SolicitanteResponse struct {
	Cnpj 	  		string 		`json:"cnpj"`
	NomeFantasia 	string 		`json:"nome_fantasia"`
	Ativo 			bool 		`json:"ativo"`
	Cep 			string 		`json:"cep"`
	Rua 			string 		`json:"rua"`
	Numero 			string 		`json:"numero"`
	Cidade 			string 		`json:"cidade"`
	Estado 			string 		`json:"estado"`
}