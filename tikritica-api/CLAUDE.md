# Tikritica API — Backend

API REST para a plataforma Tikritica de tracking e review de mídias.

## Stack

- **Linguagem:** Go 1.23
- **Router:** net/http (stdlib, Go 1.22+ pattern matching)
- **Banco:** PostgreSQL (driver lib/pq)
- **Arquitetura:** Clean Architecture (Uncle Bob)
- **Build:** make run / make build

## Estrutura (Clean Architecture)

```
├── cmd/api/main.go                          Entrypoint + Composition Root
├── internal/
│   ├── domain/                              CAMADA 1 — Entidades + Portas (mais interna)
│   │   ├── entity/                          Structs de domínio puras
│   │   │   ├── movie.go, series.go, game.go, book.go
│   │   │   ├── user.go, review.go, list.go, diary.go, social.go
│   │   └── port/                            Interfaces de repositório (contratos)
│   │       ├── movie.go, series.go, game.go, book.go
│   │       ├── user.go, review.go, list.go, diary.go, social.go
│   ├── usecase/                             CAMADA 2 — Casos de uso por feature
│   │   ├── movie/usecase.go
│   │   ├── series/usecase.go
│   │   ├── game/usecase.go
│   │   └── book/usecase.go
│   ├── adapter/                             CAMADA 3 — Interface Adapters
│   │   ├── handler/                         HTTP handlers (controllers)
│   │   │   ├── handler.go, routes.go, health.go
│   │   │   ├── movie.go, series.go, game.go, book.go
│   │   │   └── review.go, user.go, list.go, diary.go, social.go
│   │   └── repository/postgres/             Implementações PostgreSQL dos ports
│   │       ├── movie.go, series.go, game.go, book.go
│   └── infra/                               CAMADA 4 — Frameworks & Drivers
│       ├── config/config.go
│       ├── database/postgres.go
│       └── middleware/cors.go, logging.go
├── pkg/response/                            Helpers de resposta JSON
├── migrations/                              Migrações SQL
├── Makefile
└── go.mod
```

## Regra de Dependência

```
domain/entity   ← não importa nada do projeto
domain/port     ← importa apenas entity
usecase/*       ← importa port + entity
adapter/handler ← importa usecase/* + entity
adapter/repository/postgres ← importa port + entity
infra/*         ← sem dependências internas
cmd/api/main.go ← importa TUDO (composição)
```

## Convenções

- Fluxo: `handler → usecase → repository (via interface port)`
- `internal/` não é importável por projetos externos (convenção Go)
- `pkg/` contém código reutilizável e exportável
- Entidades em `domain/entity/` são structs puras sem dependências externas
- Interfaces em `domain/port/` definem contratos dos repositórios
- Use cases recebem interfaces (ports), não implementações concretas
- Handlers chamam use cases, nunca repositórios diretamente
- Campos sensíveis (email, password) usam `json:"-"`
- Endpoints seguem padrão REST: `GET /api/movies`, `GET /api/movies/{slug}`

## Endpoints

| Método | Rota                    | Descrição              | Status         |
|--------|-------------------------|------------------------|----------------|
| GET    | /health                 | Health check           | Implementado   |
| GET    | /api/movies             | Listar filmes          | Implementado   |
| GET    | /api/movies/{slug}      | Detalhe do filme       | Implementado   |
| GET    | /api/series             | Listar séries          | Implementado   |
| GET    | /api/series/{slug}      | Detalhe da série       | Implementado   |
| GET    | /api/games              | Listar jogos           | Implementado   |
| GET    | /api/games/{slug}       | Detalhe do jogo        | Implementado   |
| GET    | /api/books              | Listar livros          | Implementado   |
| GET    | /api/books/{slug}       | Detalhe do livro       | Implementado   |
| GET    | /api/reviews/{id}       | Detalhe da review      | 501            |
| POST   | /api/reviews            | Criar review           | 501            |
| GET    | /api/users/{username}   | Perfil do usuário      | 501            |
| GET    | /api/lists/{id}         | Detalhe da lista       | 501            |
| POST   | /api/follow/{userId}    | Seguir usuário         | 501            |
| DELETE | /api/follow/{userId}    | Deixar de seguir       | 501            |
| GET    | /api/feed               | Feed de atividades     | 501            |
| GET    | /api/diary              | Diário do usuário      | 501            |
| POST   | /api/diary              | Criar entrada no diário| 501            |

## Frontend

Repositório separado em Next.js: `~/projects/tikritica`
