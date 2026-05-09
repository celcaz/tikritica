import { createAction, props } from '@ngrx/store';

import { AuthUserResponse, LoginRequest, RegisterRequest } from '../domain/auth.models';

export const loginRequested = createAction(
  '[Auth] Login Requested',
  props<{ payload: LoginRequest }>(),
);

export const loginSucceeded = createAction(
  '[Auth] Login Succeeded',
  props<{ response: AuthUserResponse }>(),
);

export const loginFailed = createAction('[Auth] Login Failed', props<{ error: string }>());

export const registerRequested = createAction(
  '[Auth] Register Requested',
  props<{ payload: RegisterRequest }>(),
);

export const registerSucceeded = createAction(
  '[Auth] Register Succeeded',
  props<{ response: AuthUserResponse }>(),
);

export const registerFailed = createAction('[Auth] Register Failed', props<{ error: string }>());

export const refreshRequested = createAction('[Auth] Refresh Requested');

export const refreshSucceeded = createAction(
  '[Auth] Refresh Succeeded',
  props<{ response: AuthUserResponse }>(),
);

export const refreshFailed = createAction('[Auth] Refresh Failed', props<{ error: string }>());

export const logoutRequested = createAction('[Auth] Logout Requested');

export const logoutSucceeded = createAction('[Auth] Logout Succeeded');

export const logoutFailed = createAction('[Auth] Logout Failed', props<{ error: string }>());

export const meRequested = createAction('[Auth] Me Requested');

export const meSucceeded = createAction(
  '[Auth] Me Succeeded',
  props<{ response: AuthUserResponse }>(),
);

export const meFailed = createAction('[Auth] Me Failed', props<{ error: string }>());
