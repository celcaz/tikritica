import { Component, effect, inject, input, output } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { UserProfile } from '../../domain/user.domain';
import { InputText } from 'primeng/inputtext';

@Component({
  selector: 'app-profile-edit',
  imports: [ReactiveFormsModule, InputText],
  templateUrl: './profile-edit.html',
})
export class ProfileEditComponent {
  private readonly fb = inject(FormBuilder);
  submitInfo = output<UserProfile>();
  readonly userInfo = input.required<UserProfile | null>();

  userInfoForm: FormGroup = this.fb.group({
    displayName: ['', Validators.required],
    bio: [''],
    avatarUrl: [''],
  });

  constructor() {
    effect(() => {
      const user = this.userInfo();
      if (user) {
        this.userInfoForm.patchValue({
          displayName: user.displayName,
          bio: user.bio,
          avatarUrl: user.avatarUrl,
        });
      }
    });
  }

  onSubmit() {
    if (this.userInfoForm.invalid) {
      this.userInfoForm.markAllAsTouched();
      return;
    }
    this.submitInfo.emit(this.userInfoForm.value);
  }
}
