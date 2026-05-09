import { Injectable } from '@angular/core';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs';

import { AuthState, LoginRequest, RegisterRequest } from '../models/auth.models';
import {
  loginRequested,
  logoutRequested,
  meRequested,
  refreshRequested,
  registerRequested,
} from '../store/auth.actions';
import {
  selectAuthError,
  selectAuthStatus,
  selectAuthUser,
  selectIsAuthenticated,
} from '../store/auth.selectors';

@Injectable({ providedIn: 'root' })
export class AuthFacade {
  readonly user$: Observable<AuthState['user']>;
  readonly status$: Observable<AuthState['status']>;
  readonly error$: Observable<AuthState['error']>;
  readonly isAuthenticated$: Observable<boolean>;

  constructor(private readonly store: Store) {
    this.user$ = this.store.select(selectAuthUser);
    this.status$ = this.store.select(selectAuthStatus);
    this.error$ = this.store.select(selectAuthError);
    this.isAuthenticated$ = this.store.select(selectIsAuthenticated);
  }

  // Step 1: Dispatch login request.
  login(payload: LoginRequest): void {
    this.store.dispatch(loginRequested({ payload }));
  }

  // Step 2: Dispatch register request.
  register(payload: RegisterRequest): void {
    this.store.dispatch(registerRequested({ payload }));
  }

  // Step 3: Attempt to restore session.
  refresh(): void {
    this.store.dispatch(refreshRequested());
  }

  // Step 4: Fetch current user from cookie.
  me(): void {
    this.store.dispatch(meRequested());
  }

  // Step 5: Logout and clear auth state.
  logout(): void {
    this.store.dispatch(logoutRequested());
  }
}
