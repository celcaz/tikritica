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
import { authReducer } from './features/auth/store/auth.reducer';
import { AuthEffects } from './features/auth/store/auth.effects';
import { MovieEffects } from './features/movie/store/movie.effects';
import { MovieReducer } from './features/movie/store/movie.reducer';

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
    provideStore({ auth: authReducer, movies: MovieReducer }),
    provideEffects([AuthEffects, MovieEffects]),
  ],
};
