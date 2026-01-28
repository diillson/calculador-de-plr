# Bem-vindo ao Calculador de PLR! 🌟

## 🎉 Sobre o Projeto

Oi, pessoal! Aqui é onde a mágica acontece no mundo dos cálculos da PLR (Participação nos Lucros e Resultados). Este aplicativo, criado com Go, vai te surpreender com sua habilidade de calcular sua PLR num piscar de olhos. E adivinha? Ele é super fácil de usar, mas por trás da simplicidade, há um monte de tecnologia legal!

## 🚀 Funcionalidades

- **Cálculo da PLR**: Basta informar seu salário, multiplicador, porcentagem e meses trabalhados.
- **Logs Detalhados**: Com Logrus, você fica por dentro de tudo o que acontece durante o cálculo.
- **Interface Amigável**: Graças ao Cobra, nossa CLI é fácil até para quem nunca viu uma linha de código.

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
      
   Obs: no Linux e Mac será necessário permitir a execução do binário, basta passar o comando: chmod +x plr_calculator-OS
   caso abrir a mensagem de que o binario não pode ser executado por não haver desenvolvedor declarado,
   basta clicar com o botao direito e mandar executar com o seu terminal de preferencia,
   com isso nas proximas execuções poderá chamar o binario direto pelo terminal.
   ```
   (E veja o `plr_calculator` ganhar vida!)

## 🚀 Decolando com o Aplicativo

Depois de compilar, só rodar:

**🚀 Uso de Flags!**

Você pode usar as seguintes flags para fornecer dados ao aplicativo:

    --salario ou -s: Define o salário do funcionário.
    --multiplicador ou -m: Especifica o multiplicador da PLR.
    --participacao ou -p: Define a porcentagem de participação nos lucros.
    --meses ou -t: Informa o número de meses trabalhados.

## Exemplo de Uso com Flags

Execute o calculador da PLR com flags assim:
```sh

./plr_calculator --salario 7000 --multiplicador 2 --participacao 83 --meses 12
```

Ou de forma mais curta:
```sh

./plr_calculator -s 7000 -m 2 -p 83 -t 12
```

**Usando --help**

Quer saber mais sobre as flags? Simples! Use a flag --help para obter uma descrição detalhada de todas as opções disponíveis:

```sh

./plr_calculator --help
```

**Usando de forma interativa**

```sh
./plr_calculator
```

E seguir o que a tela te diz. Vai ser divertido, prometo!

## 🤝 Contribuições são Super Bem-vindas

Tem uma ideia brilhante? Encontrou um bugzinho? Contribuições são o coração do open source. Confira as `CONTRIBUTING.md` para saber como fazer parte desta aventura.

## 📄 Licença

Este tesouro está sob a licença [MIT](https://github.com/diillson/calculador-de-plr/blob/main/LICENSE). Veja `LICENSE` para detalhes.

## 📚 Aprenda Mais

Quer aprender mais sobre PLR, Go ou como criamos este app? Dá uma olhada nestes recursos:

- [Documentação do Go](https://golang.org/doc/)
- [Guia do Cobra](https://github.com/spf13/cobra#readme)
- [Tutoriais sobre PLR](#)

---

💡 **Dica Pro**: Explore o código, brinque com ele, quebre coisas e conserte-as. É a melhor maneira de aprender!

---

Espero que se divirta tanto explorando este projeto quanto eu me diverti criando-o. Qualquer coisa, estamos aqui! 🌈
