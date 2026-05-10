import { Component, inject } from '@angular/core';
import { ThemeService } from '../../services/theme.service';
import { AuthFacade } from '../../../features/auth/state/auth.facade';

@Component({
  selector: 'app-header',
  standalone: true,
  templateUrl: './header.component.html',
})
export class HeaderComponent {
  protected readonly themeService = inject(ThemeService);
  readonly authFacade = inject(AuthFacade);
}
