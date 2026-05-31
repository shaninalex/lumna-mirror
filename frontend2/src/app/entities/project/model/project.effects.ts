import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { catchError, exhaustMap, of } from "rxjs";
import {
    actionProjectCreate,
    actionProjectCreateFailed,
    actionProjectCreateSuccessful,
    actionProjectGetList,
    actionProjectSetList
} from "./project.actions";
import { ProjectApi } from "@entities/project/api";
import { switchMap } from "rxjs/operators";
import { HttpErrorResponse } from "@angular/common/http";
import { fromErrorResponse } from "@shared/models";

@Injectable()
export class ProjectEffects {
    private actions$ = inject(Actions);
    private api = inject(ProjectApi);

    project_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProjectGetList),
            exhaustMap(() =>
                this.api
                    .list()
                    .pipe(
                        switchMap((projectList) =>
                            of(actionProjectSetList({ list: projectList }))
                        )
                    )
            )
        )
    );

    project_create$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProjectCreate),
            switchMap((action) =>
                this.api.create(action.data).pipe(
                    switchMap((project) =>
                        of(actionProjectCreateSuccessful({ project }))
                    ),
                    catchError((err: HttpErrorResponse) =>
                        of(
                            actionProjectCreateFailed({
                                errors: fromErrorResponse(err)
                            })
                        )
                    )
                )
            )
        )
    );
}
