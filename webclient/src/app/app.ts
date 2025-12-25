import {Component, inject} from '@angular/core';
import {RouterOutlet} from '@angular/router';
import {SessionStore} from '@core/store/session.store';
import {TokenRefreshService} from '@core/api/refresh.service';
import {toObservable} from '@angular/core/rxjs-interop';
import {tap} from 'rxjs';
import {CoreService} from '@core/core.service';
import {UserStore} from '@entities/user';

@Component({
    selector: 'app-root',
    imports: [RouterOutlet],
    templateUrl: './app.html',
    styleUrl: './app.css',
})
export class App {
    readonly coreService = inject(CoreService);
    readonly sessionStore = inject(SessionStore);
    readonly tokenRefreshService = inject(TokenRefreshService);

    constructor() {
        toObservable(this.sessionStore.status)
            .pipe(tap(status => console.log('App: SessionStatus', status)))
            .subscribe();
    }
}
