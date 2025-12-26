# GoVue CRM

> Um sistema de gestão de relacionamento com o cliente (CRM) de alta performance, construído com arquitetura de microsserviços e front-end moderno.

![Project Status](https://img.shields.io/badge/status-development-orange)
![Go Version](https://img.shields.io/badge/go-1.21+-blue)
![Nuxt Version](https://img.shields.io/badge/nuxt-4-green)

## 🎯 Sobre o Projeto

O **GoVue CRM** é uma solução *fullstack* projetada para escalabilidade e manutenibilidade. O projeto adota uma abordagem de **Monorepo**, unificando o Back-end (focado em performance bruta e concorrência) e o Front-end (focado em experiência de usuário e reatividade) sob um mesmo fluxo de versionamento.

A principal filosofia arquitetural deste projeto é a **Clean Architecture** (Arquitetura Limpa) e a **Separação de Preocupações (SoC)**, aplicadas tanto no servidor quanto no cliente.

## 🛠 Tech Stack

### Back-end (API REST)
* **Linguagem:** Go (Golang)
* **Framework:** Gin Web Framework (pela baixa latência)
* **Arquitetura:** Layered Architecture (Handler -> Service -> Repository)

### Front-end (SPA/SSR)
* **Framework:** Nuxt 4 (Vue.js Ecosystem)
* **UI Library:** Shadcn-vue + Tailwind CSS
* **State Management:** Composables (Native Reactivity) & VueUse
* **Arquitetura:** Modular com foco em Composables como camada de serviço.


## 📂 Estrutura do Projeto (Monorepo)

O projeto está organizado para facilitar o desenvolvimento isolado de cada parte da stack, mantendo a coesão do produto.

```text
.
├── backend/                # API desenvolvida em Gin
├── frontend/               # Aplicação Web Nuxt 4
├── compose.yml             # Orquestração de containers (App + DBs)
└── README.md               # Documentação do projeto
```


## 🏗 Arquitetura do Front-end (`/frontend`)

O front-end foi construído seguindo as diretrizes do **Nuxt 4**, com o código-fonte isolado no diretório `app/`, mantendo a raiz do projeto limpa e organizada.

A arquitetura espelha os padrões de camadas do backend, garantindo consistência mental, previsibilidade e facilidade de manutenção.


## 📁 Estrutura Arquitetural — Tabela de Diretórios

| Diretório        | Camada Arquitetural        | Descrição Técnica |
|------------------|----------------------------|-------------------|
| `pages/`         | Controllers / View         | Responsável pelo roteamento via *File System Routing* do Nuxt. Atua como ponto de entrada da requisição do usuário, coordenando a navegação e disparando chamadas para os serviços. Não contém regras de negócio complexas. |
| `composables/`   | Service Layer              | Camada central da lógica de negócio do front-end. Encapsula estado, regras de negócio e chamadas à API (Repository Pattern via `useFetch`). Mantém a UI desacoplada e reutilizável. |
| `components/`    | UI / Presentation          | Componentes visuais isolados e reutilizáveis. Inclui componentes atômicos (`ui/`, baseados em Shadcn UI) e componentes compostos orientados ao domínio da aplicação. |
| `layouts/`       | Scaffolding                | Define as estruturas macro persistentes da aplicação, como Header, Sidebar, grids principais e containers globais. |
| `middleware/`    | Interceptors / Guards      | Guardas de rota executados antes da renderização da página. Responsável por autenticação (JWT), autorização e controle de permissões (ACL). |
| `types/`         | Domain Definitions         | Definições de interfaces e tipos TypeScript. Atua como as *Structs* do front-end, garantindo tipagem forte e consistência nos dados trafegados. |
| `lib/`           | Helpers                    | Funções puras e utilitários de uso geral, como formatação de moeda, datas, validações e regras reutilizáveis. |
| `utils/`         | Helpers                    | Utilitários auxiliares específicos da aplicação ou do Tailwind, sem estado e com foco em reutilização. |


## 🏗 Arquitetura do Back-end (`/backend`)

> **Nota:** Estrutura planejada seguindo os princípios de **Clean Code em Go** e **Clean Architecture**.

| Diretório               | Camada            | Descrição Técnica |
|-------------------------|-------------------|-------------------|
| `cmd/`                  | Main              | Ponto de entrada da aplicação (`main.go`). Responsável por inicializar configurações, rotas, middlewares e injeção de dependências. |
| `internal/`             | Private           | Código que não pode ser importado por outros projetos. Contém a lógica central (core) da aplicação. |
| `internal/handlers/`    | Transport         | Camada HTTP baseada em Gin. Recebe a requisição, valida o JSON (binding), trata erros e chama a camada de serviço. Equivalente às `pages/` do front-end. |
| `internal/services/`    | Business Logic    | Contém as regras de negócio puras. Não possui conhecimento sobre HTTP, banco de dados ou frameworks. Equivalente aos `composables/` do front-end. |
| `internal/repository/`  | Data Access       | Camada de acesso a dados. Responsável pela comunicação direta com o banco de dados e execução das queries SQL. |
| `pkg/`                 | Public            | Bibliotecas auxiliares reutilizáveis, como loggers, parsers e helpers compartilháveis entre projetos. |


## ⚙️ Gestão de Configurações & Logs

O projeto segue as premissas de [12-Factor Apps](https://12factor.net/) para configurações e observabilidade:

- **Configuração Dinâmica:** Utiliza `Viper` com mapeamento automático via Reflection.
- **Segurança:** Validação de variáveis obrigatórias no boot da aplicação (Fail Fast).
- **Logs Estruturados:** Logs emitidos em formato `JSON` via pacote nativo `slog` para facilitar a integração com stacks de monitoramento (ELK/Loki).


## 🚀 Como Rodar o Projeto

### Pré-requisitos
* **Docker & Docker Compose**

* **Node.js 20+** (para desenvolvimento local do front)

* **Go 1.21+** (para desenvolvimento local do back)

## Rodando com Docker (Recomendado)
### Na raiz do projeto:

1. **Configure o ambiente:**
   Copie o arquivo de exemplo e ajuste as variáveis se necessário:
   ```bash
   cp .env.example .env
2. **Suba os serviços:**
   ```bash
   docker compose up -d
> O Front-end estará disponível em **http://localhost:3000** e a API em **http://localhost:8080**.

### Desenvolvimento Local (Front-end)

```bash
cd frontend
pnpm install
pnpm run dev
```

## 📝 Padrões de Código
* **Commits:** Segue o padrão [Conventional Commits](https://www.conventionalcommits.org/) (ex: feat: add user login, fix: button color).

* **Linting:** ESLint + Prettier (Front-end) e golangci-lint (Back-end).

* **CSS:** Tailwind CSS com a metodologia [utility-first](https://v2.tailwindcss.com/docs/utility-first).

Feito por [Paulo Henrique](https://www.linkedin.com/in/paulo-app)


***