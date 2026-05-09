import { Routes } from '@angular/router';

import { LoginComponent } from './auth/components/login/login.component';
import { RegisterComponent } from './auth/components/register/register.component';
import { DiscoverComponent } from './discover/discover.component';
import { authGuard } from './auth/guards/auth.guard';

export const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: 'login' },
  { path: 'login', component: LoginComponent },
  { path: 'register', component: RegisterComponent },
  { path: 'forgot-password', component: LoginComponent },
  {
    path: 'discover',
    component: DiscoverComponent,
    canMatch: [authGuard],
  },
];
