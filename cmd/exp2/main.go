package main

import (
	"log"

	"github.com/AndreyShep2012/go-folder/internal/calc"
	"github.com/AndreyShep2012/go-folder/internal/models"
)

func main() {
	log.Println("start exp2")
	v := models.Expression{
		Var1: 20,
		Var2: 30,
	}

	log.Println(v.Var1, "+", v.Var2, "=", calc.Calc(v))
}
