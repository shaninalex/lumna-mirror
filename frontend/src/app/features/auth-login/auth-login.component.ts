import { Component, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { actionSessionAuthenticateStart } from '@core/store/session.actions';
import { Store } from '@ngrx/store';

@Component({
    selector: 'auth-login-feature',
    imports: [FormsModule],
    templateUrl: './auth-login.component.html',
})
export class AuthLoginFeature {
    private store = inject(Store);
    email = '';
    password = '';

    onSubmit(): void {
        this.store.dispatch(
            actionSessionAuthenticateStart({
                email: this.email,
                password: this.password,
            }),
        );
    }
}
