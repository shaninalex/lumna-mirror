import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { LocalStorageService } from '@shared/services';

export const lastRouteRedirect: CanActivateFn = () => {
    const router = inject(Router);
    const localStorageService = inject(LocalStorageService);
    const url = localStorageService.get('last_url');
    if (!url) {
        return router.createUrlTree(['app', 'workspaces']);
    }
    return router.createUrlTree([url]);
};
