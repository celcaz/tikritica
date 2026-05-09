import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Movie } from '../models/movie.model';
import { Observable } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class MovieService {
  private readonly baseUrl = '/api';
  private readonly http = inject(HttpClient);

  constructor() {}
  getMovies(): Observable<Movie[]> {
    return this.http.get<Movie[]>(`${this.baseUrl}/movies`);
  }

  getMoviesBySlug(slug: string): Observable<Movie> {
    return this.http.get<Movie>(`${this.baseUrl}/movies/${slug}`);
  }
}
