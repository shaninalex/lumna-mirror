import { inject, Injectable } from '@angular/core';
import { projectEvents, ProjectStore } from '@entities/project';
import { BoardStore } from '@entities/board';
import { ListStore } from '@entities/list';
import { TaskStore } from '@entities/task/model/task.store';
import { CoreStore } from './store/core.store';
import { sessionEvents, SessionStore } from './store/session.store';
import { Events, Dispatcher } from '@ngrx/signals/events';
import { tap } from 'rxjs';

@Injectable({
    providedIn: 'root',
})
export class CoreService {
    readonly events = inject(Events);
    readonly dispatcher = inject(Dispatcher);

    private readonly _sessionStore = inject(SessionStore);
    private readonly _coreStore = inject(CoreStore);
    private readonly _projectStore = inject(ProjectStore);
    private readonly _boardStore = inject(BoardStore);
    private readonly _listStore = inject(ListStore);
    private readonly _taskStore = inject(TaskStore);

    constructor() {
        this.events
            .on(sessionEvents.authenticated)
            .pipe(
                tap(() => {
                    console.log('CoreService: start app initialization');
                    this.dispatcher.dispatch(projectEvents.getProjects());
                }),
            )
            .subscribe();
    }
}
