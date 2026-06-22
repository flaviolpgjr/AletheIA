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
- Versionar prompts, modelos e critérios de avaliação.

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

A LLM interpreta a promessa e extrai informações estruturadas, mas não calcula o score nem define os pesos dos critérios.

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

- Gemini API
- LLMs
- Processamento de Linguagem Natural
- Extração estruturada de informações
- JSON Mode

### Dados

- IBGE
- IPEA
- Banco Central
- Dados.gov.br
- DATASUS

## Status Atual

Em desenvolvimento (Release 1).

Implementado atualmente:

### Backend

- API REST em Go;
- Health Check;
- Estrutura HTTP com handlers, routes e services;
- Injeção de dependências;
- Repository Pattern;
- Persistência em PostgreSQL;
- Cache de análises por hash da promessa;
- Histórico de análises persistido em banco de dados;
- DTOs de request/response;
- Testes automatizados.

### Inteligência Artificial

- Interface `llm.Client` para desacoplar provedores de IA;
- Integração com Gemini e OpenRouter;
- Prompt estruturado para extração de promessas;
- Resposta da LLM em JSON;
- Parse para `PromiseExtraction`;
- Extração de:
  - resumo;
  - categoria;
  - objetivo;
  - prazo;
  - indicadores;
  - valor alvo (`target_value`);
  - unidade alvo (`target_unit`);
  - critérios;
  - riscos;
  - fontes sugeridas;

- Critérios avaliados pela LLM;
- Explicações dos critérios geradas pela LLM;
- Riscos extraídos pela LLM.

### Score e Metodologia

- `ScoringModelV1`;
- Motor de cálculo de score;
- Confidence calculada a partir dos critérios;
- Critérios ponderados por peso;
- Metodologia documentada em `docs/methodology.md`;
- Critério de plausibilidade baseado em evidências públicas;
- Comparação automática entre meta e linha de base pública;
- Explicação quantitativa da plausibilidade.

### Dados Públicos

- Integração com CNES/DATASUS;
- Coleta automática de baseline público;
- Persistência de baselines em PostgreSQL;
- Evidências públicas incorporadas à análise;
- Contagem nacional de hospitais ativos utilizando dados oficiais do CNES;
- Sistema preparado para múltiplos provedores de dados públicos.

### Frontend

- React + TypeScript;
- Consumo da API de análise;
- Exibição de:
  - score;
  - confidence;
  - resumo;
  - critérios;
  - riscos;
  - fontes;
  - evidências públicas;

- Interface dark/light;
- Componentes de resultado e visualização de score.

---

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
┌─────────────────────┐
│ LLM Provider        │
│ Gemini / OpenRouter │
└─────────────────────┘
↓
PromiseExtraction
↓
Public Data Provider
↓
CNES / DATASUS
↓
Evidence
↓
ScoringModelV1
↓
ScoreCalculatorService
↓
Analysis
↓
PostgreSQL
↓
AnalyzePromiseResponse
```

### Fluxo de Análise

```txt
Promessa
↓
LLM
↓
Extração estruturada
↓
Busca de evidências públicas
↓
Linha de base (baseline)
↓
Avaliação dos critérios
↓
Cálculo de Score
↓
Cálculo de Confidence
↓
Persistência
↓
Resposta para o frontend
```

## Próximos Passos

- Banco Central (IPCA);
- Múltiplos provedores de dados públicos;
- Melhorias na metodologia de plausibilidade;
- Transparência da metodologia no frontend;
- Docker Compose completo;
- Release 1.

## Contrato da Extração LLM

A LLM deve responder em JSON estruturado no formato:

```json
{
  "summary": "string",
  "category": "string",
  "goal": "string",
  "deadline": "string",
  "indicators": ["string"],
  "risks": ["string"],
  "criteria": [
    {
      "key": "string",
      "status": "yes|partial|no",
      "explanation": "string"
    }
  ]
}
```

Critérios oficiais:

- `clarity`
- `measurability`
- `deadline`
- `public_data`
- `historical_baseline`
- `risks_dependencies`

Status permitidos:

- `yes`
- `partial`
- `no`

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
- Persistência das análises;
- Campo `analysis_data` em JSONB preparado para futuro Knowledge Graph;
- Docker;
- Deploy.

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

- Persistir análises no PostgreSQL;
- Criar campo `analysis_data` em JSONB contendo promessa, extração, critérios, riscos, fontes, limitações e metadados;
- Versionar `prompt_version`, `score_model_version`, `llm_provider` e `llm_model`;
- Iniciar integração com uma fonte pública real;
- Adicionar Docker para backend, frontend e PostgreSQL;
- Preparar deploy da Release 1.

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
