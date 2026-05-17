import { inject } from "@angular/core";
import { CanActivateFn, CanMatchFn, Router } from "@angular/router";
import { Store } from "@ngrx/store";
import { filter, map, switchMap, take, tap } from "rxjs/operators";
import { actionUserGet } from "@entities/user";

type SessionShape = {
    session: { authenticated: boolean; checked: boolean };
};

const checkAuth = () => {
    const store = inject(Store<SessionShape>);
    const router = inject(Router);
    const session$ = store.select((state) => state.session);

    return session$.pipe(
        take(1),
        tap((s) => {
            if (!s.checked) {
                store.dispatch(actionUserGet());
            }
        }),
        switchMap(() =>
            session$.pipe(
                filter((s) => s.checked),
                take(1)
            )
        ),
        map(
            ({ authenticated }) =>
                authenticated || router.createUrlTree(["/auth/login"])
        )
    );
};

export const authCanActivate: CanActivateFn & CanMatchFn = () => checkAuth();
