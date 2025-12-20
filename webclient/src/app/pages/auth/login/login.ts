import { Component } from '@angular/core';
import { AuthLayout } from '@shared/ui/layouts'
import {RouterLink} from '@angular/router';
import {LoginFormFeature} from '@features/auth';

@Component({
    selector: 'app-login',
    imports: [AuthLayout, RouterLink, LoginFormFeature],
    templateUrl: './login.html',
    styleUrl: './login.css',
})
export class Login {

}
