package finance

import (
	"github.com/diillson/calculador-de-plr/internal/domain"
	"testing"
)

func TestCalculator_CalcularPLR(t *testing.T) {
	calculator := NewCalculator()

	tests := []struct {
		name    string
		dados   domain.PLRDados
		wantPLR float64
		wantErr bool
	}{
		{
			name: "Caso padrão",
			dados: domain.PLRDados{
				Salario:                 10000,
				Multiplicador:           2,
				PorcentagemParticipacao: 0.8,
				MesesTrabalhados:        12,
			},
			wantPLR: 16000,
			wantErr: false,
		},
		// Adicionar mais casos de teste conforme necessário
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPLR, err := calculator.CalcularPLR(tt.dados)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculator.CalcularPLR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPLR != tt.wantPLR {
				t.Errorf("Calculator.CalcularPLR() = %v, want %v", gotPLR, tt.wantPLR)
			}
		})
	}
}
