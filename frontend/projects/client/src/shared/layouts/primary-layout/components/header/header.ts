import {Component, inject} from '@angular/core';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {map, Observable} from 'rxjs';
import {AsyncPipe} from '@angular/common';
import {RouterLink} from '@angular/router';
import {UserLogoutFeature} from '@client/features/user';
import {selectUser} from '@client/entities/user';

import { ThemeSwitcherComponent } from './theme-switcher.component'

@Component({
    selector: 'fr-header',
    imports: [
        AsyncPipe,
        RouterLink,
        UserLogoutFeature,
        ThemeSwitcherComponent,
    ],
    template: `
        <div class="py-2 px-4 flex items-center justify-between border-b border-base-300 bg-base-100">
            <div class="flex items-center gap-2 ms-auto">
                <fr-theme-switcher />
                @if (code$ | async; as code) {
                    <div class="dropdown">
                        <div tabindex="0" role="button" class="flex items-center gap-2">
                            <div>{{ code }}</div>
                            <img src="img/1.png" class="rounded-full w-8" alt="">
                        </div>
                        <ul tabindex="0" class="dropdown-content right-0 menu bg-base-100 rounded-box z-1 w-36 p-2 shadow-sm">
                            <li><a [routerLink]="['settings']">Settings</a></li>
                            <li><kr-user-logout-feature /></li>
                        </ul>
                    </div>
                }
            </div>
        </div>
    `
})
export class Header {
    private store: Store<AppState> = inject(Store<AppState>);
    code$: Observable<string | undefined> = this.store.select(selectUser).pipe(
        map(user => user?.code)
    )
}
