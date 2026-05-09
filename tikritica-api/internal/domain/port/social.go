package port

import "context"

type SocialRepository interface {
	Follow(ctx context.Context, followerID, followingID string) error
	Unfollow(ctx context.Context, followerID, followingID string) error
}
