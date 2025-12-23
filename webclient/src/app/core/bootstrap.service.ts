import {Dispatcher} from '@ngrx/signals/events';
import {inject, Injectable} from '@angular/core';
import {toObservable} from '@angular/core/rxjs-interop';
import {filter, take, tap} from 'rxjs';
import {projectEvents, ProjectStore} from '@entities/project';
import {SessionStore} from './store/session.store'

@Injectable({providedIn: 'root'})
export class AppBootstrapService {
    private initialized = false;
    readonly sessionStore = inject(SessionStore)
    readonly dispatcher = inject(Dispatcher)

    private readonly projectStore = inject(ProjectStore)

    constructor() {
        this.listenSession();
    }

    private listenSession() {
        toObservable(this.sessionStore.status)
            .pipe(
                filter(status => status === 'authenticated'),
                take(1)
            )
            .subscribe(() => {
                if (!this.initialized) {
                    this.initialized = true;
                    this.dispatcher.dispatch(projectEvents.getProjects());
                }
            });
    }
}
