import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router, RouterLink } from '@angular/router';

import { AuthFacade } from '../../services/auth.facade';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './register.component.html',
})
export class RegisterComponent {
  private readonly fb = inject(FormBuilder);
  private readonly authFacade = inject(AuthFacade);
  private readonly router = inject(Router);
  constructor() {
    // Step 1: Build the reactive form for registration.
    this.form = this.fb.group({
      username: ['', [Validators.required]],
      displayName: [''],
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(8)]],
    });

    this.status$ = this.authFacade.status$;
    this.error$ = this.authFacade.error$;
  }

  readonly form;
  readonly status$;
  readonly error$;

  // Step 2: Submit register request through facade.
  submit(): void {
    if (this.form.invalid) {
      this.form.markAllAsTouched();
      return;
    }

    const { username, displayName, email, password } = this.form.getRawValue();
    if (!username || !email || !password) {
      return;
    }

    this.authFacade.register({
      username,
      displayName: displayName || undefined,
      email,
      password,
    });
  }

  // Step 3: Navigate back to login.
  goToLogin(): void {
    this.router.navigate(['/login']);
  }
}
