import {Component, inject, Input} from '@angular/core';
import {UserStore} from '@entities/user';
import {CdkMenu, CdkMenuItem, CdkMenuTrigger,} from '@angular/cdk/menu';
import {Dispatcher} from '@ngrx/signals/events';
import {sessionEvents} from '@core/store/session.store';

@Component({
    selector: 'app-header',
    imports: [
        CdkMenuTrigger,
        CdkMenu,
        CdkMenuItem,
    ],
    templateUrl: './header.html',
    styleUrl: './header.css',
})
export class Header {
    @Input() title: string
    readonly userStore = inject(UserStore);
    readonly user = this.userStore.user;
    private readonly dispatcher = inject(Dispatcher);

    logout(): void {
        this.dispatcher.dispatch(sessionEvents.logout())
    }
}
