import { Actions, createEffect, ofType } from "@ngrx/effects";
import { inject, Injectable } from "@angular/core";
import { actionUserClear, actionUserGet, actionUserSet } from "./user.actions";
import { exhaustMap, of, switchMap, tap } from "rxjs";
import { UserApi } from "../api/user.service";

@Injectable()
export class UserEffects {
    private actions$ = inject(Actions);
    private api = inject(UserApi);

    get_user$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionUserGet),
            exhaustMap(() =>
                this.api
                    .me()
                    .pipe(switchMap((user) => of(actionUserSet({ user }))))
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
