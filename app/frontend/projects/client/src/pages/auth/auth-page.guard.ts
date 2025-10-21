import {CanActivateFn, Router} from '@angular/router';
import {inject} from '@angular/core';
import {AppState} from '@client/shared/store';
import {Store} from '@ngrx/store';
import {selectUser} from '@client/entities/user';

export const authPageGuard: CanActivateFn = () => {
    // const store = inject(Store<AppState>);
    // const router = inject(Router);
    //
    // return router.parseUrl('/');
    return true
};
