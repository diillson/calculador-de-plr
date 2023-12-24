package validation

import (
	"errors"
	"github.com/diillson/calculador-de-plr/internal/domain"
)

func ValidarPLRInput(dados domain.PLRDados) error {
	if dados.Salario <= 0 {
		return errors.New("salário inválido")
	}
	// Adicionar mais validações conforme necessário
	return nil
}
