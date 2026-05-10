import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { UserProfile } from '../domain/user.domain';

@Injectable({ providedIn: 'root' })
export class ProfileApi {
  private readonly baseUrl = '/api/users';
  private readonly http = inject(HttpClient);

  getUserProfile(username: string): Observable<UserProfile> {
    return this.http.get<UserProfile>(`${this.baseUrl}/${username}`);
  }
}
