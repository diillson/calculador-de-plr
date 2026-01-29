# Bem-vindo ao Calculador de PLR! 🌟

## 🎉 Sobre o Projeto

Oi, pessoal! Aqui é onde a mágica acontece no mundo dos cálculos da PLR (Participação nos Lucros e Resultados). Este aplicativo, criado com Go, vai te surpreender com sua habilidade de calcular sua PLR num piscar de olhos. E adivinha? Ele é super fácil de usar, mas por trás da simplicidade, há um monte de tecnologia legal!

## 🚀 Funcionalidades

- **Cálculo da PLR**: Basta informar seu salário, multiplicador, porcentagem e meses trabalhados.
- **Cálculo de IRPF**: Calcula automaticamente o Imposto de Renda sobre a PLR com as tabelas 2025.
- **Escolha de Período**: Selecione entre as tabelas de Janeiro-Abril/2025 ou Maio+/2025.
- **Logs Detalhados**: Com Logrus, você fica por dentro de tudo o que acontece durante o cálculo.
- **Interface Amigável**: Graças ao Cobra, nossa CLI é fácil até para quem nunca viu uma linha de código.

## 🐠 Tabelas IRPF 2025

O calculador utiliza as tabelas oficiais de IRPF para PLR:

### Janeiro a Abril/2025

| Faixa | Alíquota | Dedução |
|--------|----------|----------|
| Até R$ 7.640,00 | 0% (isento) | R$ 0,00 |
| R$ 7.640,01 a R$ 9.922,28 | 7,5% | R$ 573,06 |
| R$ 9.922,29 a R$ 13.167,00 | 15% | R$ 1.317,23 |
| R$ 13.167,01 a R$ 16.380,38 | 22,5% | R$ 2.304,76 |
| Acima de R$ 16.380,38 | 27,5% | R$ 3.123,78 |

### A partir de Maio/2025 ✅ (padrão)

| Faixa | Alíquota | Dedução |
|--------|----------|----------|
| Até R$ 8.214,40 | 0% (isento) | R$ 0,00 |
| R$ 8.214,41 a R$ 9.922,28 | 7,5% | R$ 616,08 |
| R$ 9.922,29 a R$ 13.167,00 | 15% | R$ 1.360,25 |
| R$ 13.167,01 a R$ 16.380,38 | 22,5% | R$ 2.347,78 |
| Acima de R$ 16.380,38 | 27,5% | R$ 3.166,80 |

## 🛠 Tecnologias Usadas

- **[Cobra](https://github.com/spf13/cobra)**: Uma biblioteca para criar CLIs poderosas.
- **[Logrus](https://github.com/sirupsen/logrus)**: Uma ferramenta de logging com superpoderes.
- **Go Padrão**: Para todas as outras coisas que tornam esse app possível.

## 🏗 Como Montar seu Ambiente

Você vai precisar do Go (1.24+), Git e uma xícara de café (ou chá, se preferir).

### 📋 Passo a Passo para a Diversão

1. **Clone este repositório:**
   ```sh
   git clone https://github.com/diillson/calculador-de-plr.git
   ```
2. **Entre no nosso mundo:**
   ```sh
   cd calculador-de-plr
   ```
   (Isso é onde a mágica começa.)

3. **Compile o projeto:**
   ```sh
   go build -o plr_calculator ./cmd
   ```
   (E veja o `plr_calculator` ganhar vida!)

4. **Teve problemas ou dificuldades no build? Relaxa tenho a solução para você :**
   ```sh
   Acesse o link: https://github.com/diillson/calculador-de-plr/releases
   na ultima Release, terá já buildado pelo workflow,versões para Mac(Darwin), Linux e Windows.
   Basta baixar, extrair e executar em seu terminal Linux e Mac ou Prompt/PowerShell sendo Windows.
       
   Obs: no Linux e Mac será necessário permitir a execução do binário, basta passar o comando: chmod +x *
   caso abrir a mensagem de que o binario não pode ser executado por não haver desenvolvedor declarado,
   basta clicar com o botao direito e mandar executar com o seu terminal de preferencia,
   com isso nas proximas execuções poderá chamar o binario direto pelo terminal.
   ```
   (E veja o `plr_calculator` ganhar vida!)

## 🚀 Decolando com o Aplicativo

Depois de compilar, só rodar:

### 🚀 Uso de Flags!

Você pode usar as seguintes flags para fornecer dados ao aplicativo:

    --salario ou -s: Define o salário do funcionário.
    --multiplicador ou -m: Especifica o multiplicador da PLR.
    --participacao ou -p: Define a porcentagem de participação nos lucros.
    --meses ou -t: Informa o número de meses trabalhados.
    --periodo: Escolhe a tabela IRPF: jan-abr ou maio-mais (padrão).

### Exemplo de Uso com Flags

**Usando tabela Maio+/2025 (padrão):**
```sh
./plr_calculator -s 7000 -m 2 -p 83 -t 12
```

**Usando tabela Janeiro-Abril/2025:**
```sh
./plr_calculator -s 7000 -m 2 -p 83 -t 12 --periodo jan-abr
```

**Forma curta:**
```sh
./plr_calculator --salario 7000 --multiplicador 2 --participacao 83 --meses 12 --periodo maio-mais
```

### **Usando --help**

Quer saber mais sobre as flags? Simples! Use a flag --help para obter uma descrição detalhada de todas as oqções disponíveis:

```sh
./plr_calculator --help
```

### **Usando de forma interativa**

```sh
./plr_calculator
```

E seguir o que a tela te diz. Você poderá escolher o período da tabela IRPF:

```
Digite o salário: 7000
Digite o multiplicador da PLR: 2
Digite a porcentagem de participação nos lucros: 83
Digite o número de meses trabalhados: 12
Escolha o período da tabela IRPF:
1) Janeiro a Abril/2025
2) A partir de Maio/2025 (padrão)
Digite sua escolha (1 ou 2): 2
```

## 🤤 Contribuições são Super Bem-vindas

Tem uma ideia brilhante? Encontrou um bugzinho? Contribuições são o coração do open source. Confira as `CONTRIBUTING.md` para saber como fazer parte desta aventura.

## 📄 Licença

Este tesouro está sob a licença [MIT](https://github.com/diillson/calculador-de-plr/blob/main/LICENSE). Veja `LICENSE` para detalhes.

## 📺 Aprenda Mais

Quer aprender mais sobre PLR, Go ou como criamos este app? Dá uma olhada nestes recursos:

- [Documentação do Go](https://golang.org/doc/)
- [Guia do Cobra](https://github.com/spf13/cobra#readme)
- [Tutoriais sobre PLR](#)

---

💡 **Dica Pro**: Explore o código, brinque com ele, quebre coisas e conserte-as. É a melhor maneira de aprender!

---

Espero que se divirta tanto explorando este projeto quanto eu me diverti criando-o. Qualquer coisa, estamos aqui! 🌈
