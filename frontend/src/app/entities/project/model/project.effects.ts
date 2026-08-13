import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import {
    actionProjectCreate, actionProjectCreateFailed, actionProjectDelete, actionProjectDeleteSuccess,
    actionProjectList,
    actionProjectSetCurrent,
    actionProjectsSetList,
    actionProjectUpdate,
    actionProjectUpsert,
} from './project.actions';
import { catchError, exhaustMap, of, switchMap, tap } from 'rxjs';
import { ProjectApi } from '../api/project.service';
import { HttpErrorResponse } from '@angular/common/http';
import { fromErrorResponse } from '@shared/models';

@Injectable()
export class ProjectEffects {
    private actions$ = inject(Actions);
    private projectsApi = inject(ProjectApi);

    get_projects_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProjectList),
            exhaustMap((action) =>
                this.projectsApi.GetProjects(action.workspace_id).pipe(
                    switchMap((data) => of(actionProjectsSetList({ projects: data })))
                ),
            ),
        ),
    );

    create_project$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProjectCreate),
            exhaustMap((action) =>
                this.projectsApi.CreateProject(action.payload).pipe(
                    switchMap((data) => of(actionProjectUpsert({ project: data }))),
                    catchError((err: HttpErrorResponse) => of(actionProjectCreateFailed({ errors: fromErrorResponse(err) })))
                ),
            ),
        ),
    );

    update_project$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProjectUpdate),
            exhaustMap((action) =>
                this.projectsApi.Patch(action.id, action.data).pipe(
                    switchMap((data) => of(actionProjectUpsert({ project: data })))
                ),
            ),
        ),
    );

    delete_project$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProjectDelete),
            exhaustMap((action) =>
                this.projectsApi.DeleteProject(action.project_id).pipe(
                    switchMap(() => of(actionProjectDeleteSuccess({ project_id: action.project_id }))))
                )
        ),
    );

    set_current_project$ = createEffect(() => 
        this.actions$.pipe(
            ofType(actionProjectSetCurrent),
            tap((action) => localStorage.setItem("last_project_id", String(action.id)))
        ),
        { dispatch: false }
    );
}
