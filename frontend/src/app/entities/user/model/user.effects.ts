import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { actionUserClear, actionUserSet } from './user.actions';
import { tap } from 'rxjs';

@Injectable()
export class UserEffects {
    private actions$ = inject(Actions);

    set_user$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionUserSet),
            // dispatch event to reconnect websocket for example
            tap((action) => console.log(`user ${action.user.full_name} added to store`)),
        ),
        { dispatch: false },
    );

    clear_user$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionUserClear),
            // dispatch event to close websocket connection for example
            tap(() => console.log('user store cleaned')),
        ),
        { dispatch: false },
    );
}
