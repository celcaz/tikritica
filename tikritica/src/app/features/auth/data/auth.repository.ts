import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { AuthUserResponse, LoginRequest, RegisterRequest } from '../domain/auth.models';
import { AuthApi } from './auth.api';

@Injectable({ providedIn: 'root' })
export class AuthRepository {
  private readonly api = inject(AuthApi);

  login(payload: LoginRequest): Observable<AuthUserResponse> {
    return this.api.login(payload);
  }

  register(payload: RegisterRequest): Observable<AuthUserResponse> {
    return this.api.register(payload);
  }

  refresh(): Observable<AuthUserResponse> {
    return this.api.refresh();
  }

  me(): Observable<AuthUserResponse> {
    return this.api.me();
  }

  logout(): Observable<{ message: string }> {
    return this.api.logout();
  }
}
