import { Component, inject } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { Store } from '@ngrx/store';
import { AsyncPipe } from '@angular/common';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';

import { UiService } from '@shared/ui';
import { selectUser, UserState } from '@entities/user';
import { actionSessionLoggingOut } from '@core/store/index';
import { ThemeSwitcher } from './components';
import { BreadCrumbs } from './components/breadcrumbs/breadcrumbs.component';

@Component({
    selector: 'app-header',
    imports: [CdkMenuTrigger, CdkMenu, CdkMenuItem, AsyncPipe, ThemeSwitcher, BreadCrumbs],
    templateUrl: './header.html',
    styleUrl: './header.css',
})
export class Header {
    private readonly ui = inject(UiService);
    readonly title = toSignal(this.ui.pageTitle);
    readonly store = inject(Store<UserState>);
    readonly user = this.store.select(selectUser);

    logout(): void {
        this.store.dispatch(actionSessionLoggingOut());
    }
}
