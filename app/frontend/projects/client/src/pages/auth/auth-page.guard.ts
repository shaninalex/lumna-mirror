import { CanActivateFn } from '@angular/router'

export const authPageGuard: CanActivateFn = () => {
    // const store = inject(Store<AppState>);
    // const router = inject(Router);
    //
    // return router.parseUrl('/');
    return true
}
