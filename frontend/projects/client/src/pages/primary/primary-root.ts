import {Component, inject, OnDestroy, OnInit} from '@angular/core';
import {PrimaryLayout} from '@client/app/layouts';
import {RouterOutlet} from '@angular/router';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {selectSession} from '@client/entities/session';
import {LoaderComponent} from '@dev/ui/loader';
import {filter, tap, Subscription} from 'rxjs';

@Component({
    selector: 'fr-root',
    imports: [
        PrimaryLayout,
        RouterOutlet,
        LoaderComponent
    ],
    template: `
        @if (ready) {
            <fr-primary-layout>
                <router-outlet/>
            </fr-primary-layout>
        } @else {
            <ui-loader />
        }
    `
})
export class PrimaryRoot implements OnInit, OnDestroy {
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
