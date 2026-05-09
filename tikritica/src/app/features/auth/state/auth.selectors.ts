import { createFeatureSelector, createSelector } from '@ngrx/store';

import { AuthState } from '../domain/auth.models';

export const selectAuthState = createFeatureSelector<AuthState>('auth');

export const selectAuthUser = createSelector(selectAuthState, (state) => state.user);

export const selectAuthStatus = createSelector(selectAuthState, (state) => state.status);

export const selectAuthError = createSelector(selectAuthState, (state) => state.error);

export const selectIsAuthenticated = createSelector(
  selectAuthState,
  (state) => state.status === 'authenticated' && !!state.user,
);
