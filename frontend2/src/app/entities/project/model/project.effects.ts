import { inject, Injectable } from "@angular/core";
import { Actions, createEffect, ofType } from "@ngrx/effects";
import { exhaustMap, of } from "rxjs";
import { actionProjectGetList, actionProjectSetList } from "./project.actions";
import { ProjectApi } from "@entities/project/api";
import { switchMap } from "rxjs/operators";

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
}
