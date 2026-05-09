# Estrutura do Frontend — Tikritica

Stack: Angular 19 (standalone), NgRx, TypeScript

## Visão geral

```
tikritica/
├── src/
│   ├── main.ts                         # Bootstrap da aplicação
│   ├── styles.css                      # Estilos globais
│   └── app/
│       ├── app.ts                      # Componente raiz
│       ├── app.html                    # Template raiz
│       ├── app.css                     # Estilos raiz
│       ├── app.routes.ts               # Definição de rotas
│       ├── app.config.ts               # Providers globais (NgRx, HTTP, Router)
│       │
│       ├── core/                       # Singleton — carregado uma vez, sem estado de negócio
│       │   ├── guards/
│       │   │   └── auth.guard.ts       # CanMatchFn — redireciona para /login se não autenticado
│       │   ├── models/
│       │   │   └── auth.models.ts      # Interfaces: User, AuthState, LoginRequest, etc.
│       │   └── services/
│       │       ├── auth.service.ts     # HTTP: login, register, refresh, me, logout
│       │       └── app-startup.service.ts  # Restaura sessão no bootstrap
│       │
│       └── features/                   # Features lazy-loadable, auto-contidas
│           ├── auth/
│           │   ├── components/
│           │   │   ├── login/
│           │   │   │   ├── login.component.ts
│           │   │   │   └── login.component.html
│           │   │   └── register/
│           │   │       ├── register.component.ts
│           │   │       └── register.component.html
│           │   └── store/              # NgRx slice de autenticação
│           │       ├── auth.actions.ts     # 13 actions (login, register, refresh, me, logout)
│           │       ├── auth.reducer.ts     # Estado: user, status, error
│           │       ├── auth.effects.ts     # Side effects + navegação pós-auth
│           │       ├── auth.selectors.ts   # Seletores: user$, status$, error$, isAuthenticated$
│           │       └── auth.facade.ts      # API pública do store para os componentes
│           │
│           └── discover/
│               └── components/
│                   ├── discover.component.ts
│                   └── discover.component.html
│
├── docs/
│   └── estrutura.md                    # Este arquivo
├── angular.json
├── package.json
├── tsconfig.json
└── proxy.conf.json                     # Proxy /api → backend Go em dev
```

## Camadas e responsabilidades

| Camada | Regra |
|---|---|
| `core/` | Serviços singleton (`providedIn: 'root'`). Nunca importados dentro de features. |
| `core/models/` | Interfaces e types globais. Sem lógica. |
| `core/services/` | Chamadas HTTP e inicialização. Sem estado local. |
| `core/guards/` | Guards de rota. Leem o NgRx store via seletores. |
| `features/` | Feature auto-contida. Só conhece `core/` e `shared/` (quando existir). |
| `features/*/store/` | NgRx: actions, reducer, effects, selectors e facade por feature. |
| `features/*/components/` | Componentes standalone. Interagem com o store via facade. |

## Fluxo de autenticação

```
main.ts
  └── AppStartupService.init()
        └── AuthFacade.refresh()
              └── [Auth] Refresh Requested
                    └── AuthEffects.refresh$
                          ├── OK  → refreshSucceeded → status: 'authenticated'
                          └── ERR → refreshFailed    → navigate('/login')
```

## Convenções

- Arquivos: `kebab-case` (`auth.facade.ts`, `login.component.ts`)
- Classes/Interfaces: `PascalCase` (`AuthFacade`, `LoginComponent`)
- Injeção via `inject()` (Angular 14+), sem constructor quando possível
- Componentes standalone (`standalone: true`) — sem NgModules
- Propriedades do formulário inicializadas na declaração, não no constructor
