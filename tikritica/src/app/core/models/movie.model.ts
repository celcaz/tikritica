export interface Movie {
  id: string;
  title: string;
  slug: string;
  releaseDate: string;
  rating: number;
  posterUrl: string;
}

export type MovieStatus = 'success' | 'loading' | 'error';

export interface MovieState {
  movies: Movie[];
  selectedMovie: Movie | null;
  status: MovieStatus;
  error: string | null;
}
