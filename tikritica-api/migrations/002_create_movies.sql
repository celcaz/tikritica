CREATE TABLE movies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug VARCHAR(255) UNIQUE NOT NULL,
    title VARCHAR(500) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    cover_url TEXT,
    release_date DATE,
    director VARCHAR(255),
    runtime INT,
    genres TEXT[] NOT NULL DEFAULT '{}',
    average_rating NUMERIC(3,2) NOT NULL DEFAULT 0,
    ratings_count INT NOT NULL DEFAULT 0,
    tmdb_id INT UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_movies_slug ON movies(slug);
CREATE INDEX idx_movies_tmdb_id ON movies(tmdb_id);
