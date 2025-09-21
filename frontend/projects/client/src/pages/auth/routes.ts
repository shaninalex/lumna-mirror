import {Routes} from '@angular/router';
import {AuthService} from '@client/entities/auth';
import {LoginPageComponent} from '@client/pages/auth/login';
import {RegisterPageComponent} from '@client/pages/auth/register';


export const authRoutes: Routes = [
    {
        path: "auth/login",
        component: LoginPageComponent,
    },
    {
        path: "auth/register",
        component: RegisterPageComponent,
    }
]
