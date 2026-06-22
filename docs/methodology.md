# AletheIA Score V1

## Objetivo

O AletheIA é uma plataforma de análise de promessas públicas baseada em evidências verificáveis.

O objetivo do score não é prever o sucesso de uma promessa, mas avaliar sua qualidade, verificabilidade e auditabilidade utilizando critérios objetivos e dados públicos disponíveis.

A metodologia do AletheIA é inspirada em práticas internacionais de avaliação de políticas públicas e análise baseada em evidências.

---

## Como o score funciona

Cada promessa é avaliada em sete critérios.

Cada critério recebe um peso específico e pode assumir um dos seguintes status:

| Status  | Fator |
| ------- | ----: |
| yes     |  100% |
| partial |   50% |
| no      |    0% |

A pontuação final varia de 0 a 100 pontos.

---

## Critérios

| Critério                                      | Peso |
| --------------------------------------------- | ---: |
| Clareza                                       |   10 |
| Mensurabilidade                               |   15 |
| Prazo definido                                |   10 |
| Dados públicos disponíveis                    |   15 |
| Histórico comparável                          |   15 |
| Plausibilidade baseada em evidências públicas |   20 |
| Riscos e dependências                         |   15 |

Total: 100 pontos.

---

## Clareza

Avalia se a promessa possui objetivo claro, compreensível e verificável.

### Exemplos

**Sim**

- Construir 100 hospitais públicos.
- Reduzir a inflação para 4%.

**Não**

- Melhorar a saúde da população.
- Transformar o país.

---

## Mensurabilidade

Avalia se existem indicadores capazes de medir o cumprimento da promessa.

### Exemplos

**Sim**

- Construir 100 hospitais.
- Criar 1 milhão de empregos.

**Parcial**

- Melhorar a qualidade da educação.

---

## Prazo definido

Avalia se existe um horizonte temporal explícito para acompanhamento.

### Exemplos

**Sim**

- Em 2 anos.
- Até 2030.

**Não**

- Sem prazo definido.

---

## Dados públicos disponíveis

Avalia se existem fontes públicas capazes de acompanhar a execução da promessa.

### Exemplos de fontes

- CNES
- DATASUS
- Banco Central
- IBGE
- INEP
- Portal da Transparência
- Tesouro Nacional

---

## Histórico comparável

Avalia se existem dados históricos que permitam comparar a evolução do indicador ao longo do tempo.

Exemplos:

- Número de hospitais.
- Taxa de inflação.
- Número de matrículas escolares.

---

## Plausibilidade baseada em evidências públicas

Avalia a relação entre a meta proposta e a linha de base pública identificada.

### Fórmula atual

Meta proposta ÷ Linha de base pública

### Faixas atuais

| Relação meta / baseline | Status  |
| ----------------------- | ------- |
| Até 10%                 | yes     |
| Acima de 10% até 50%    | partial |
| Acima de 50%            | no      |

### Exemplo

Meta:

100 hospitais

Linha de base:

5115 hospitais

Resultado:

100 ÷ 5115 = 1,96%

Status:

yes

### Observação

As faixas atuais são experimentais e poderão evoluir conforme novas fontes públicas e critérios forem incorporados ao sistema.

---

## Riscos e dependências

Avalia fatores externos que podem impactar a execução da promessa.

Exemplos:

- Dependência orçamentária.
- Aprovação legislativa.
- Licenciamento ambiental.
- Capacidade operacional.
- Contratações públicas.

---

## Confidence

O indicador de confiança representa o nível de completude da análise.

Ele é calculado a partir da quantidade e qualidade dos critérios avaliados.

Quanto maior a disponibilidade de dados, fontes públicas e informações verificáveis, maior tende a ser a confiança da análise.

---

## Limitações

O AletheIA não prevê resultados futuros.

O score não representa probabilidade de sucesso.

O score representa apenas a qualidade da promessa sob a perspectiva de:

- clareza;
- mensurabilidade;
- disponibilidade de dados públicos;
- histórico comparável;
- plausibilidade baseada em evidências;
- riscos identificáveis.

Novas fontes públicas e melhorias metodológicas poderão alterar análises futuras.
