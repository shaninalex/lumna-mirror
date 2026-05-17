import { inject } from "@angular/core";
import { CanActivateFn, CanMatchFn, Router } from "@angular/router";
import { Store } from "@ngrx/store";
import { filter, map, take } from "rxjs/operators";

type SessionShape = {
    session: { authenticated: boolean; checked: boolean };
};

const checkAuth = () => {
    const store = inject(Store<SessionShape>);
    const router = inject(Router);

    return store.select((state) => state.session).pipe(
        filter((s) => s.checked),
        take(1),
        map(({ authenticated }) => authenticated || router.createUrlTree(["/auth/login"]))
    );
};

export const authCanActivate: CanActivateFn & CanMatchFn = () => checkAuth();
