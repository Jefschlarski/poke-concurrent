# PokeConcurrent

**PokeConcurrent** é um pequeno projeto de teste desenvolvido em Go que utiliza a API do Pokémon para buscar informações sobre Pokémon de forma concorrente. Este projeto visa simplesmente demonstra a utilização de conceitos fundamentais de concorrência em Go, como goroutines, channels, mutexes, wait groups e operações atômicas de forma simples e descomplicada para fins de estudo.

## Conceitos Utilizados

1. **Goroutines**: Funções que são executadas concorrentemente, permitindo que múltiplas chamadas à API sejam realizadas ao mesmo tempo, otimizando o desempenho.

2. **Channels**: Mecanismos de comunicação entre goroutines que facilitam a troca de dados de forma segura e eficiente.

3. **Sincronização**:
   - **Mutexes (Locks)**: Controle de acesso a recursos compartilhados, prevenindo inconsistências nos dados.
   - **WaitGroups**: Sincronização de goroutines, garantindo que todas as tarefas em execução sejam concluídas antes que o programa continue.

4. **Atomicidade**: Garantia da integridade de variáveis compartilhadas, permitindo operações seguras em ambientes concorrentes.

## Funcionalidade

O sistema busca informações sobre Pokémon, separando-os por tipo e gerando um relatório. O uso de concorrência melhora a eficiência do sistema e torna o gerenciamento de tarefas simultâneas mais intuitivo.

## Como Executar o Projeto

### Pré-requisitos

- **Go**: Certifique-se de que você tenha o Go instalado. Você pode baixar e instalar o Go a partir do site oficial: [golang.org](https://golang.org/dl/).

### Passos para Execução

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/Jefschlarski/poke-concurrent
   cd poke-concurrent
   ```

2. **Compile e execute o projeto**:
   ```bash
   go run main.go
   ```

### Resultados Esperados

Após a execução, o programa fará chamadas à API do Pokémon e exibirá um relatório no console com os Pokémon separados por tipo e o total de Pokémon processados.
