import {patchState, signalStore, withHooks, withMethods} from '@ngrx/signals';
import {addEntities, addEntity, removeEntity, withEntities} from '@ngrx/signals/entities';
import {Events, withEventHandlers} from '@ngrx/signals/events';
import {inject} from '@angular/core';
import {BoardApi} from '@entities/board/api/board.api';
import {boardEvents} from '@entities/board/model/board.events';
import {switchMap, tap} from 'rxjs';
import {mapResponse} from '@ngrx/operators';
import {BoardModel} from '@entities/board';

export const BoardStore = signalStore(
    { providedIn: 'root' },
    withEntities<BoardModel>(),
    withHooks({
        onInit() {
            console.log('BoardStore initialized');
        },
    }),
    withMethods((store) => ({
        projectBoards(projectId: number): BoardModel[] {
            console.log(store.entities())
            return store.entities().filter(p => p.project_id === projectId)
        }
    })),
    withEventHandlers((
        store,
        events = inject(Events),
        boardApi = inject(BoardApi),
    ) => ({
        actionGetProjectBoards$: events
            .on(boardEvents.getList)
            .pipe(
                tap(() => console.log(boardEvents.getList.type)),
                switchMap(e =>
                    boardApi.List(e.payload).pipe(
                        mapResponse({
                            next: boards => boardEvents.setList(boards),
                            error: error => boardEvents.failed(error)
                        })
                    )
                )
            ),

        actionCreateBoard$: events
            .on(boardEvents.create)
            .pipe(
                tap(() => console.log(boardEvents.create.type)),
                switchMap(e =>
                    boardApi.Create(e.payload.project_id, e.payload).pipe(
                        mapResponse({
                            next: board => boardEvents.set(board),
                            error: error => boardEvents.failed(error)
                        })
                    )
                )
            ),

        actionPatchBoard$: events
            .on(boardEvents.patch)
            .pipe(
                tap(() => console.log(boardEvents.patch.type)),
                switchMap(e =>
                    boardApi.Patch(e.payload.boardId, e.payload.data).pipe(
                        mapResponse({
                            next: board => boardEvents.set(board),
                            error: error => boardEvents.failed(error)
                        })
                    )
                )
            ),

        actionDeleteBoard$: events
            .on(boardEvents.delete)
            .pipe(
                tap(() => console.log(boardEvents.delete.type)),
                switchMap(e =>
                    boardApi.Delete(e.payload).pipe(
                        mapResponse({
                            next: () => boardEvents._deleteSuccess(e.payload),
                            error: error => boardEvents.failed(error)
                        })
                    )
                )
            ),

        _deleteBoard$: events
            .on(boardEvents._deleteSuccess)
            .pipe(
                tap(e => patchState(store, removeEntity(e.payload)))
            ),

        _setBoard$: events
            .on(boardEvents.set)
            .pipe(
                tap(e => patchState(store, addEntity(e.payload)))
            ),

        _setList$: events
            .on(boardEvents.setList)
            .pipe(
                tap(e => patchState(store, addEntities(e.payload)))
            )
    }))
)
