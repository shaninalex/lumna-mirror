import { Component, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../api/auth.service';
import { AuthLayout } from '@core/layouts';

@Component({
    selector: 'app-login',
    imports: [FormsModule, AuthLayout],
    templateUrl: './login.html',
    styleUrl: './login.css',
})
export class Login {
    private authService = inject(AuthService);
    private router = inject(Router);

    email = '';
    password = '';

    onSubmit(): void {
        this.authService.login(this.email, this.password).subscribe({
            next: (user) => {
                this.router.navigate(['/']);
            },
        });
    }
}
