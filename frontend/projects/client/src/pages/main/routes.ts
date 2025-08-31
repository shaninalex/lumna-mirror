import {ResolveFn, Routes} from '@angular/router';
import {MainRoot} from './main-root';
import {Overview} from './overview/overview';
import {GetSessionAction} from '@client/entities/session';
import {Store} from '@ngrx/store';
import {inject} from '@angular/core';
import {AppState} from '@client/shared/store';

export const sessionResolver: ResolveFn<void> = () => {
    const store = inject(Store<AppState>);
    store.dispatch(GetSessionAction());
    return undefined;
};

export const mainRoutes: Routes = [
    {
        path: "",
        component: MainRoot,
        resolve: { session: sessionResolver },
        children: [
            {
                path: "",
                component: Overview,
            }
        ]
    }
];
