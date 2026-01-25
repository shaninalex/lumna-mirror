import { Component, inject } from '@angular/core';
import { Router } from '@angular/router';
import { UiService } from '@shared/ui';
import { toSignal } from '@angular/core/rxjs-interop';
import { UserService, UserStore } from '@entities/user';
import { CdkMenuTrigger } from '@angular/cdk/menu';

@Component({
    selector: 'app-header',
    imports: [CdkMenuTrigger],
    templateUrl: './header.html',
    styleUrl: './header.css',
})
export class Header {
    private readonly ui = inject(UiService);
    private readonly router = inject(Router);
    private readonly userService = inject(UserService);
    readonly title = toSignal(this.ui.pageTitle);

    readonly userStore = inject(UserStore);
    readonly user = this.userStore.user;

    logout(): void {
        this.userService.logout().subscribe(() => {
            this.userStore.setUser(null);
            this.router.navigate(['/auth/login']);
        });
    }
}
