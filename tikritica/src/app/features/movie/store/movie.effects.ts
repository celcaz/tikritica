import { inject, Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { MovieService } from '../../../core/services/movie.service';
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
  movieService = inject(MovieService);

  getMovies$ = createEffect(() =>
    this.actions$.pipe(ofType(moviesRequested)).pipe(
      switchMap(() =>
        this.movieService.getMovies().pipe(
          map((movies) => moviesSucceeded({ movies })),
          catchError((error) => of(moviesFailed({ error: this.toMessage(error) }))),
        ),
      ),
    ),
  );

  getMovieBySlug$ = createEffect(() =>
    this.actions$.pipe(ofType(movieBySlugRequested)).pipe(
      switchMap(({ slug }) =>
        this.movieService.getMoviesBySlug(slug).pipe(
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
