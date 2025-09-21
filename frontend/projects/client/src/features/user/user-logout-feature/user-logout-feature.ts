import {Component, inject} from '@angular/core';
import {LoaderComponent} from '@client/shared/ui/loader';
import {filter, finalize, tap} from 'rxjs';
import {APIResponse} from '@client/shared/models';
import {Router} from '@angular/router';
import {UserService} from '@client/entities/user';
import {TokenService} from '@client/shared/common';

@Component({
    selector: 'kr-user-logout-feature',
    imports: [
        LoaderComponent
    ],
    template: `
        <button class="cursor-pointer" (click)="logoutAction()">
            @if (!loading) {
                Logout
            } @else {
                <ui-loader />
            }
        </button>
        @if (errors) {
            @for (err of errors; track $index) {
                <div class="text-red-500 text-sm">{{ err }}</div>
            }
        }
    `,
})
export class UserLogoutFeature {
    loading: boolean = false;
    errors: string[] = [];
    api = inject(UserService);
    tokenService = inject(TokenService)
    router = inject(Router);

    logoutAction(): void {
        this.errors = []
        this.loading = true
        this.api.logout().pipe(
            tap({
                error: (err: APIResponse<any>) => this.errors = err.messages,
            }),
            filter(resp => resp.status === true),
            finalize(() => this.loading = false),
        ).subscribe({
            next: data => {
                this.tokenService.removeAuthToken()
                this.router.navigate(['/auth/login'])
            },
        })
    }
}
