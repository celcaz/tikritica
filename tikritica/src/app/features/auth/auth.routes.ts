import { Routes } from '@angular/router';

import { LoginComponent } from './ui/login/login.component';
import { RegisterComponent } from './ui/register/register.component';

export const authRoutes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: 'register', component: RegisterComponent },
  { path: 'forgot-password', component: LoginComponent },
];
