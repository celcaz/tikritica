import { Injectable } from '@angular/core';
import { signal, computed } from '@angular/core';
import { AuthService } from '../../auth/services/auth.service';

@Injectable({ providedIn: 'root' })
export class AuthSignalStore {
  // Example-only: minimal signal-based auth store for study.
  readonly user = signal<any | null>(null);
  readonly status = signal<'idle' | 'loading' | 'authenticated' | 'error'>('idle');
  readonly error = signal<string | null>(null);

  readonly isAuthenticated = computed(() => this.status() === 'authenticated' && !!this.user());

  constructor(private readonly authService: AuthService) {}

  // Step 1: Login using signals (example only).
  login(email: string, password: string): void {
    this.status.set('loading');
    this.error.set(null);

    this.authService.login({ email, password }).subscribe({
      next: (response) => {
        this.user.set(response.user);
        this.status.set('authenticated');
      },
      error: (err) => {
        this.user.set(null);
        this.status.set('error');
        this.error.set(err?.error?.error ?? 'unexpected error');
      },
    });
  }
}
