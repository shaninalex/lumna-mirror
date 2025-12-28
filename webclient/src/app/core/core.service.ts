import {inject, Injectable, OnInit} from '@angular/core';
import {Dispatcher, Events} from '@ngrx/signals/events';
import {sessionEvents} from '@core/store/session.store';
import {tap} from 'rxjs';
import {projectEvents, ProjectStore} from '@entities/project';
import {UserStore} from '@entities/user';
import {BoardStore} from '@entities/board';
import {ListStore} from '@entities/list';

@Injectable({
    providedIn: 'root',
})
export class CoreService {
    readonly events = inject(Events);
    readonly dispatcher = inject(Dispatcher)

    // Init stores
    readonly userStore = inject(UserStore)
    readonly projectStore = inject(ProjectStore)
    readonly boardStore = inject(BoardStore)
    readonly listStore = inject(ListStore)

    constructor() {
        this.events
            .on(sessionEvents.authenticated)
            .pipe(
                tap(() => {
                    console.log("CoreService: start app initialization")
                    this.dispatcher.dispatch(projectEvents.getProjects())
                })
            )
            .subscribe();
    }
}
