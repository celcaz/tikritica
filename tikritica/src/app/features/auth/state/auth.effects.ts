import { inject, Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { catchError, map, of, switchMap, tap } from 'rxjs';

import { AuthRepository } from '../data/auth.repository';
import {
  loginFailed, loginRequested, loginSucceeded,
  logoutFailed, logoutRequested, logoutSucceeded,
  meFailed, meRequested, meSucceeded,
  refreshFailed, refreshRequested, refreshSucceeded,
  registerFailed, registerRequested, registerSucceeded,
} from './auth.actions';

@Injectable()
export class AuthEffects {
  private readonly actions$ = inject(Actions);
  private readonly authRepository = inject(AuthRepository);
  private readonly router = inject(Router);

  login$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loginRequested),
      switchMap(({ payload }) =>
        this.authRepository.login(payload).pipe(
          map((response) => loginSucceeded({ response })),
          catchError((error) => of(loginFailed({ error: this.toMessage(error) }))),
        ),
      ),
    ),
  );

  register$ = createEffect(() =>
    this.actions$.pipe(
      ofType(registerRequested),
      switchMap(({ payload }) =>
        this.authRepository.register(payload).pipe(
          map((response) => registerSucceeded({ response })),
          catchError((error) => of(registerFailed({ error: this.toMessage(error) }))),
        ),
      ),
    ),
  );

  refresh$ = createEffect(() =>
    this.actions$.pipe(
      ofType(refreshRequested),
      switchMap(() =>
        this.authRepository.refresh().pipe(
          map((response) => refreshSucceeded({ response })),
          catchError((error) => of(refreshFailed({ error: this.toMessage(error) }))),
        ),
      ),
    ),
  );

  me$ = createEffect(() =>
    this.actions$.pipe(
      ofType(meRequested),
      switchMap(() =>
        this.authRepository.me().pipe(
          map((response) => meSucceeded({ response })),
          catchError((error) => of(meFailed({ error: this.toMessage(error) }))),
        ),
      ),
    ),
  );

  logout$ = createEffect(() =>
    this.actions$.pipe(
      ofType(logoutRequested),
      switchMap(() =>
        this.authRepository.logout().pipe(
          map(() => logoutSucceeded()),
          catchError((error) => of(logoutFailed({ error: this.toMessage(error) }))),
        ),
      ),
    ),
  );

  loginRedirect$ = createEffect(
    () => this.actions$.pipe(
      ofType(loginSucceeded, registerSucceeded),
      tap(() => this.router.navigate(['/discover'])),
    ),
    { dispatch: false },
  );

  logoutRedirect$ = createEffect(
    () => this.actions$.pipe(
      ofType(logoutSucceeded),
      tap(() => this.router.navigate(['/login'])),
    ),
    { dispatch: false },
  );

  private toMessage(error: unknown): string {
    if (error && typeof error === 'object' && 'error' in error) {
      const httpError = error as { error?: { error?: string } };
      return httpError.error?.error ?? 'unexpected error';
    }
    return 'unexpected error';
  }
}
