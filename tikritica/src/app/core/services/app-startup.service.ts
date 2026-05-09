import { inject, Injectable } from '@angular/core';
import { AuthFacade } from '../../features/auth/state/auth.facade';

@Injectable({ providedIn: 'root' })
export class AppStartupService {
  private readonly authFacade = inject(AuthFacade);

  init(): void {
    this.authFacade.refresh();
  }
}
