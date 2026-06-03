# AletheIA

AletheIA é uma plataforma que utiliza Inteligência Artificial e dados públicos para analisar promessas públicas de forma explicável, transparente e auditável.

O objetivo do projeto não é prever o futuro nem determinar verdades absolutas. O foco é fornecer uma análise estruturada baseada em critérios explícitos, evidências verificáveis e fontes públicas confiáveis.

O nome "AletheIA" é inspirado em _Aletheia_, conceito da filosofia grega associado à verdade e revelação, combinado com "IA" (Inteligência Artificial).

## Princípios

O AletheIA é construído sobre alguns princípios fundamentais:

- Não inventar dados.
- Exibir as fontes utilizadas.
- Explicar como a conclusão foi alcançada.
- Tornar os critérios visíveis ao usuário.
- Expor limitações da análise.
- Informar o nível de confiança da análise.
- Evitar modelos de decisão em caixa-preta.
- Separar interpretação por IA de cálculo de score.

## Objetivo

Transformar promessas públicas escritas em linguagem natural em análises estruturadas e explicáveis.

Fluxo pretendido:

```txt
Promessa pública
↓
LLM
↓
Extração de informações estruturadas
↓
Consulta a dados públicos
↓
Avaliação por critérios
↓
Score
↓
Confidence
↓
Explicação da análise
↓
Persistência e histórico
```

## O que o AletheIA analisa

A análise considera fatores como:

- Clareza da promessa;
- Mensurabilidade;
- Existência de prazo;
- Disponibilidade de dados públicos;
- Histórico comparável;
- Dependências e riscos;
- Evidências encontradas em fontes públicas.

## Score AletheIA

O score não representa verdade absoluta, intenção política ou probabilidade exata de cumprimento.

O Score AletheIA representa o grau de sustentação verificável de uma promessa pública com base em critérios objetivos, dados disponíveis e evidências auditáveis.

Critérios iniciais do modelo v1:

- Clareza da promessa;
- Mensurabilidade;
- Prazo definido;
- Dados públicos disponíveis;
- Histórico comparável;
- Dependências e riscos.

A LLM pode interpretar a promessa e extrair informações estruturadas, mas não calcula o score nem define os pesos dos critérios.

## Stack

### Backend

- Go
- Chi Router
- PostgreSQL
- Docker

### Frontend

- React
- TypeScript
- TailwindCSS
- React Router

### Inteligência Artificial

- LLMs
- Processamento de Linguagem Natural
- Extração estruturada de informações

### Dados

- IBGE
- IPEA
- Banco Central
- Dados.gov.br
- DATASUS

## Status Atual

Em desenvolvimento.

Implementado atualmente:

- API REST em Go;
- Health Check;
- Estrutura HTTP com handlers, routes e services;
- Handler de análise estruturado com injeção de dependência;
- Modelo inicial de score;
- Critérios de avaliação;
- ScoringModelV1;
- Motor de cálculo de score;
- Confidence inicial;
- DTOs de request/response;
- Frontend consumindo score, confidence, critérios e riscos;
- Interface inicial para integração com LLM;
- GeminiClient estruturado;
- API Key configurada via variável de ambiente;
- Testes automatizados.

## Arquitetura Atual

```txt
Frontend React
↓
POST /promises/analyze
↓
AnalyzePromiseHandler
↓
PromiseAnalyzerService
↓
ScoreCalculatorService
↓
Analysis
↓
AnalyzePromiseResponse
```

Camada LLM em preparação:

```txt
llm.Client
↓
GeminiClient
↓
PromiseExtraction
```

## Roadmap

### Release 1 — Core

Primeira versão pública do produto.

Inclui:

- LLM;
- Dados públicos;
- Score;
- Confidence;
- Critérios explicáveis;
- Fontes utilizadas;
- Limitações da análise;
- PostgreSQL;
- Histórico de análises.

### Release 2 — Explainability

- Grafo lógico da análise;
- Visualização do raciocínio utilizado;
- Explicação detalhada dos critérios.

### Release 3 — Multi-Fonte

- Integração com múltiplas bases públicas;
- Cruzamento de indicadores;
- Maior robustez das análises.

### Release 4 — Knowledge Graph

- Grafo como modelo de dados;
- Relacionamento entre promessas, indicadores, fontes e evidências.

### Release 5 — Inteligência Acumulada

- Reutilização de análises anteriores;
- Comparação entre promessas semelhantes;
- Aprendizado sobre padrões históricos.

### Release 6 — Plataforma

- Comparações;
- Dashboards;
- Busca;
- Monitoramento contínuo.

## Próximos Passos

- Implementar a primeira chamada HTTP real para a API do Gemini;
- Utilizar o prompt de extração estruturada;
- Converter a resposta da LLM em `PromiseExtraction`;
- Substituir gradualmente heurísticas locais por interpretação da LLM;
- Iniciar integração com dados públicos;
- Persistir análises no PostgreSQL.

## Objetivo Técnico

Além do impacto social, o projeto também serve como estudo prático de:

- Go;
- Arquitetura de Software;
- APIs REST;
- Integração com IA;
- Engenharia de Dados;
- Sistemas Explicáveis;
- Testes Automatizados;
- Processamento de Dados Públicos;
- Boas Práticas de Engenharia de Software.

## Motivação

O AletheIA surgiu da ideia de utilizar tecnologia para tornar análises públicas mais transparentes e acessíveis.

Em vez de produzir respostas opacas, o projeto busca apresentar critérios, evidências, fontes e limitações de forma clara, permitindo que qualquer pessoa compreenda como uma conclusão foi construída.
