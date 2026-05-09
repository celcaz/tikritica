import { createReducer, on } from '@ngrx/store';
import { MovieState } from '../domain/movie.model';
import {
  movieBySlugFailed,
  movieBySlugRequested,
  movieBySlugSucceeded,
  moviesFailed,
  moviesRequested,
  moviesSucceeded,
} from './movie.actions';

export const initialMovieState: MovieState = {
  movies: [],
  selectedMovie: null,
  status: 'loading',
  error: null,
};

export const MovieReducer = createReducer(
  initialMovieState,
  on(moviesRequested, movieBySlugRequested, (state) => ({
    ...state,
    status: 'loading',
    error: null,
  })),
  on(moviesSucceeded, (state, { movies }) => ({
    ...state,
    movies,
    status: 'success',
    error: null,
  })),
  on(movieBySlugSucceeded, (state, { movie }) => ({
    ...state,
    selectedMovie: movie,
    status: 'success',
    error: null,
  })),
  on(moviesFailed, movieBySlugFailed, (state, { error }) => ({
    ...state,
    movies: [],
    status: 'error',
    error,
  })),
);
