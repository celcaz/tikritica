# Estrutura do Frontend — Tikritica

Stack: Angular 19 (standalone), NgRx + @ngrx/signals, PrimeNG, Tailwind CSS v4, TypeScript

## Visão geral

```
tikritica/
├── src/
│   ├── main.ts                              Bootstrap da aplicação
│   ├── styles.css                           Tailwind + @custom-variant dark
│   └── app/
│       ├── app.ts                           Componente raiz — inicializa ThemeService
│       ├── app.html                         <app-header /> + <router-outlet />
│       ├── app.routes.ts                    Lazy loading por feature
│       ├── app.config.ts                    Providers globais (NgRx, HTTP, PrimeNG)
│       │
│       ├── core/                            Singleton — sem estado de negócio
│       │   ├── guards/
│       │   │   └── auth.guard.ts            CanMatchFn — aguarda status auth antes de avaliar
│       │   ├── interceptors/
│       │   │   └── credentials.interceptor.ts  withCredentials em todas as requests
│       │   └── services/
│       │       └── app-startup.service.ts   Chama AuthFacade.refresh() no bootstrap
│       │
│       ├── shared/                          Reutilizável entre features
│       │   ├── components/
│       │   │   └── header/
│       │   │       ├── header.component.ts
│       │   │       └── header.component.html  Navbar + toggle dark/light
│       │   ├── directives/
│       │   ├── pipes/
│       │   ├── services/
│       │   │   └── theme.service.ts         Signal isDark, toggle, persistência localStorage
│       │   └── utils/
│       │
│       └── features/                        Features lazy-loadable, auto-contidas
│           │
│           ├── auth/
│           │   ├── data/
│           │   │   ├── auth.api.ts          HTTP: login, register, refresh, me, logout
│           │   │   └── auth.repository.ts   Abstração sobre AuthApi
│           │   ├── domain/
│           │   │   └── auth.models.ts       User, AuthState, LoginRequest, RegisterRequest
│           │   ├── state/                   NgRx slice de autenticação
│           │   │   ├── auth.actions.ts
│           │   │   ├── auth.reducer.ts      Estado: user, status ('idle'|'loading'|'authenticated'|'error')
│           │   │   ├── auth.effects.ts      Side effects — usa AuthRepository
│           │   │   ├── auth.selectors.ts
│           │   │   └── auth.facade.ts       API pública via Signals (toSignal)
│           │   ├── ui/
│           │   │   ├── login/
│           │   │   └── register/
│           │   └── auth.routes.ts           /login, /register, /forgot-password
│           │
│           ├── discover/
│           │   ├── ui/
│           │   │   ├── discover.component.ts
│           │   │   └── discover.component.html
│           │   └── discover.routes.ts       /discover (protegida por authGuard)
│           │
│           ├── movie/
│           │   ├── data/
│           │   │   ├── movie.api.ts         HTTP: GET /api/movies, /api/movies/{slug}
│           │   │   └── movie.repository.ts
│           │   ├── domain/
│           │   │   └── movie.model.ts       Movie, MovieState
│           │   ├── state/
│           │   │   └── movie.store.ts       SignalStore — getMovies, getMovieBySlug
│           │   └── store/                   NgRx alternativo (facade, actions, reducer, effects)
│           │
│           └── userProfile/
│               ├── data/
│               │   ├── userProfile.api.ts   HTTP: GET /api/users/{username}
│               │   └── userProfile.repository.ts
│               ├── domain/
│               │   └── user.domain.ts       UserProfile, UserProfileState
│               ├── ui/
│               │   ├── profile-page/        Smart component — injeta store, passa @Input para filhos
│               │   ├── profile-header/      Dumb component — @Input user, @Output follow
│               │   └── profile-stats/       Dumb component — @Input counts
│               └── userProfile.routes.ts    /profile/:username (protegida por authGuard)
```

## Padrão por feature

```
feature/
├── data/       api.ts (HTTP puro) + repository.ts (abstração)
├── domain/     interfaces e types (sem lógica)
├── state/      SignalStore (@ngrx/signals) — withState, withComputed, withMethods, rxMethod
├── ui/         componentes standalone
└── *.routes.ts lazy routes da feature
```

## Smart vs Dumb components

| Tipo             | Responsabilidade                                                                 |
| ---------------- | -------------------------------------------------------------------------------- |
| **Smart** (page) | Injeta store/facade, busca dados, passa via `@Input` para filhos                 |
| **Dumb**         | Recebe dados via `@Input`, emite eventos via `@Output`, sem dependência de store |

## Dark mode

- `styles.css` define `@custom-variant dark (&:where(.my-app-dark, .my-app-dark *))`
- `ThemeService` togla a classe `.my-app-dark` no `<html>` e persiste no `localStorage`
- PrimeNG configurado com `darkModeSelector: '.my-app-dark'`
- Usar prefixo `dark:` nas classes Tailwind normalmente

## Fluxo de autenticação

```
main.ts
  └── AppStartupService.init()
        └── AuthFacade.refresh()
              └── [Auth] Refresh Requested
                    └── AuthEffects → AuthRepository → AuthApi
                          ├── OK  → refreshSucceeded → status: 'authenticated'
                          └── ERR → refreshFailed    → status: 'error'

authGuard: aguarda status !== 'idle' && !== 'loading' antes de avaliar
```

## Convenções

- Arquivos: `kebab-case` (`auth.facade.ts`, `login.component.ts`)
- Classes/Interfaces: `PascalCase` (`AuthFacade`, `LoginComponent`)
- Injeção via `inject()`, sem constructor DI
- Componentes standalone (`standalone: true`) — sem NgModules
- Signals expostos via `toSignal()` no facade ou diretamente no SignalStore
- Formulários inicializados na declaração, não no constructor
