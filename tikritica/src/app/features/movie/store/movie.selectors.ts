import { createFeatureSelector, createSelector } from '@ngrx/store';
import { MovieState } from '../domain/movie.model';

export const selectMovieState = createFeatureSelector<MovieState>('movies');

export const selectMovies = createSelector(selectMovieState, (state) => state.movies);

export const selectMoviesStatus = createSelector(selectMovieState, (state) => state.status);

export const selectMoviesError = createSelector(selectMovieState, (state) => state.error);

export const selectSelectedMovie = createSelector(selectMovieState, (state) => state.selectedMovie);
