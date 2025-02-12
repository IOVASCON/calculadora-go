package calculator

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Calculator é a estrutura que armazena o estado da calculadora
type Calculator struct{}

// NewCalculator cria uma nova instância da calculadora
func NewCalculator() *Calculator {
	return &Calculator{}
}

// Add realiza a soma de dois números e retorna o resultado formatado
func (c *Calculator) Add(a, b float64) string {
	result := a + b
	return formatNumber(result)
}

// Subtract realiza a subtração de dois números e retorna o resultado formatado
func (c *Calculator) Subtract(a, b float64) string {
	result := a - b
	return formatNumber(result)
}

// Multiply realiza a multiplicação de dois números e retorna o resultado formatado
func (c *Calculator) Multiply(a, b float64) string {
	result := a * b
	return formatNumber(result)
}

// Divide realiza a divisão de dois números e retorna o resultado formatado
func (c *Calculator) Divide(a, b float64) (string, error) {
	if b == 0 {
		return "", fmt.Errorf("divisão por zero não permitida")
	}
	result := a / b
	return formatNumber(result), nil
}

// formatNumber formata um número no padrão brasileiro
func formatNumber(value float64) string {
	p := message.NewPrinter(language.Portuguese)
	return p.Sprintf("%.2f", value)
}
