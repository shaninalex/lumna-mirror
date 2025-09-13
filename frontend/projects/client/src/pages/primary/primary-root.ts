import {Component, inject, OnInit} from '@angular/core';
import {PrimaryLayout} from '@client/shared/layouts';
import {RouterOutlet} from '@angular/router';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {filter, Observable, map} from 'rxjs';
import {LoaderComponent} from '@client/shared/ui/loader';
import {selectSession} from '@client/entities/auth';
import {GetUserAction} from '@client/entities/user';
import {AsyncPipe} from '@angular/common';

@Component({
    selector: 'fr-root',
    imports: [
        PrimaryLayout,
        RouterOutlet,
        LoaderComponent,
        AsyncPipe
    ],
    template: `
        @if (ready$ | async) {
            <fr-primary-layout>
                <router-outlet/>
            </fr-primary-layout>
        } @else {
            <ui-loader />
        }
    `
})
export class PrimaryRoot implements OnInit {
    private store: Store<AppState> = inject(Store<AppState>);
    ready$: Observable<boolean> = this.store.select(selectSession).pipe(
        filter(session => !!session),
        map(session => !!session),
    );

    ngOnInit() {
        this.store.dispatch(GetUserAction());
    }
}
