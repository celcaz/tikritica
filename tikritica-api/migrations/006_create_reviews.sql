CREATE TABLE reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    media_id UUID NOT NULL,
    media_type VARCHAR(10) NOT NULL CHECK (media_type IN ('movie', 'series', 'game', 'book')),
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    content TEXT NOT NULL DEFAULT '',
    contains_spoilers BOOLEAN NOT NULL DEFAULT false,
    likes_count INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE(user_id, media_id, media_type)
);

CREATE INDEX idx_reviews_user_id ON reviews(user_id);
CREATE INDEX idx_reviews_media ON reviews(media_id, media_type);
CREATE INDEX idx_reviews_created_at ON reviews(created_at DESC);
