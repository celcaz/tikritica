import { computed, inject } from '@angular/core';
import { tapResponse } from '@ngrx/operators';
import { patchState, signalStore, withComputed, withMethods, withState } from '@ngrx/signals';
import { rxMethod } from '@ngrx/signals/rxjs-interop';
import { pipe, switchMap, tap } from 'rxjs';

import { Movie, MovieState } from '../domain/movie.model';
import { MovieRepository } from '../data/movie.repository';

function toMessage(error: unknown): string {
  if (error && typeof error === 'object' && 'error' in error) {
    const httpError = error as { error?: { error?: string } };
    return httpError.error?.error ?? 'unexpected error';
  }
  return 'unexpected error';
}

export const MovieSignalStore = signalStore(
  { providedIn: 'root' },
  withState<MovieState>({
    movies: [],
    selectedMovie: null,
    status: 'loading',
    error: null,
  }),
  withComputed((store) => ({
    isLoading: computed(() => store.status() === 'loading'),
    hasError: computed(() => store.status() === 'error'),
    hasMovies: computed(() => store.movies().length > 0),
  })),
  withMethods((store, movieRepository = inject(MovieRepository)) => ({
    getMovies: rxMethod<void>(
      pipe(
        tap(() => patchState(store, { status: 'loading', error: null })),
        switchMap(() =>
          movieRepository.findAll().pipe(
            tapResponse({
              next: (movies: Movie[]) =>
                patchState(store, { movies, status: 'success', error: null }),
              error: (err: unknown) =>
                patchState(store, { status: 'error', error: toMessage(err) }),
            }),
          ),
        ),
      ),
    ),

    getMovieBySlug: rxMethod<string>(
      pipe(
        tap(() => patchState(store, { status: 'loading', error: null })),
        switchMap((slug) =>
          movieRepository.findBySlug(slug).pipe(
            tapResponse({
              next: (movie: Movie) =>
                patchState(store, { selectedMovie: movie, status: 'success', error: null }),
              error: (err: unknown) =>
                patchState(store, { status: 'error', error: toMessage(err) }),
            }),
          ),
        ),
      ),
    ),
  })),
);
