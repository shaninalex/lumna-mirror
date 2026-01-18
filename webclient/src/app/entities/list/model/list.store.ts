import { patchState, signalStore, withComputed, withHooks, withMethods } from '@ngrx/signals';
import {
    addEntities,
    addEntity,
    removeEntity,
    updateEntities,
    updateEntity,
    withEntities,
} from '@ngrx/signals/entities';
import { Events, withEventHandlers } from '@ngrx/signals/events';
import { inject } from '@angular/core';
import { map, switchMap, tap } from 'rxjs';
import { mapResponse } from '@ngrx/operators';
import { ListModel } from './list.model';
import { ListApi } from '../api/list.api';
import { listEvents } from './list.events';

export const ListStore = signalStore(
    { providedIn: 'root' },
    withEntities<ListModel>(),
    withHooks({
        onInit() {},
    }),
    withMethods((store) => ({
        boardLists(boardId: string): ListModel[] {
            return store
                .entities()
                .filter((p) => p.board_id === boardId)
                .sort((a, b) => a.order - b.order);
        },
    })),
    withEventHandlers((store, events = inject(Events), listApi = inject(ListApi)) => ({
        actionGetBoardLists$: events.on(listEvents.getLists).pipe(
            switchMap((e) =>
                listApi.List(e.payload).pipe(
                    mapResponse({
                        next: (lists) => listEvents.setLists(lists),
                        error: (error) => listEvents.failed(error),
                    }),
                ),
            ),
        ),

        actionCreateList$: events.on(listEvents.create).pipe(
            switchMap((e) =>
                listApi.Create(e.payload.boardId, e.payload.data).pipe(
                    mapResponse({
                        next: (list) => listEvents.setList(list),
                        error: (error) => listEvents.failed(error),
                    }),
                ),
            ),
        ),

        actionPatchList$: events.on(listEvents.patch).pipe(
            switchMap((e) =>
                listApi.Patch(e.payload.listId, e.payload.data).pipe(
                    mapResponse({
                        next: (list) => listEvents._patchSuccess(list),
                        error: (error) => listEvents.failed(error),
                    }),
                ),
            ),
        ),

        actionDeleteList$: events.on(listEvents.delete).pipe(
            switchMap((e) =>
                listApi.Delete(e.payload).pipe(
                    mapResponse({
                        next: () => listEvents._deleteSuccess(e.payload),
                        error: (error) => listEvents.failed(error),
                    }),
                ),
            ),
        ),

        changeOrder$: events.on(listEvents.changeOrder).pipe(
            map((data) =>
                patchState(
                    store,
                    updateEntities({
                        ids: data.payload.lists.map((l) => l.id),
                        changes: (list) => {
                            const lists = data.payload.lists;
                            const l = lists.find((l) => l.id === list.id);
                            if (!l) return {};
                            return {
                                order: l.order,
                            };
                        },
                    }),
                ),
            ),
        ),

        _deleteList$: events
            .on(listEvents._deleteSuccess)
            .pipe(tap((e) => patchState(store, removeEntity(e.payload)))),

        _patchSuccess$: events.on(listEvents._patchSuccess).pipe(
            tap((e) =>
                patchState(
                    store,
                    updateEntity({
                        id: e.payload.id,
                        changes: e.payload,
                    }),
                ),
            ),
        ),

        _setList$: events
            .on(listEvents.setList)
            .pipe(tap((e) => patchState(store, addEntity(e.payload)))),

        _setLists$: events
            .on(listEvents.setLists)
            .pipe(tap((e) => patchState(store, addEntities(e.payload)))),
    })),
);
