import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { map } from 'rxjs';
import { actionSessionAuthenticatedSuccessfull } from './session.actions';
import { actionWorkspaceGetList } from '@entities/workspace';

@Injectable()
export class AppEffects {
    private actions$ = inject(Actions);

    authenticated_successfull$ = createEffect(() => {
        return this.actions$.pipe(
            ofType(actionSessionAuthenticatedSuccessfull),
            map(() => actionWorkspaceGetList()),
        );
    });

}
