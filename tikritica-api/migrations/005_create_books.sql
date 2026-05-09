CREATE TABLE books (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug VARCHAR(255) UNIQUE NOT NULL,
    title VARCHAR(500) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    cover_url TEXT,
    release_date DATE,
    author VARCHAR(255),
    pages INT,
    isbn VARCHAR(20) UNIQUE,
    genres TEXT[] NOT NULL DEFAULT '{}',
    average_rating NUMERIC(3,2) NOT NULL DEFAULT 0,
    ratings_count INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_books_slug ON books(slug);
CREATE INDEX idx_books_isbn ON books(isbn);
