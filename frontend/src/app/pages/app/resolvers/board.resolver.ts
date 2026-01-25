import {
    ActivatedRouteSnapshot,
    ResolveFn,
    Router,
    RouterStateSnapshot,
    UrlTree,
} from '@angular/router';
import { inject } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import { filter, firstValueFrom, map, tap } from 'rxjs';
import { BoardModel, BoardStore } from '@entities/board';
import { Dispatcher } from '@ngrx/signals/events';
import { listEvents } from '@entities/list/model/list.events';
import { taskEvents } from '@entities/task/model/task.events';

export const boardResolver: ResolveFn<BoardModel | UrlTree> = (
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot,
) => {
    const store = inject(BoardStore);
    const dispatcher = inject(Dispatcher);
    const boardId = route.paramMap.get('boardId')!;
    return firstValueFrom(
        toObservable(store.entities).pipe(
            filter((boards) => boards.some((p) => p.id === boardId)),
            map((boards) => boards.find((p) => p.id === boardId)!),
            tap((board) => {
                dispatcher.dispatch(listEvents.getLists(board.id));
                dispatcher.dispatch(taskEvents.getTasks({ board_id: board.id }));
            }),
        ),
    );
};
