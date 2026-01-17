import { Component, inject } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { SessionStore } from '@core/store/session.store';
import { CoreService } from '@core/core.service';

@Component({
    selector: 'app-root',
    imports: [RouterOutlet],
    templateUrl: './app.html',
    styleUrl: './app.css',
})
export class App {
    readonly coreService = inject(CoreService);
    readonly sessionStore = inject(SessionStore);

    constructor() {
        // toObservable(this.sessionStore.status)
        //     .pipe(tap(status => console.log('App: SessionStatus', status)))
        //     .subscribe();
    }
}
