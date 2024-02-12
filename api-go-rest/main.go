package main

import (
	"api-rest/database"
	"api-rest/models"
	"api-rest/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciado o servidor Rest com Go")

	models.Personalidades = []models.Personalidade {
		{Id: 1, Nome: "Nome 1", Historia: "história 1"}, 
		{Id: 2, Nome: "Nome 2", Historia: "história 2"},
	}

	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
