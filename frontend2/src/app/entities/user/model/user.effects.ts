import { Actions, createEffect, ofType } from "@ngrx/effects";
import { inject, Injectable } from "@angular/core";
import { actionUserClear, actionUserGet, actionUserSet } from "./user.actions";
import { catchError, exhaustMap, map, of, switchMap, tap } from "rxjs";
import { UserApi } from "../api/user.service";
import { actionSessionFailed } from "@core/store/session/session.actions";
import { Error } from "@shared/models";

@Injectable()
export class UserEffects {
    private actions$ = inject(Actions);
    private api = inject(UserApi);

    get_user$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionUserGet),
            exhaustMap(() =>
                this.api.me().pipe(
                    map((user) => actionUserSet({ user })),
                    catchError(() => of(actionSessionFailed()))
                )
            )
        )
    );

    set_user$ = createEffect(
        () =>
            this.actions$.pipe(
                ofType(actionUserSet),
                tap((action) =>
                    console.log(`user ${action.user.full_name} added to store`)
                )
            ),
        { dispatch: false }
    );

    clear_user$ = createEffect(
        () =>
            this.actions$.pipe(
                ofType(actionUserClear),
                tap(() => console.log("user store cleaned"))
            ),
        { dispatch: false }
    );
}
