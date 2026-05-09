import { createReducer, on } from '@ngrx/store';

import { AuthState } from '../../../core/models/auth.models';
import {
  loginFailed,
  loginRequested,
  loginSucceeded,
  logoutSucceeded,
  meFailed,
  meRequested,
  meSucceeded,
  refreshFailed,
  refreshRequested,
  refreshSucceeded,
  registerFailed,
  registerRequested,
  registerSucceeded,
} from './auth.actions';

export const initialAuthState: AuthState = {
  user: null,
  status: 'idle',
  error: null,
};

export const authReducer = createReducer(
  initialAuthState,
  on(loginRequested, registerRequested, refreshRequested, meRequested, (state) => ({
    ...state,
    status: 'loading',
    error: null,
  })),
  on(loginSucceeded, registerSucceeded, refreshSucceeded, meSucceeded, (state, { response }) => ({
    ...state,
    user: response.user,
    status: 'authenticated',
    error: null,
  })),
  on(logoutSucceeded, (state) => ({
    ...state,
    user: null,
    status: 'idle',
    error: null,
  })),
  on(loginFailed, registerFailed, refreshFailed, meFailed, (state, { error }) => ({
    ...state,
    user: null,
    status: 'error',
    error,
  })),
);
