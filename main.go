package main

import (
	"calculadora-go/calculator"
	"calculadora-go/gui"
)

func main() {
	// Inicializa a calculadora
	calculator := calculator.NewCalculator()

	// Inicializa a interface gráfica
	gui.StartGUI(calculator)
}
