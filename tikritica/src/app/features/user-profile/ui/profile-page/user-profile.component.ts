import { Component, effect, inject, input } from '@angular/core';
import { HeaderProfileComponent } from '../profile-header/header.component';
import { StatsProfileComponent } from '../profile-stats/stats.component';
import { ProfileEditComponent } from '../profile-edit/profile-edit';
import { ProfileStore } from '../../state/user-profile.store';
import { UserProfile } from '../../domain/user.domain';

@Component({
  selector: 'app-user-profile',
  imports: [HeaderProfileComponent, StatsProfileComponent, ProfileEditComponent],
  templateUrl: './user-profile.component.html',
  styleUrl: './user-profile.component.css',
})
export class UserProfileComponent {
  readonly store = inject(ProfileStore);
  readonly username = input.required<string>();

  constructor() {
    effect(() => {
      this.store.getUserProfile(this.username());
    });
  }

  updateUserProfile(updatedProfile: UserProfile) {
    console.log('Profile updated with:', updatedProfile);
  }
}
