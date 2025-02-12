package calculator

import "testing"

func TestAdd(t *testing.T) {
	calc := NewCalculator()
	result := calc.Add(2, 3)
	if result != "5,00" { // Resultado formatado no padrão brasileiro
		t.Errorf("Erro na soma: esperado %v, obtido %v", "5,00", result)
	}
}

func TestSubtract(t *testing.T) {
	calc := NewCalculator()
	result := calc.Subtract(5, 3)
	if result != "2,00" { // Resultado formatado no padrão brasileiro
		t.Errorf("Erro na subtração: esperado %v, obtido %v", "2,00", result)
	}
}

func TestMultiply(t *testing.T) {
	calc := NewCalculator()
	result := calc.Multiply(4, 3)
	if result != "12,00" { // Resultado formatado no padrão brasileiro
		t.Errorf("Erro na multiplicação: esperado %v, obtido %v", "12,00", result)
	}
}

func TestDivide(t *testing.T) {
	calc := NewCalculator()
	result, err := calc.Divide(10, 2)
	if err != nil || result != "5,00" { // Resultado formatado no padrão brasileiro
		t.Errorf("Erro na divisão: esperado %v, obtido %v", "5,00", result)
	}

	_, err = calc.Divide(10, 0)
	if err == nil {
		t.Errorf("Erro esperado na divisão por zero, mas nenhum erro foi retornado")
	}
}
