import { Component, inject } from '@angular/core';
import { MovieFacade } from '../../movie/store/movie.facade';

@Component({
  selector: 'app-discover',
  standalone: true,
  imports: [],
  templateUrl: './discover.component.html',
})
export class DiscoverComponent {
  protected readonly movieFacade = inject(MovieFacade);

  constructor() {
    this.movieFacade.getMovies();
  }
}
