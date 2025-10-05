import {Component, inject} from '@angular/core';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {filter, map, Observable} from 'rxjs';
import {AsyncPipe} from '@angular/common';
import {RouterLink} from '@angular/router';
import {UserLogoutFeature} from '@client/features/user';
import {selectUser} from '@client/entities/user';
import {OverlayModule} from '@angular/cdk/overlay';

@Component({
    selector: 'fr-header',
    imports: [
        AsyncPipe,
        RouterLink,
        UserLogoutFeature,
        OverlayModule,
    ],
    template: `
        <div class="py-2 px-4 flex items-center justify-between border-b border-gray-300">
            <div class="flex items-center gap-2 ms-auto">
                @if (email$ | async; as email) {
                    <div class="dropdown">
                        <div tabindex="0" role="button" class="flex items-center gap-2 cursor-pointer" (click)="isOpen = !isOpen"
                             cdkOverlayOrigin #trigger="cdkOverlayOrigin">
                            <div>{{ email }}</div>
                            <img src="img/1.png" class="rounded-full w-8" alt="">
                        </div>
                        <ng-template
                            cdkConnectedOverlay
                            cdkConnectedOverlayBackdropClass="cdk-overlay-transparent-backdrop"
                            [cdkConnectedOverlayOrigin]="trigger"
                            [cdkConnectedOverlayOpen]="isOpen"
                            [cdkConnectedOverlayHasBackdrop]="true"
                            (backdropClick)="isOpen = false"
                        >
                            <ul class="p-4 bg-white border rounded">
                                <li><a [routerLink]="['settings']">Settings</a></li>
                                <li><kr-user-logout-feature /></li>
                            </ul>
                        </ng-template>
                    </div>
                }
            </div>
        </div>
    `
})
export class HeaderComponent {
    private store: Store<AppState> = inject(Store<AppState>);
    email$: Observable<string> = this.store.select(selectUser).pipe(
        filter(user => !!user),
        map(user => user.email)
    )
    isOpen = false;
}
