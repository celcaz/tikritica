# Setup — Tikritica API

Guia passo a passo para subir o backend localmente.

## Requisitos

- [Go 1.23+](https://go.dev/dl/)
- [Docker](https://docs.docker.com/get-docker/)

## 1. Subir o banco de dados

```bash
docker compose up -d
```

Isso cria um container `tikritica-db` com PostgreSQL 17 (Alpine) na porta **5432**.

Credenciais padrão (configuráveis via variáveis de ambiente no `docker-compose.yml`):

| Variável          | Valor padrão |
|-------------------|--------------|
| `POSTGRES_USER`   | tikritica    |
| `POSTGRES_PASSWORD` | tikritica  |
| `POSTGRES_DB`     | tikritica    |

Para verificar se o banco está rodando:

```bash
docker ps
```

## 2. Aplicar as migrations

```bash
psql "postgres://tikritica:tikritica@localhost:5432/tikritica?sslmode=disable" \
  -f migrations/001_create_users.sql \
  -f migrations/002_create_movies.sql \
  -f migrations/003_create_series.sql \
  -f migrations/004_create_games.sql \
  -f migrations/005_create_books.sql \
  -f migrations/006_create_reviews.sql \
  -f migrations/007_create_lists.sql \
  -f migrations/008_create_diary.sql \
  -f migrations/009_create_social.sql
```

Ou, se preferir rodar todas de uma vez:

```bash
for f in migrations/*.sql; do
  psql "postgres://tikritica:tikritica@localhost:5432/tikritica?sslmode=disable" -f "$f"
done
```

## 3. Rodar a API

```bash
DATABASE_URL="postgres://tikritica:tikritica@localhost:5432/tikritica?sslmode=disable" make run
```

O servidor vai subir em `http://localhost:8080`.

## 4. Testar

```bash
curl http://localhost:8080/health
# {"status":"ok"}
```

## Comandos úteis

| Comando                    | Descricao                      |
|----------------------------|--------------------------------|
| `docker compose up -d`     | Subir o banco                  |
| `docker compose down`      | Parar o banco                  |
| `docker compose down -v`   | Parar e apagar os dados        |
| `docker logs tikritica-db` | Ver logs do PostgreSQL         |
| `make run`                 | Rodar a API (precisa do DATABASE_URL) |
| `make build`               | Compilar o binario em `bin/api` |
| `make test`                | Rodar os testes                |
