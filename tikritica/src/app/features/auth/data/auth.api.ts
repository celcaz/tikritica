import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { AuthUserResponse, LoginRequest, RegisterRequest } from '../domain/auth.models';

@Injectable({ providedIn: 'root' })
export class AuthApi {
  private readonly baseUrl = '/api/auth';
  private readonly http = inject(HttpClient);

  login(payload: LoginRequest): Observable<AuthUserResponse> {
    return this.http.post<AuthUserResponse>(`${this.baseUrl}/login`, payload);
  }

  register(payload: RegisterRequest): Observable<AuthUserResponse> {
    return this.http.post<AuthUserResponse>(`${this.baseUrl}/register`, payload);
  }

  refresh(): Observable<AuthUserResponse> {
    return this.http.post<AuthUserResponse>(`${this.baseUrl}/refresh`, {});
  }

  me(): Observable<AuthUserResponse> {
    return this.http.get<AuthUserResponse>(`${this.baseUrl}/me`);
  }

  logout(): Observable<{ message: string }> {
    return this.http.post<{ message: string }>(`${this.baseUrl}/logout`, {});
  }
}
