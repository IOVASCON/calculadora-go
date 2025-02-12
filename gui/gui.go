package gui

import (
	"calculadora-go/calculator"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// StartGUI inicia a interface gráfica da calculadora
func StartGUI(calc *calculator.Calculator) {
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculadora")

	// Campos de entrada
	inputA := widget.NewEntry()
	inputB := widget.NewEntry()
	result := widget.NewLabel("Resultado aparecerá aqui")

	// Botões de operação
	addButton := widget.NewButton("Somar", func() {
		a, _ := parseInput(inputA.Text)
		b, _ := parseInput(inputB.Text)
		result.SetText(fmt.Sprintf("Resultado: %s", calc.Add(a, b))) // Usando resultado formatado
	})

	subtractButton := widget.NewButton("Subtrair", func() {
		a, _ := parseInput(inputA.Text)
		b, _ := parseInput(inputB.Text)
		result.SetText(fmt.Sprintf("Resultado: %s", calc.Subtract(a, b))) // Usando resultado formatado
	})

	multiplyButton := widget.NewButton("Multiplicar", func() {
		a, _ := parseInput(inputA.Text)
		b, _ := parseInput(inputB.Text)
		result.SetText(fmt.Sprintf("Resultado: %s", calc.Multiply(a, b))) // Usando resultado formatado
	})

	divideButton := widget.NewButton("Dividir", func() {
		a, _ := parseInput(inputA.Text)
		b, _ := parseInput(inputB.Text)
		res, err := calc.Divide(a, b)
		if err != nil {
			result.SetText(fmt.Sprintf("Erro: %s", err.Error()))
		} else {
			result.SetText(fmt.Sprintf("Resultado: %s", res)) // Usando resultado formatado
		}
	})

	// Layout
	content := container.NewVBox(
		widget.NewLabel("Número A:"),
		inputA,
		widget.NewLabel("Número B:"),
		inputB,
		addButton,
		subtractButton,
		multiplyButton,
		divideButton,
		result,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}

// parseInput converte a entrada de texto para float64
func parseInput(input string) (float64, error) {
	var value float64
	_, err := fmt.Sscanf(input, "%f", &value)
	return value, err
}
