import { ActivatedRouteSnapshot, ResolveFn, RouterStateSnapshot } from '@angular/router';
import { Store } from '@ngrx/store';
import { actionBoardGet, BoardModel, BoardState } from '@entities/board';
import { inject } from '@angular/core';
import { filter, tap } from 'rxjs';
import { selectBoardById } from '@entities/board/model/board.selectors';

export const boardResolver: ResolveFn<BoardModel | null> = (
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot,
) => {
    const store = inject(Store<BoardState>);
    const boardId = Number(route.paramMap.get('id'));

    return store.select(selectBoardById(boardId)).pipe(
        tap((board) => {
            if (!board) store.dispatch(actionBoardGet({ boardId }));
        }),
        filter((board) => !!board),
    );
};
