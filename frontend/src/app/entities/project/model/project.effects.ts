import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import {
    actionProjectCreate,
    actionProjectList,
    actionProjectsSetList,
    actionProjectUpdate,
    actionProjectUpsert,
} from './project.actions';
import { exhaustMap, map, of, switchMap, tap } from 'rxjs';
import { ProjectApi } from '../api/project.service';
import { actionBoardSetList, BoardModel } from '@entities/board';

@Injectable()
export class ProjectEffects {
    private actions$ = inject(Actions);
    private projectsApi = inject(ProjectApi);

    get_projects_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProjectList),
            exhaustMap(() =>
                this.projectsApi
                    .GetProjects()
                    .pipe(switchMap((data) => of(actionProjectsSetList({ projects: data })))),
            ),
        ),
    );

    create_project$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProjectCreate),
            exhaustMap((action) =>
                this.projectsApi
                    .CreateProject(action.payload)
                    .pipe(switchMap((data) => of(actionProjectUpsert({ project: data })))),
            ),
        ),
    );

    update_project$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProjectUpdate),
            exhaustMap((action) =>
                this.projectsApi
                    .Patch(action.id, action.data)
                    // TODO: upsert instead of add!
                    .pipe(switchMap((data) => of(actionProjectUpsert({ project: data })))),
            ),
        ),
    );

    set_boards$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProjectsSetList),
            map((data) => {
                const boards: BoardModel[] = [];
                for (let i = 0; i < data.projects.length; i++) {
                    for (let bi = 0; bi < data.projects[i].boards.length; bi++) {
                        boards.push(data.projects[i].boards[bi]);
                    }
                }
                return actionBoardSetList({ boards });
            }),
        ),
    );
}
