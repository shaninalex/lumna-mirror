import {patchState, signalStore, withHooks, withMethods} from '@ngrx/signals';
import {addEntities, withEntities} from '@ngrx/signals/entities';
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

        _setList$: events
            .on(boardEvents.setList)
            .pipe(
                tap(e => patchState(store, addEntities(e.payload)))
            )
    }))
)
