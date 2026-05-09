import { inject, Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { MovieRepository } from '../data/movie.repository';
import {
  movieBySlugFailed,
  movieBySlugRequested,
  movieBySlugSucceeded,
  moviesFailed,
  moviesRequested,
  moviesSucceeded,
} from './movie.actions';
import { catchError, map, of, switchMap } from 'rxjs';

@Injectable()
export class MovieEffects {
  actions$ = inject(Actions);
  movieRepository = inject(MovieRepository);

  getMovies$ = createEffect(() =>
    this.actions$.pipe(ofType(moviesRequested)).pipe(
      switchMap(() =>
        this.movieRepository.findAll().pipe(
          map((movies) => moviesSucceeded({ movies })),
          catchError((error) => of(moviesFailed({ error: this.toMessage(error) }))),
        ),
      ),
    ),
  );

  getMovieBySlug$ = createEffect(() =>
    this.actions$.pipe(ofType(movieBySlugRequested)).pipe(
      switchMap(({ slug }) =>
        this.movieRepository.findBySlug(slug).pipe(
          map((movie) => movieBySlugSucceeded({ movie })),
          catchError((error) => of(movieBySlugFailed({ error: this.toMessage(error) }))),
        ),
      ),
    ),
  );

  private toMessage(error: unknown): string {
    if (typeof error === 'string') {
      return error;
    }

    if (error && typeof error === 'object' && 'error' in error) {
      const httpError = error as { error?: { error?: string } };
      return httpError.error?.error ?? 'unexpected error';
    }

    return 'unexpected error';
  }
}
