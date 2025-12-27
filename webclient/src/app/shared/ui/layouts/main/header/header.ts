import {ChangeDetectionStrategy, Component, inject, Input} from '@angular/core';
import {UserStore} from '@entities/user';
import {Dispatcher} from '@ngrx/signals/events';
import {sessionEvents} from '@core/store/session.store';
import {UiService} from '@shared/ui';
import {CdkMenuTrigger} from '@angular/cdk/menu';

@Component({
    selector: 'app-header',
    imports: [
        CdkMenuTrigger
    ],
    templateUrl: './header.html',
    styleUrl: './header.css',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class Header {
    @Input() title: string
    readonly userStore = inject(UserStore);
    private ui: UiService = inject(UiService)
    readonly user = this.userStore.user;
    private readonly dispatcher = inject(Dispatcher);

    logout(): void {
        this.dispatcher.dispatch(sessionEvents.logout())
    }
}
