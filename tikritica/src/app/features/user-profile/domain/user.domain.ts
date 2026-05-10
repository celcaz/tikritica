export interface UserProfile {
  avatarUrl: string;
  bio: string;
  createdAt: string;
  displayName: string;
  followersCount: number;
  followingCount: number;
  id: string;
  listsCount: number;
  reviewsCount: number;
  username: string;
}

export type ProfileStatus = 'loading' | 'success' | 'error';

export interface ProfileState {
  profile: UserProfile | null;
  status: ProfileStatus;
  error: string | null;
}
