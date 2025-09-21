import {Routes} from '@angular/router';
import {LoginPageComponent} from '@client/pages/auth/login';
import {RegisterPageComponent} from '@client/pages/auth/register';
import {authPageGuard} from '@client/pages/auth/auth-page.guard';


export const authRoutes: Routes = [
    {
        path: "auth/login",
        component: LoginPageComponent,
        canActivate: [authPageGuard]
    },
    {
        path: "auth/register",
        component: RegisterPageComponent,
        canActivate: [authPageGuard]
    }
]
