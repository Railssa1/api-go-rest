package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	conexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	if DB, err = gorm.Open(postgres.Open(conexao)); err != nil {
		log.Panic("Erro ao conectar ao banco de dados.")
	}
}