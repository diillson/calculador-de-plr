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
		wantPLR string
		wantErr bool
	}{
		{
			name: "Caso padrão",
			dados: domain.PLRDados;
				Salario:                 10000,
				Multiplicador:           2,
				PorcentagemParticipacao: 0.8,
				MesesTrabalhados:        12,
			},
			wantPLR: "16,000.00",
			wantErr: false,
		},
		{
			name: "PLR com porcentagem maior que 1",
			dados: domain.PLRDados{
				Salario:                 7000,
				Multiplicador:           2,
				PorcentagemParticipacao: 83, // Será convertido para 0.83
				MesesTrabalhados:        12,
			},
			wantPLR: "11,620.00",
			wantErr: false,
		},
		{
			name: "PLR com meses parciais",
			dados: domain.PLRDados;
				Salario:                 5000,
				Multiplicador:           1.5,
				PorcentagemParticipacao: 0.9,
				MesesTrabalhados:        6,
			},
			wantPLR: "3,375.00",
			wantErr: false,
		},
		{
			name: "Valores inválidos - salário negativo",
			dados: domain.PLRDados{
				Salario:                 -1000,
				Multiplicador:           2,
				PorcentagemParticipacao: 0.8,
				MesesTrabalhados:        12,
			},
			wantPLR: "",
			wantErr: true,
		},
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

func TestCalculator_CalcularIRPF(t *testing.T) {
	calculator := NewCalculator()
	tabelaIRPF := domain.TabelaIRPF()

	tests := []struct {
		name           string
		plr            float64
		wantAliquota  float64
		wantImposto  float64
		wantLiquido  float64
		wantErr       bool
	}{
		{
			name:           "Isento - PLR até R$ 28.467,20",
			plr:            15000.00,
			wantAliquota:  0.0,
			wantImposto:  0.0,
			wantLiquido:  15000.00,
			wantErr:       false,
		},
		{
			name:           "Primeira faixa - 7,5%",
			plr:            30000.00,
			wantAliquota:  7.5,
			wantImposto:  115.04,
			wantLiquido:  29884.96,
			wantErr:       false,
		},
		{
			name:           "Segunda faixa - 15%",
			plr:            40000.00,
			wantAliquota:  15.0,
			wantImposto:  1320.97,
			wantLiquido:  38679.03,
			wantErr:       false,
		},
		{
			name:           "Terceira faixa - 22,5%",
			plr:            50000.00,
			wantAliquota:  22.5,
			wantImposto:  3195.03,
			wantLiquido:  46804.97,
			wantErr:       false,
		},
		{
			name:           "Quarta faixa - 27,5%",
			plr:            60000.00,
			wantAliquota:  27.5,
			wantImposto:  5646.22,
			wantLiquido:  54353.78,
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResultado, err := calculator.CalcularIRPF(tt.plr, tabelaIRPF)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculator.CalcularIRPF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if gotResultado.Aliquota != tt.wantAliquota {
					t.Errorf("Alíquota = %.1f, want %.1f", gotResultado.Aliquota, tt.wantAliquota)
				}
				if gotResultado.ImpostoApurado < tt.wantImposto-0.01 || gotResultado.ImpostoApurado > tt.wantImposto+0.01 {
					t.Errorf("Imposto Apurado = %.2f, want %.2f", gotResultado.ImpostoApurado, tt.wantImposto)
				}
				if gotResultado.ValorLiquido < tt.wantLiquido-0.01 || gotResultado.ValorLiquido > tt.wantLiquido+0.01 {
					t.Errorf("Valor Líquido = %.2f, want %.2f", gotResultado.ValorLiquido, tt.wantLiquido)
				}
			}
		})
	}
}
