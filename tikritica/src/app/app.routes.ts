import { Routes } from '@angular/router';
import { authGuard } from './core/guards/auth.guard';

export const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: 'login' },
  {
    path: '',
    loadChildren: () => import('./features/auth/auth.routes').then((m) => m.authRoutes),
  },
  {
    path: 'discover',
    loadChildren: () => import('./features/discover/discover.routes').then((m) => m.discoverRoutes),
    canMatch: [authGuard],
  },
];
