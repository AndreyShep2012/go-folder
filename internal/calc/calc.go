package calc

import "github.com/AndreyShep2012/go-folder/internal/models"

//Calc -
func Calc(v models.Expression) int {
	return v.Var1 + v.Var2
}
