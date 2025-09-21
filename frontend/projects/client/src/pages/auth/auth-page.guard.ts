import {CanActivateFn, Router} from '@angular/router';
import {inject} from '@angular/core';
import {TokenService} from '@client/shared/common';

export const authPageGuard: CanActivateFn = () => {
    const tokenService = inject(TokenService);
    const router = inject(Router);

    if (tokenService.getAuthToken() === "") {
        // no token - allow access to auth pages
        return true;
    }

    return router.parseUrl('/');
};
