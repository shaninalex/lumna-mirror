import {CanActivateFn, CanMatchFn, Router} from '@angular/router';
import {inject} from '@angular/core';
import {userEvents, UserStore} from '@entities/user';
import {Dispatcher} from '@ngrx/signals/events';
import {SessionStore} from '@core/store/session.store';
import {toObservable} from '@angular/core/rxjs-interop';
import {filter, map, take, tap} from 'rxjs';
import {appEvents} from '@core/store/app.events';
import {projectEvents, ProjectStore} from '@entities/project';
import {ApplicationStore} from '@core/store/app.store';

// export const authGuard: CanActivateFn = (route, state) => {
//     const userStore = inject(UserStore);
//     const user = userStore.user;
//     if (!user()) {
//         const dispatcher = inject(Dispatcher)
//         dispatcher.dispatch(userEvents.getUser());
//     }
//     return true;
// };


export const authGuard: CanMatchFn = () => {
    const dispatcher = inject(Dispatcher);
    const sessionStore = inject(SessionStore);
    const _appStore  = inject(ApplicationStore)
    const _projectStore = inject(ProjectStore)
    const router = inject(Router);

    if (sessionStore.status() === 'idle') {
        dispatcher.dispatch(userEvents.getUser());
        console.log("idle")
    }

    return toObservable(sessionStore.status).pipe(
        filter(status => status !== 'loading' && status !== 'idle'),
        take(1),
        tap(status => {
            console.log(status)
            if (status === 'unauthenticated') {
                router.navigate(['/auth/login']);
            }
        }),
        filter(status => status === 'authenticated'),
        tap(() => {
            dispatcher.dispatch(appEvents.applicationReady(true))
            dispatcher.dispatch(projectEvents.getProjects())
        }),
        // map to boolean implicitly true
        map(status => status === "authenticated")
    );
};
