import { Routes } from '@angular/router';
import { UserProfileComponent } from './ui/profile-page/user-profile.component';

export const userProfileRoutes: Routes = [{ path: ':username', component: UserProfileComponent }];
