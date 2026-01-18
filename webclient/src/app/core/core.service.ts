import { inject, Injectable } from '@angular/core';
import { ProjectStore } from '@entities/project';
import { BoardStore } from '@entities/board';
import { ListStore } from '@entities/list';
import { TaskStore } from '@entities/task/model/task.store';
import { CoreStore } from './store/core.store';
import { SessionStore } from './store/session.store';

@Injectable({
    providedIn: 'root',
})
export class CoreService {
    private readonly _sessionStore = inject(SessionStore);
    private readonly _coreStore = inject(CoreStore);
    private readonly _projectStore = inject(ProjectStore);
    private readonly _boardStore = inject(BoardStore);
    private readonly _listStore = inject(ListStore);
    private readonly _taskStore = inject(TaskStore);
}
