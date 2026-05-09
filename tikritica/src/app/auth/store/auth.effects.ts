import { Injectable, inject } from '@angular/core';
import { Router } from '@angular/router';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { catchError, map, of, switchMap, tap } from 'rxjs';

import { AuthService } from '../services/auth.service';
import {
  loginFailed,
  loginRequested,
  loginSucceeded,
  logoutFailed,
  logoutRequested,
  logoutSucceeded,
  meFailed,
  meRequested,
  meSucceeded,
  refreshFailed,
  refreshRequested,
  refreshSucceeded,
  registerFailed,
  registerRequested,
  registerSucceeded,
} from './auth.actions';

@Injectable()
export class AuthEffects {
  private readonly actions$ = inject(Actions);
  private readonly authService = inject(AuthService);
  private readonly router = inject(Router);

  // Step 1: Handle login request.
  login$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loginRequested),
      switchMap(({ payload }) =>
        this.authService.login(payload).pipe(
          map((response) => loginSucceeded({ response })),
          catchError((error) => of(loginFailed({ error: this.toMessage(error) }))),
        ),
      ),
    ),
  );

  // Step 2: Handle registration request.
  register$ = createEffect(() =>
    this.actions$.pipe(
      ofType(registerRequested),
      switchMap(({ payload }) =>
        this.authService.register(payload).pipe(
          map((response) => registerSucceeded({ response })),
          catchError((error) => of(registerFailed({ error: this.toMessage(error) }))),
        ),
      ),
    ),
  );

  // Step 3: Refresh session on app start.
  refresh$ = createEffect(() =>
    this.actions$.pipe(
      ofType(refreshRequested),
      switchMap(() =>
        this.authService.refresh().pipe(
          map((response) => refreshSucceeded({ response })),
          catchError((error) => of(refreshFailed({ error: this.toMessage(error) }))),
        ),
      ),
    ),
  );

  // Step 4: Resolve current user from cookie.
  me$ = createEffect(() =>
    this.actions$.pipe(
      ofType(meRequested),
      switchMap(() =>
        this.authService.me().pipe(
          map((response) => meSucceeded({ response })),
          catchError((error) => of(meFailed({ error: this.toMessage(error) }))),
        ),
      ),
    ),
  );

  // Step 5: Logout and clear state.
  logout$ = createEffect(() =>
    this.actions$.pipe(
      ofType(logoutRequested),
      switchMap(() =>
        this.authService.logout().pipe(
          map(() => logoutSucceeded()),
          catchError((error) => of(logoutFailed({ error: this.toMessage(error) }))),
        ),
      ),
    ),
  );

  // Step 6: Navigate to discover after login/register.
  loginRedirect$ = createEffect(
    () =>
      this.actions$.pipe(
        ofType(loginSucceeded, registerSucceeded),
        tap(() => {
          this.router.navigate(['/discover']);
        }),
      ),
    { dispatch: false },
  );

  // Step 7: Navigate back to login if session refresh fails.
  refreshFailedRedirect$ = createEffect(
    () =>
      this.actions$.pipe(
        ofType(refreshFailed, meFailed),
        tap(() => {
          this.router.navigate(['/login']);
        }),
      ),
    { dispatch: false },
  );

  // Step 8: Navigate to login after logout.
  logoutRedirect$ = createEffect(
    () =>
      this.actions$.pipe(
        ofType(logoutSucceeded),
        tap(() => {
          this.router.navigate(['/login']);
        }),
      ),
    { dispatch: false },
  );

  private toMessage(error: unknown): string {
    if (typeof error === 'string') {
      return error;
    }

    if (error && typeof error === 'object' && 'error' in error) {
      const httpError = error as { error?: { error?: string } };
      return httpError.error?.error ?? 'unexpected error';
    }

    return 'unexpected error';
  }
}
