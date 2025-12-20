import {ChangeDetectionStrategy, Component, inject, signal} from '@angular/core';
import {form, Field, required, email, debounce} from '@angular/forms/signals';

import {LoginCredentials} from '../model/login.model';
import {LoginService} from '../api/login.service';
import {map} from 'rxjs';

@Component({
    selector: 'app-login-form-feature',
    imports: [Field],
    templateUrl: './login-form.html',
    styleUrl: './login-form.css',
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LoginFormFeature {
    service = inject(LoginService)

    loginModel = signal<LoginCredentials>({
        email: '',
        password: '',
    });

    loginForm = form(this.loginModel, (schemaPath) => {
        debounce(schemaPath.email, 500);
        required(schemaPath.email, {message: 'Email is required'});
        required(schemaPath.password, {message: 'Password is required'});
        email(schemaPath.email, {message: 'Email is in wrong format'});
    });

    submit(event: Event): void {
        event.preventDefault()
        this.service.Login(this.loginModel()).subscribe(data => {
            console.log(data);
        })
    }
}
