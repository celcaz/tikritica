import { Routes } from '@angular/router';

import { LoginComponent } from './features/auth/components/login/login.component';
import { RegisterComponent } from './features/auth/components/register/register.component';
import { DiscoverComponent } from './features/discover/components/discover.component';
import { authGuard } from './core/guards/auth.guard';

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
