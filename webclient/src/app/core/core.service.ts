import { inject, Injectable, OnInit } from '@angular/core';
import { Dispatcher, Events } from '@ngrx/signals/events';
import { sessionEvents } from '@core/store/session.store';
import { tap } from 'rxjs';
import { projectEvents, ProjectStore } from '@entities/project';
import { UserStore } from '@entities/user';
import { BoardStore } from '@entities/board';
import { ListStore } from '@entities/list';
import { TaskStore } from '@entities/task/model/task.store';
import { CoreStore } from './store/core.store';

@Injectable({
    providedIn: 'root',
})
export class CoreService {
    readonly events = inject(Events);
    readonly dispatcher = inject(Dispatcher);

    // Init stores
    readonly userStore = inject(UserStore);
    readonly projectStore = inject(ProjectStore);
    readonly boardStore = inject(BoardStore);
    readonly listStore = inject(ListStore);
    readonly taskStore = inject(TaskStore);
    readonly coreStpre = inject(CoreStore);

    constructor() {
        this.events
            .on(sessionEvents.authenticated)
            .pipe(tap(() => this.dispatcher.dispatch(projectEvents.getProjects())))
            .subscribe();
    }
}
