import {ChangeDetectionStrategy, Component, inject, signal} from '@angular/core';
import {form, Field, required, email, debounce} from '@angular/forms/signals';

import {LoginCredentials} from '../model/login.model';
import {LoginService} from '../api/login.service';
import {catchError, filter, finalize, map, of} from 'rxjs';
import {NONE_TYPE} from '@angular/compiler';
import {HttpErrorResponse} from '@angular/common/http';
import {Router} from '@angular/router';

@Component({
    selector: 'app-login-form-feature',
    imports: [Field],
    templateUrl: './login-form.html',
    styleUrl: './login-form.css',
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LoginFormFeature {
    service = inject(LoginService)
    router = inject(Router)
    loading = signal<boolean>(false);
    loginModel = signal<LoginCredentials>({
        email: '',
        password: '',
    });
    errors = signal<string[]>([]);

    loginForm = form(this.loginModel, (schemaPath) => {
        debounce(schemaPath.email, 500);
        required(schemaPath.email, {message: 'Email is required'});
        required(schemaPath.password, {message: 'Password is required'});
        email(schemaPath.email, {message: 'Email is in wrong format'});
    });

    submit(event: Event): void {
        event.preventDefault()
        this.loading.set(true)
        this.service.Login(this.loginModel()).pipe(
            catchError((err: HttpErrorResponse) => {
                if (!err.error.status) {
                    this.errors.set(err.error.errors)
                }
                return of(null)
            }),
            finalize(() => this.loading.set(false)),
            filter(data => !!data),
            map(data => {
                if (data.status) {
                    this.router.navigate(["/"])
                    // TODO: data.status - send to notification service
                }
            })
        ).subscribe()
    }
}
