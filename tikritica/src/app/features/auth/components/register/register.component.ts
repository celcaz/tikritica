import { Component, inject } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';

import { AuthFacade } from '../../store/auth.facade';

@Component({
  selector: 'app-register',
  standalone: true,
  imports: [ReactiveFormsModule],
  templateUrl: './register.component.html',
})
export class RegisterComponent {
  private readonly fb = inject(FormBuilder);
  private readonly router = inject(Router);
  protected readonly auth = inject(AuthFacade);

  readonly form = this.fb.group({
    username: ['', [Validators.required]],
    displayName: [''],
    email: ['', [Validators.required, Validators.email]],
    password: ['', [Validators.required, Validators.minLength(8)]],
  });

  submit(): void {
    if (this.form.invalid) {
      this.form.markAllAsTouched();
      return;
    }

    const { username, displayName, email, password } = this.form.getRawValue();
    if (!username || !email || !password) return;

    this.auth.register({
      username,
      displayName: displayName || undefined,
      email,
      password,
    });
  }

  goToLogin(): void {
    this.router.navigate(['/login']);
  }
}
