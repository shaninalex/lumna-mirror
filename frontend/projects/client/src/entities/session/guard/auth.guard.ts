import {CanMatchFn} from '@angular/router';
import {inject} from '@angular/core';
import {Store} from '@ngrx/store';
import {GetSessionAction, selectSession} from '../model';
import {AppState} from '@client/shared/store';
import {filter, map, take} from "rxjs"


export const AuthGuard: CanMatchFn = () => {
    const store = inject(Store<AppState>);

    return store.select(selectSession).pipe(
        filter(session => session !== undefined), // wait until effect resolves
        take(1),
        map(session => {
            if (!session) return false;
            const expDate = new Date(session.expires_at as string);
            return expDate > new Date();
        })
    );
};

// export const AuthGuard: CanMatchFn = () => {
//     const store = inject(Store<AppState>);
//
//     return store.select(selectSession).pipe(
//         take(1), // only need one emission
//         map(session => {
//             if (!session) {
//                 store.dispatch(GetSessionAction())
//                 return false;
//             }
//
//             const expDate = new Date(session.expires_at as string);
//             const currentDate = new Date();
//
//             if (expDate < currentDate) {
//                 store.dispatch(GetSessionAction())
//                 return false;
//             }
//
//             return true;
//         })
//     );
// };
