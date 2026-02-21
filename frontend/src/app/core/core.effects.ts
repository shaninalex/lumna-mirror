import { inject, Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { routerNavigatedAction } from '@ngrx/router-store';

@Injectable()
export class RouterEffects {
    private actions$ = inject(Actions);

    routeChanged$ = createEffect(
        () =>
            this.actions$.pipe(
                ofType(routerNavigatedAction),
                // tap((action) => this.drillRouteState(action.payload.routerState)),
            ),
        { dispatch: false },
    );
    //
    // private drillRouteState(state: SerializedRouterStateSnapshot) {
    //     // console.clear()
    //     this.walk(state.root);
    // }
    //
    // private walk(route: any) {
    //     console.log('------');
    //     console.log('Route:', route.routeConfig?.path);
    //     console.log('Params:', route.params);
    //     console.log('Data:', route.data);
    //
    //     if (route.firstChild) {
    //         this.walk(route.firstChild);
    //     }
    // }
}
