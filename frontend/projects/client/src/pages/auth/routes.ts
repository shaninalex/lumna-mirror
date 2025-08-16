import {Routes} from '@angular/router';
import {AuthComponent} from '@client/pages/auth/auth.component';
import {LoginComponent} from '@client/pages/auth/login.component';
import {VerificationComponent} from '@client/pages/auth/verification.component';
import {RegistrationComponent} from '@client/pages/auth/registration.component';
import {RecoveryComponent} from '@client/pages/auth/recovery.component';
import {ErrorComponent} from '@client/pages/auth/error.component';

export const authRoutes: Routes = [
    {
        path: 'auth',
        component: AuthComponent,
        children: [
            {
                path: 'login',
                component: LoginComponent,
            },
            {
                path: 'verification',
                component: VerificationComponent,
            },
            {
                path: 'registration',
                component: RegistrationComponent,
            },
            {
                path: 'recovery',
                component: RecoveryComponent,
            },
            {
                path: 'error',
                component: ErrorComponent,
            },
        ],
    },
];
