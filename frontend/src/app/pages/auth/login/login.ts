import { Component, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { SessionStore } from '@core/store/session.store';
import { UserStore } from '@entities/user';
import { AuthService } from '../api/auth.service';
import { AuthLayout } from '@core/layouts';
import { Store } from '@ngrx/store';
import { actionSessionAuthenticate } from '@core/store/session.actions';

@Component({
    selector: 'app-login',
    imports: [FormsModule, AuthLayout],
    templateUrl: './login.html',
    styleUrl: './login.css',
})
export class Login {
    private authService = inject(AuthService);
    private sessionStore = inject(SessionStore);
    private store = inject(Store);
    private userStore = inject(UserStore);
    private router = inject(Router);

    email = '';
    password = '';

    onSubmit(): void {
        this.authService.login(this.email, this.password).subscribe({
            next: (user) => {
                this.store.dispatch(actionSessionAuthenticate({ user: user }));
                this.sessionStore.setAuthenticated(true);
                this.userStore.setUser(user);
                this.router.navigate(['/']);
            },
        });
    }
}
