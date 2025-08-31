
import { Routes } from '@angular/router';
import {SessionExpiredComponent} from './session-expired/session-expired.component';
import {LoginRequiredComponent} from './login-required/login-required.component';

export const systemRoutes: Routes = [
    {
        path: "session-expired",
        component: SessionExpiredComponent,
    },
    {
        path: "login-required",
        component: LoginRequiredComponent,
    }
];
