import {Component, inject, OnDestroy, OnInit} from '@angular/core';
import {MainLayout} from '@client/app/layouts';
import {RouterOutlet} from '@angular/router';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {selectSession} from '@client/entities/session';
import {LoaderComponent} from '@dev/ui/loader';
import {filter, tap, map, Subscription, take} from 'rxjs';

@Component({
    selector: 'fr-root',
    imports: [
        MainLayout,
        RouterOutlet,
        LoaderComponent
    ],
    template: `
        @if (ready) {
            <fr-main-layout>
                <router-outlet/>
            </fr-main-layout>
        } @else {
            <ui-loader />
        }
    `
})
export class MainRoot implements OnInit, OnDestroy {
    private store: Store<AppState> = inject(Store<AppState>);
    private _sub: Subscription = new Subscription();
    ready = false;
    ngOnInit() {
        this._sub.add(
            this.store.select(selectSession).pipe(
                filter(session => !!session),
                tap(() => {
                    this.ready = true;
                }),
            ).subscribe()
        )
    }

    ngOnDestroy() {
        this._sub.unsubscribe();
    }
}
