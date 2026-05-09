import { inject, Injectable } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { Store } from '@ngrx/store';

import { MovieState } from '../../../core/models/movie.model';
import { movieBySlugRequested, moviesRequested } from './movie.actions';
import { selectMovies, selectMoviesError, selectMoviesStatus, selectSelectedMovie } from './movie.selectors';

@Injectable({ providedIn: 'root' })
export class MovieFacade {
  private readonly store = inject(Store);

  readonly movies = toSignal(this.store.select(selectMovies), { initialValue: [] });
  readonly selectedMovie = toSignal(this.store.select(selectSelectedMovie), { initialValue: null });
  readonly status = toSignal(this.store.select(selectMoviesStatus), { initialValue: 'loading' as MovieState['status'] });
  readonly error = toSignal(this.store.select(selectMoviesError), { initialValue: null });

  getMovies(): void {
    this.store.dispatch(moviesRequested());
  }

  getMovieBySlug(slug: string): void {
    this.store.dispatch(movieBySlugRequested({ slug }));
  }
}
