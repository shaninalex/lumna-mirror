import {Routes} from '@angular/router';
import {AuthRoot} from './auth-root';
import {Login} from './login/login';
import {Register} from './register/registration';
import {Verification} from './verification/verification';
import {loginFlowResolver} from '@client/pages/auth/login/resolver';
import {registrationFlowResolver} from '@client/pages/auth/register/resolver';
import {verificationFlowResolver} from '@client/pages/auth/verification/resolver';

export const authRoutes: Routes = [
    {
        path: "auth",
        component: AuthRoot,
        children: [
            {
                path: "login",
                component: Login,
                resolve: { loginForm: loginFlowResolver }
            },
            {
                path: "registration",
                component: Register,
                resolve: { registrationForm: registrationFlowResolver }
            },
            {
                path: "verification",
                component: Verification,
                resolve: { verificationForm: verificationFlowResolver }
            }
        ]
    }
];
