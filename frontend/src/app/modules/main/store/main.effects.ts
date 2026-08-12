import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { map, tap } from 'rxjs';
import { ROUTER_NAVIGATED, RouterNavigatedPayload } from '@ngrx/router-store';
import { LastUrlService } from '@root/src/app/modules/main/services';

@Injectable()
export class AppEffects {
    private actions$ = inject(Actions);
    private lastUrlService = inject(LastUrlService);

    routeChanged$ = createEffect(
        () =>
            this.actions$.pipe(
                ofType(ROUTER_NAVIGATED),
                map((action) => action.payload),
                tap((payload: RouterNavigatedPayload) => this.lastUrlService.set(payload.event.url)),
            ),
        { dispatch: false },
    );
}
