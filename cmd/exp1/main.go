package main

import (
	"log"

	"github.com/AndreyShep2012/go-folder/internal/calc"
	"github.com/AndreyShep2012/go-folder/internal/models"
)

func main() {
	log.Println("start exp1")
	v := models.Expression{
		Var1: 10,
		Var2: 20,
	}

	log.Println(v.Var1, "+", v.Var2, "=", calc.Calc(v))
}
