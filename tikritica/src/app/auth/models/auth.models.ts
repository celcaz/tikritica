export interface User {
  id?: string;
  username?: string;
  displayName?: string;
  avatarUrl?: string;
  bio?: string;
  followersCount?: number;
  followingCount?: number;
  listsCount?: number;
  reviewsCount?: number;
  createdAt?: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  username: string;
  displayName?: string;
  email: string;
  password: string;
}

export interface AuthUserResponse {
  user: User;
}

export type AuthStatus = 'idle' | 'loading' | 'authenticated' | 'error';

export interface AuthState {
  user: User | null;
  status: AuthStatus;
  error: string | null;
}
