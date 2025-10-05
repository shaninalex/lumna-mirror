import {Component, inject} from '@angular/core';
import {LoaderComponent} from '@client/shared/ui/loader';
import {filter, finalize, tap} from 'rxjs';
import {APIResponse} from '@client/shared/models';
import {Router} from '@angular/router';
import {UserService} from '@client/entities/user';

@Component({
    selector: 'kr-user-logout-feature',
    imports: [
        LoaderComponent
    ],
    template: `
        <button class="cursor-pointer" (click)="logoutAction()">
            @if (!loading) {
                <i class="i-logout"></i>
                Logout
            } @else {
                <ui-loader />
            }
        </button>
        @if (errors) {
            @for (err of errors; track $index) {
                <div class="text-warning text-sm">{{ err }}</div>
            }
        }
    `,
})
export class UserLogoutFeature {
    loading: boolean = false;
    errors: string[] = [];
    api = inject(UserService);
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
                this.router.navigate(['/auth/login'])
            },
        })
    }
}
