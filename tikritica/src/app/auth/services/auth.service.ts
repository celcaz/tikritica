import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { AuthUserResponse, LoginRequest, RegisterRequest } from '../models/auth.models';

@Injectable({ providedIn: 'root' })
export class AuthService {
  private readonly baseUrl = '/api/auth';

  constructor(private readonly http: HttpClient) {}

  // Step 1: Send login credentials; backend sets HttpOnly cookie.
  login(payload: LoginRequest): Observable<AuthUserResponse> {
    return this.http.post<AuthUserResponse>(`${this.baseUrl}/login`, payload);
  }

  // Step 2: Register user; backend sets HttpOnly cookie.
  register(payload: RegisterRequest): Observable<AuthUserResponse> {
    return this.http.post<AuthUserResponse>(`${this.baseUrl}/register`, payload);
  }

  // Step 3: Refresh session from cookie.
  refresh(): Observable<AuthUserResponse> {
    return this.http.post<AuthUserResponse>(`${this.baseUrl}/refresh`, {});
  }

  // Step 4: Resolve current user from cookie.
  me(): Observable<AuthUserResponse> {
    return this.http.get<AuthUserResponse>(`${this.baseUrl}/me`);
  }

  // Step 5: Clear auth cookie.
  logout(): Observable<{ message: string }> {
    return this.http.post<{ message: string }>(`${this.baseUrl}/logout`, {});
  }
}
