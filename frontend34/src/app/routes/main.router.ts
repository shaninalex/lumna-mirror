import {Routes} from '@angular/router';

import {appRouter} from './app';
import {authRouter} from './auth';

export const mainRoutes: Routes = [
    ...authRouter,
    ...appRouter,
];
