import { Injectable } from '@angular/core';
import { AuthFacade } from './auth/services/auth.facade';

@Injectable({ providedIn: 'root' })
export class AppStartupService {
  constructor(private readonly authFacade: AuthFacade) {}

  // Step 1: Restore session from auth cookie on app startup.
  init(): void {
    this.authFacade.refresh();
  }
}
