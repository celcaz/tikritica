import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { Movie } from '../domain/movie.model';

@Injectable({ providedIn: 'root' })
export class MovieApi {
  private readonly baseUrl = '/api';
  private readonly http = inject(HttpClient);

  getMovies(): Observable<Movie[]> {
    return this.http.get<Movie[]>(`${this.baseUrl}/movies`);
  }

  getMovieBySlug(slug: string): Observable<Movie> {
    return this.http.get<Movie>(`${this.baseUrl}/movies/${slug}`);
  }
}
