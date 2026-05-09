import { inject, Injectable } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { Store } from '@ngrx/store';

import { AuthState, LoginRequest, RegisterRequest } from '../domain/auth.models';
import {
  loginRequested, logoutRequested, meRequested, refreshRequested, registerRequested,
} from './auth.actions';
import {
  selectAuthError, selectAuthStatus, selectAuthUser, selectIsAuthenticated,
} from './auth.selectors';

@Injectable({ providedIn: 'root' })
export class AuthFacade {
  private readonly store = inject(Store);

  readonly user = toSignal(this.store.select(selectAuthUser));
  readonly status = toSignal(this.store.select(selectAuthStatus), { initialValue: 'idle' as AuthState['status'] });
  readonly error = toSignal(this.store.select(selectAuthError), { initialValue: null });
  readonly isAuthenticated = toSignal(this.store.select(selectIsAuthenticated), { initialValue: false });

  login(payload: LoginRequest): void { this.store.dispatch(loginRequested({ payload })); }
  register(payload: RegisterRequest): void { this.store.dispatch(registerRequested({ payload })); }
  refresh(): void { this.store.dispatch(refreshRequested()); }
  me(): void { this.store.dispatch(meRequested()); }
  logout(): void { this.store.dispatch(logoutRequested()); }
}
