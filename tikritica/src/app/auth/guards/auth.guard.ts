import { inject } from '@angular/core';
import { CanMatchFn, Router, UrlTree } from '@angular/router';
import { Store } from '@ngrx/store';
import { map } from 'rxjs';

import { selectIsAuthenticated } from '../store/auth.selectors';

export const authGuard: CanMatchFn = () => {
  const store = inject(Store);
  const router = inject(Router);

  // Step 1: Read auth state from NgRx.
  // Step 2: If not authenticated, redirect to login with returnUrl.
  return store.select(selectIsAuthenticated).pipe(
    map((isAuthenticated): boolean | UrlTree => {
      if (isAuthenticated) {
        return true;
      }

      return router.createUrlTree(['/login'], {
        queryParams: { returnUrl: '/discover' },
      });
    }),
  );
};
