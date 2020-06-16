package main

import (
	"fmt"
)

func main() {

	titulo := `
	*******************
	    CALCULADORA
	*******************`
	fmt.Println(titulo)
	fmt.Println("Quieres hacer una operacion?")
	fmt.Println(`
	1-> Suma		
	2-> Resta
	3-> Multiplicación
	4-> División
	`)
	fmt.Println("Cual deseas realizar?")
	var operacion int
	fmt.Scanln(&operacion)
	var numero int
	var numero1 int

	// Suma

	if operacion == 1 {
		fmt.Println("Dime el primer valor")
		fmt.Scan(&numero)
		fmt.Println("Indicame el sugundo valor")
		fmt.Scan(&numero1)
		animacion := `
	***********************
	*resultado de la Suma *
	***********************	`
		fmt.Println(animacion)
		resultado := numero + numero1
		fmt.Println(resultado)
	}

	//Resta

	if operacion == 2 {
		fmt.Println("Dime el primer valor")
		fmt.Scan(&numero)
		fmt.Println("Indicame el sugundo valor")
		fmt.Scan(&numero1)
		animacion := `
	***********************
	*resultado de la Resta *
	***********************	`
		fmt.Println(animacion)
		resultado := numero - numero1
		fmt.Println(resultado)
	}

	// Multiplicación

	if operacion == 3 {
		fmt.Println("Dime el primer valor")
		fmt.Scan(&numero)
		fmt.Println("Indicame el sugundo valor")
		fmt.Scan(&numero1)
		animacion := `
	***********************
	*resultado de la Multiplicación *
	***********************	`
		fmt.Println(animacion)
		resultado := numero * numero1
		fmt.Println(resultado)
	}

	// División

	if operacion == 4 {
		fmt.Println("Dime el primer valor")
		fmt.Scan(&numero)
		fmt.Println("Indicame el sugundo valor")
		fmt.Scan(&numero1)
		animacion := `
	***********************
	*resultado de la  División *
	***********************	`
		fmt.Println(animacion)
		resultado := numero / numero1
		fmt.Println(resultado)
	}

	fmt.Println("Aplicación de un solo uso! Presione cualquier numero para salir!\n")
	var salir int
	fmt.Scan(&salir)
	fmt.Println("Adiooooooos!!!")
}
