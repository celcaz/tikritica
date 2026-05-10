import { inject } from '@angular/core';
import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';
import { firstValueFrom } from 'rxjs';

import { ProfileState } from '../domain/user.domain';
import { ProfileRepository } from '../data/userProfile.repository';

const INITIAL_PROFILE_STATE: ProfileState = {
  profile: null,
  status: 'loading',
  error: null,
};

export const ProfileStore = signalStore(
  { providedIn: 'root' },
  withState<ProfileState>(INITIAL_PROFILE_STATE),
  withMethods((store, profileRepository = inject(ProfileRepository)) => ({
    async getUserProfile(username: string): Promise<void> {
      try {
        patchState(store, { status: 'loading', error: null });
        const profile = await firstValueFrom(profileRepository.fetchUserProfile(username));
        patchState(store, { profile, status: 'success', error: null });
      } catch {
        patchState(store, { status: 'error', error: 'Failed to load user profile' });
      }
    },
  })),
);
