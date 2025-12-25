import {Component, inject} from '@angular/core';
import {RouterOutlet} from '@angular/router';
import {SessionStore} from '@core/store/session.store';
import {TokenRefreshService} from '@core/api/refresh.service';
import {toObservable} from '@angular/core/rxjs-interop';
import {tap} from 'rxjs';

@Component({
    selector: 'app-root',
    imports: [RouterOutlet],
    templateUrl: './app.html',
    styleUrl: './app.css',
})
export class App {
    readonly sessionStore = inject(SessionStore);
    readonly tokenRefreshService = inject(TokenRefreshService);

    constructor() {
        toObservable(this.sessionStore.status)
            .pipe(tap(status => console.log('App: ', status)))
            .subscribe();
    }
}
