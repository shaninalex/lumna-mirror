import {ChangeDetectionStrategy, Component, inject, signal} from '@angular/core';
import {debounce, email, Field, form, required} from '@angular/forms/signals';
import {Dispatcher, Events} from '@ngrx/signals/events';

import {LoginCredentials} from '../model/login.model';
import {Router} from '@angular/router';
import {sessionEvents} from '@core/store/session.store';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';

@Component({
    selector: 'app-login-form-feature',
    imports: [Field],
    templateUrl: './login-form.html',
    styleUrl: './login-form.css',
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LoginFormFeature {
    router = inject(Router)
    loading = signal<boolean>(false);
    loginModel = signal<LoginCredentials>({
        email: '',
        password: '',
    });
    errors = signal<string[]>([]);

    private readonly events = inject(Events);
    readonly dispatcher = inject(Dispatcher)

    loginForm = form(this.loginModel, (schemaPath) => {
        debounce(schemaPath.email, 500);
        required(schemaPath.email, {message: 'Email is required'});
        required(schemaPath.password, {message: 'Password is required'});
        email(schemaPath.email, {message: 'Email is in wrong format'});
    });

    constructor() {
        this.events
            .on(sessionEvents.authenticated)
            .pipe(takeUntilDestroyed())
            .subscribe(() => this.router.navigate(["/"]));

        this.events
            .on(sessionEvents.sessionFailed)
            .pipe(takeUntilDestroyed())
            .subscribe(data => {
                this.errors.set(data.payload.error.errors)
            });
    }

    submit(event: Event): void {
        event.preventDefault()
        this.loading.set(true)
        this.dispatcher.dispatch(sessionEvents.authenticate(this.loginModel()));
    }
}
