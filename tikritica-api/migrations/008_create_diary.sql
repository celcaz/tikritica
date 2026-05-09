CREATE TABLE diary_entries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    media_id UUID NOT NULL,
    media_type VARCHAR(10) NOT NULL CHECK (media_type IN ('movie', 'series', 'game', 'book')),
    watched_at DATE NOT NULL DEFAULT CURRENT_DATE,
    rating INT CHECK (rating >= 1 AND rating <= 5),
    review_id UUID REFERENCES reviews(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_diary_user_id ON diary_entries(user_id);
CREATE INDEX idx_diary_watched_at ON diary_entries(user_id, watched_at DESC);
