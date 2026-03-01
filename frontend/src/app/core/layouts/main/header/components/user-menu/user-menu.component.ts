import {Component, inject} from '@angular/core';
import {Store} from '@ngrx/store';
import {selectUser, UserState} from '@entities/user';
import { actionSessionLoggingOut } from '@core/store/index';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';
import {AsyncPipe} from '@angular/common';

@Component({
    selector: 'app-user-menu',
    imports: [CdkMenuTrigger, CdkMenu, CdkMenuItem, AsyncPipe],
    template: `
        @if (user | async; as user) {
            <button [cdkMenuTriggerFor]="menu" class="btn btn-sm">
                <div class="d-flex align-items-center gap-2">
                    <span class="text-sm">{{ user.email }}</span>
                    <img class="rounded-circle" src="img/7.png" style="width: 32px; height: 32px;"/>
                </div>
            </button>

            <ng-template #menu>
                <div class="dropdown-menu d-block" cdkMenu>
                    <button (click)="logout()" class="dropdown-item" cdkMenuItem>Sign out</button>
                </div>
            </ng-template>
        }
    `
})
export class UserMenuComponent {
    readonly store = inject(Store<UserState>);
    readonly user = this.store.select(selectUser);

    logout(): void {
        this.store.dispatch(actionSessionLoggingOut());
    }
}
