import { computed, inject } from '@angular/core';
import { tapResponse } from '@ngrx/operators';
import { patchState, signalStore, withComputed, withMethods, withState } from '@ngrx/signals';
import { rxMethod } from '@ngrx/signals/rxjs-interop';
import { pipe, switchMap, tap } from 'rxjs';

import { Movie, MovieState } from '../../../core/models/movie.model';
import { MovieService } from '../../../core/services/movie.service';

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
  withMethods((store, movieService = inject(MovieService)) => ({
    getMovies: rxMethod<void>(
      pipe(
        tap(() => patchState(store, { status: 'loading', error: null })),
        switchMap(() =>
          movieService.getMovies().pipe(
            tapResponse({
              next: (movies: Movie[]) => patchState(store, { movies, status: 'success', error: null }),
              error: (err: unknown) => patchState(store, { status: 'error', error: toMessage(err) }),
            }),
          ),
        ),
      ),
    ),

    getMovieBySlug: rxMethod<string>(
      pipe(
        tap(() => patchState(store, { status: 'loading', error: null })),
        switchMap((slug) =>
          movieService.getMoviesBySlug(slug).pipe(
            tapResponse({
              next: (movie: Movie) => patchState(store, { selectedMovie: movie, status: 'success', error: null }),
              error: (err: unknown) => patchState(store, { status: 'error', error: toMessage(err) }),
            }),
          ),
        ),
      ),
    ),
  })),
);
