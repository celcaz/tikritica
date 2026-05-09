CREATE TABLE series (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug VARCHAR(255) UNIQUE NOT NULL,
    title VARCHAR(500) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    cover_url TEXT,
    release_date DATE,
    creator VARCHAR(255),
    seasons INT NOT NULL DEFAULT 0,
    episodes INT NOT NULL DEFAULT 0,
    genres TEXT[] NOT NULL DEFAULT '{}',
    average_rating NUMERIC(3,2) NOT NULL DEFAULT 0,
    ratings_count INT NOT NULL DEFAULT 0,
    tmdb_id INT UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_series_slug ON series(slug);
CREATE INDEX idx_series_tmdb_id ON series(tmdb_id);
