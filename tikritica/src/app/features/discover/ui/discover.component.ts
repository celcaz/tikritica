import { Component, inject } from '@angular/core';
import { MovieFacade } from '../../movie/store/movie.facade';
import { MovieSignalStore } from '../../movie/state/movie.store';

@Component({
  selector: 'app-discover',
  standalone: true,
  imports: [],
  templateUrl: './discover.component.html',
})
export class DiscoverComponent {
  protected readonly movieFacade = inject(MovieFacade);
  movieStore = inject(MovieSignalStore);
  constructor() {
    this.movieFacade.getMovies();
    this.movieStore.getMovies();
  }
}
