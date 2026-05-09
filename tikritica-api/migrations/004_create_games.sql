CREATE TABLE games (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug VARCHAR(255) UNIQUE NOT NULL,
    title VARCHAR(500) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    cover_url TEXT,
    release_date DATE,
    developer VARCHAR(255),
    publisher VARCHAR(255),
    platforms TEXT[] NOT NULL DEFAULT '{}',
    genres TEXT[] NOT NULL DEFAULT '{}',
    average_rating NUMERIC(3,2) NOT NULL DEFAULT 0,
    ratings_count INT NOT NULL DEFAULT 0,
    igdb_id INT UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_games_slug ON games(slug);
CREATE INDEX idx_games_igdb_id ON games(igdb_id);
