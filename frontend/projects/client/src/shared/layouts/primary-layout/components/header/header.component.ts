import {Component, inject, OnInit} from '@angular/core';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {filter, map, Observable} from 'rxjs';
import {AsyncPipe} from '@angular/common';
import {NavigationStart, Router, RouterLink, RouterLinkActive} from '@angular/router';
import {UserLogoutFeature} from '@client/features/user';
import {selectUser} from '@client/entities/user';
import {OverlayModule} from '@angular/cdk/overlay';
import {UiService} from '@client/shared/ui/ui.service';
import {BreadcrumbsComponent} from './breadcrumbs';
import {ThemeSwitcherComponent} from '@client/shared/layouts/primary-layout/components/header/theme-switcher.component';

@Component({
    selector: 'lu-header',
    imports: [
        AsyncPipe,
        RouterLink,
        UserLogoutFeature,
        OverlayModule,
        RouterLinkActive,
        BreadcrumbsComponent,
        ThemeSwitcherComponent,
    ],
    template: `
        <div class="py-2 px-4 flex items-center justify-between border-b border-gray-300 dark:bg-gray-800 dark:border-gray-600">
            <div class="flex items-center gap-4">
                <button (click)="toggleSidebar()" class="cursor-pointer">
                    @if (closeSidebar) {
                        <i class="i-arrow-right"></i>
                    } @else {
                        <i class="i-arrow-left"></i>
                    }
                </button>
                <lu-breadcrumbs />
            </div>
            <div class="flex items-center gap-2 ms-auto">
                <lu-theme-switcher />
                @if (email$ | async; as email) {
                    <div tabindex="0" role="button" class="flex items-center gap-2 cursor-pointer" (click)="isOpen = !isOpen"
                         cdkOverlayOrigin #trigger="cdkOverlayOrigin">
                        <div>{{ email }}</div>
                        <img src="img/1.png" class="rounded-full w-8" alt="">
                    </div>
                    <ng-template
                        cdkConnectedOverlay
                        [cdkConnectedOverlayOrigin]="trigger"
                        [cdkConnectedOverlayOpen]="isOpen"
                        [cdkConnectedOverlayHasBackdrop]="true"
                        (backdropClick)="isOpen = false"
                    >
                        <ul class="dropdown">
                            <li>
                                <a [routerLink]="['settings']" [routerLinkActive]="'active-link'">
                                    <i class="i-settings"></i>
                                    Settings
                                </a>
                            </li>
                            <li><kr-user-logout-feature /></li>
                        </ul>
                    </ng-template>
                }
            </div>
        </div>
    `
})
export class HeaderComponent implements OnInit {
    private uiService = inject(UiService);
    private store: Store<AppState> = inject(Store<AppState>);
    private router = inject(Router);

    email$: Observable<string> = this.store.select(selectUser).pipe(
        filter(user => !!user),
        map(user => user.email)
    )
    isOpen = false;
    closeSidebar: boolean = false;

    toggleSidebar() {
        this.closeSidebar = !this.closeSidebar;
        this.uiService.setExtendSidebar(this.closeSidebar)
    }

    ngOnInit() {
        this.router.events.pipe(filter(e => e instanceof NavigationStart)).subscribe(() => this.isOpen = false);
    }
}
