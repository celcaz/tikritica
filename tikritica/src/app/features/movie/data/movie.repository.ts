import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { Movie } from '../domain/movie.model';
import { MovieApi } from './movie.api';

@Injectable({ providedIn: 'root' })
export class MovieRepository {
  private readonly api = inject(MovieApi);

  findAll(): Observable<Movie[]> {
    return this.api.getMovies();
  }

  findBySlug(slug: string): Observable<Movie> {
    return this.api.getMovieBySlug(slug);
  }
}
