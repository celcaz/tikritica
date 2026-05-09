import { inject } from '@angular/core';
import { CanMatchFn, Router, UrlTree } from '@angular/router';
import { Store } from '@ngrx/store';
import { filter, map, take } from 'rxjs';

import { selectAuthState } from '../../features/auth/state/auth.selectors';

export const authGuard: CanMatchFn = () => {
  const store = inject(Store);
  const router = inject(Router);

  return store.select(selectAuthState).pipe(
    filter((state) => state.status !== 'idle' && state.status !== 'loading'),
    take(1),
    map((state): boolean | UrlTree => {
      if (state.status === 'authenticated' && !!state.user) return true;
      return router.createUrlTree(['/login'], { queryParams: { returnUrl: '/discover' } });
    }),
  );
};
