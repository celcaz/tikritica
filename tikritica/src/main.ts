import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { App } from './app/app';
import { AppStartupService } from './app/core/services/app-startup.service';

bootstrapApplication(App, appConfig)
  .then((appRef) => {
    // Step 1: Kick off session refresh after bootstrap.
    const startup = appRef.injector.get(AppStartupService);
    startup.init();
  })
  .catch((err) => console.error(err));
