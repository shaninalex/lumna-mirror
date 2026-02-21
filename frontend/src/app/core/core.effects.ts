import { inject, Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { routerNavigatedAction } from '@ngrx/router-store';
import { UiService } from '@shared/ui';
import { tap } from 'rxjs/operators';

@Injectable()
export class RouterEffects {
    private actions$ = inject(Actions);
    private ui = inject(UiService);

    routeChanged$ = createEffect(
        () =>
            this.actions$.pipe(
                ofType(routerNavigatedAction),
                tap((action) => {
                    this.ui.setUrl(action.payload.routerState.url)
                }),
            ),
        { dispatch: false },
    );
}
