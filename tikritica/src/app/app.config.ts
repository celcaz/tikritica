import { ApplicationConfig, provideBrowserGlobalErrorListeners } from '@angular/core';
import {
  HttpInterceptorFn,
  provideHttpClient,
  withFetch,
  withInterceptors,
} from '@angular/common/http';
import { provideRouter } from '@angular/router';
import { provideEffects } from '@ngrx/effects';
import { provideStore } from '@ngrx/store';

import { routes } from './app.routes';
import { authReducer } from './auth/store/auth.reducer';
import { AuthEffects } from './auth/store/auth.effects';

const withCredentialsInterceptor: HttpInterceptorFn = (req, next) => {
  // Step 1: Always send cookies to the API for cookie-based auth.
  const request = req.clone({ withCredentials: true });
  return next(request);
};

export const appConfig: ApplicationConfig = {
  providers: [
    provideBrowserGlobalErrorListeners(),
    provideRouter(routes),
    provideHttpClient(withFetch(), withInterceptors([withCredentialsInterceptor])),
    provideStore({ auth: authReducer }),
    provideEffects([AuthEffects]),
  ],
};
