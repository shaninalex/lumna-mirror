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
        LoaderComponent,
    ],
    template: `
        <fr-auth-layout title="Register">
            <form [formGroup]="form" (ngSubmit)="onSubmit()" class="flex flex-col mt-4 items-start">
                <div class="mb-4">
                    <input class="input" type="email" formControlName="email" placeholder="Email"/>
                </div>
                <div class="mb-4">
                    <input class="input" type="password" formControlName="password" placeholder="Password"/>
                </div>
                <button class="btn btn-primary mb-4" [disabled]="!form.valid">
                    @if (loading) {
                        <ui-loader />
                    } @else {
                        Register
                    }
                </button>
                @if (errors) {
                    @for (err of errors; track $index) {
                        <div class="text-warning text-sm">{{ err }}</div>
                    }
                }
            </form>
            <hr class="mb-4">
            <a [routerLink]="['/auth/login']" class="text-sm">Login</a>
        </fr-auth-layout>
    `
})
export class RegisterPageComponent {
    api = inject(AuthService)
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
        this.api.register(this.form.value).pipe(
            tap({
                error: (err: APIResponse<any>) => this.errors = err.messages,
            }),
            filter(resp => resp.status === true),
            finalize(() => this.loading = false),
        ).subscribe({
            next: data => this.router.navigate(['/auth/login']),
        })
    }
}
