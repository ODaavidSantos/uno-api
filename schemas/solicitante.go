package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Solicitante struct {
	gorm.Model
	Cnpj string `gorm:"unique;size:14"`
	NomeFantasia string `gorm:"unique"`
	Ativo bool `gorm:"default:true"`
	Cep string `gorm:"size:8"`
	Rua string 
	Numero string
	Cidade string
	Estado string
}

type SolicitanteResponse struct {
	ID        		uint      	`json:"id"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"updated_at"`
	DeletedAt 		time.Time 	`json:"deteledAt,omitempty"`
	Cnpj 	  		string 		`json:"cnpj"`
	NomeFantasia 	string 		`json:"nome_fantasia"`
	Ativo 			bool 		`json:"ativo"`
	Cep 			string 		`json:"cep"`
	Rua 			string 		`json:"rua"`
	Numero 			string 		`json:"numero"`
	Cidade 			string 		`json:"cidade"`
	Estado 			string 		`json:"estado"`
}