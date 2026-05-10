import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ProfileApi } from './userProfile.api';
import { UserProfile } from '../domain/user.domain';

@Injectable({
  providedIn: 'root',
})
export class ProfileRepository {
  readonly api = inject(ProfileApi);

  fetchUserProfile(username: string): Observable<UserProfile> {
    return this.api.getUserProfile(username);
  }
}
