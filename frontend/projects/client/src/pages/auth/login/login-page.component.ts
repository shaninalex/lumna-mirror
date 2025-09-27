import {Component, inject} from '@angular/core';
import {AuthLayout} from '@client/shared/layouts';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {AuthService} from '@client/entities/auth';
import {Router, RouterLink} from '@angular/router';
import {filter, finalize, tap} from 'rxjs';
import {APIResponse} from '@client/shared/models';
import {LoaderComponent} from '@client/shared/ui';

@Component({
    selector: "fr-login-page",
    imports: [
        AuthLayout,
        ReactiveFormsModule,
        RouterLink,
        LoaderComponent
    ],
    template: `
        <fr-auth-layout title="Login">
            <form [formGroup]="form" (ngSubmit)="onSubmit()" class="mb-4">
                <input type="email" formControlName="email" class="input block mb-4" placeholder="email"/>
                <input type="password" formControlName="password" class="input block mb-4" placeholder="password"/>
                <button class="btn btn-primary mb-4" [disabled]="!form.valid">
                    @if (loading) {
                        <ui-loader />
                    } @else {
                        Login
                    }
                </button>
                @if (errors) {
                    @for (err of errors; track $index) {
                        <div class="text-warning text-sm">{{ err }}</div>
                    }
                }
            </form>
            <hr class="mb-4">
            <a [routerLink]="['/auth/register']" class="text-sm">Register</a>
        </fr-auth-layout>
    `
})
export class LoginPageComponent {
    api = inject(AuthService);

    router = inject(Router)
    form: FormGroup = new FormGroup({
        "email": new FormControl("", [Validators.required, Validators.email]),
        "password": new FormControl("", [Validators.required])
    })
    errors: string[] = [];
    loading: boolean = false;

    onSubmit(): void {
        this.loading = true;
        this.errors = [];
        this.api.login(this.form.value).pipe(
            tap({
                error: (err: APIResponse<any>) => this.errors = err.messages,
            }),
            filter(resp => resp.status === true),
            finalize(() => this.loading = false),
        ).subscribe({
            next: (data: APIResponse<{ "token": string }>) => {
                this.router.navigate(['/'])
            },
        })
    }
}
