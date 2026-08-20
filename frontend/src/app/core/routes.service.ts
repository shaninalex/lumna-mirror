import { inject, Injectable } from '@angular/core';
import { selectCurrentProjectId } from '@entities/project';
import { selectCurrentWorkspaceId } from '@entities/workspace';
import { Store } from '@ngrx/store';

@Injectable()
export class AppRoutes {
    private readonly store = inject(Store);

    private readonly workspaceId = this.store.selectSignal(selectCurrentWorkspaceId);

    private readonly projectId = this.store.selectSignal(selectCurrentProjectId);

    private projectRoute(): unknown[] {
        return ['/app', 'w', this.workspaceId(), 'p', this.projectId()];
    }

    projectsCreate(): unknown[] {
        return ['/app', 'w', this.workspaceId(), 'projects', 'create']
    }

    backlog(): unknown[] {
        return [...this.projectRoute(), 'backlog'];
    }

    boards(): unknown[] {
        return [...this.projectRoute(), 'boards'];
    }

    board(id: number): unknown[] {
        return [...this.projectRoute(), 'board', id];
    }

    boardsCreate(): unknown[] {
        return [...this.projectRoute(), 'board', 'create'];
    }

    createTask(): unknown[] {
        return [...this.projectRoute(), 'task', 'create'];
    }

    task(id: number): unknown[] {
        return [...this.projectRoute(), 'task', id];
    }

    editTask(id: string): unknown[] {
        return [...this.projectRoute(), 'task', id, 'edit'];
    }

}
