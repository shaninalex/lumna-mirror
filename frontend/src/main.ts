import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from '@core';
import { App } from '@root/src/app/app';

bootstrapApplication(App, appConfig)
    .catch((err) => console.error(err));
