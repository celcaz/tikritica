import { createAction, props } from '@ngrx/store';
import { Movie } from '../domain/movie.model';

export const moviesRequested = createAction('[Movie] Movies Requested');

export const moviesSucceeded = createAction(
  '[Movie] Movies Succeeded',
  props<{ movies: Movie[] }>(),
);

export const moviesFailed = createAction('[Movie] Movies Failed', props<{ error: string }>());

export const movieBySlugRequested = createAction(
  '[Movie] Movie By Slug Requested',
  props<{ slug: string }>(),
);

export const movieBySlugSucceeded = createAction(
  '[Movie] Movie By Slug Succeeded',
  props<{ movie: Movie }>(),
);

export const movieBySlugFailed = createAction(
  '[Movie] Movie By Slug Failed',
  props<{ error: string }>(),
);
