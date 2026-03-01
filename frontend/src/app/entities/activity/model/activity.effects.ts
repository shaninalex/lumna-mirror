import { inject, Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { ActivityApi } from '../api/activity.api';
import {
    actionActivityGetList, actionActivitySetList,
} from './activity.actions';
import { exhaustMap, of, switchMap } from 'rxjs';

@Injectable()
export class ActivityEffects {
    private actions$ = inject(Actions);
    private taskApi = inject(ActivityApi);

    activity_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionActivityGetList),
            exhaustMap((action) =>
                this.taskApi
                    .List(action.entity_id, action.entity_type)
                    .pipe(switchMap((list) => of(actionActivitySetList({ list })))),
            ),
        ),
    );
}
