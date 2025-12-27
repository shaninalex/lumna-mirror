import {Component, inject} from '@angular/core';
import {UserStore} from '@entities/user';
import {Dispatcher} from '@ngrx/signals/events';
import {sessionEvents} from '@core/store/session.store';
import {UiService} from '@shared/ui';
import {CdkMenuTrigger} from '@angular/cdk/menu';
import {toSignal} from '@angular/core/rxjs-interop';

@Component({
    selector: 'app-header',
    imports: [
        CdkMenuTrigger
    ],
    templateUrl: './header.html',
    styleUrl: './header.css',
})
export class Header {
    private readonly ui = inject(UiService);
    readonly userStore = inject(UserStore);
    readonly user = this.userStore.user;
    private readonly dispatcher = inject(Dispatcher);
    readonly title = toSignal(this.ui.pageTitle);

    logout(): void {
        this.dispatcher.dispatch(sessionEvents.logout())
    }
}
