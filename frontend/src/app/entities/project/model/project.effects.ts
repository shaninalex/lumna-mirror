import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import {
    actionProjectCreate,
    actionProjectList,
    actionProjectAdd,
    actionProjectsSetList,
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
            ofType(actionProjectList.type),
            exhaustMap(() =>
                this.projectsApi
                    .GetProjects()
                    .pipe(switchMap((data) => of(actionProjectsSetList({ projects: data })))),
            ),
        ),
    );

    create_project$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProjectCreate.type),
            exhaustMap((action) =>
                this.projectsApi
                    .CreateProject(action.payload)
                    .pipe(switchMap((data) => of(actionProjectAdd({ project: data })))),
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
