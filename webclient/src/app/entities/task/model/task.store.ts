import { inject } from '@angular/core'
import {patchState, signalStore, withMethods} from "@ngrx/signals";
import {addEntities, addEntity, updateEntities, withEntities} from "@ngrx/signals/entities";
import { TaskModel } from "./task.model";
import { Events, withEventHandlers } from "@ngrx/signals/events";
import { taskEvents } from './task.events';
import {TaskApi} from '@entities/task/api/task.api';
import {map, switchMap, tap} from 'rxjs';
import {mapResponse} from '@ngrx/operators';
import {ListModel} from '@entities/list';

export const TaskStore = signalStore(
    { providedIn: 'root' },
    withEntities<TaskModel>(),
    withMethods((store) => ({
        boardTasks(boardId: number): TaskModel[] {
            return store.entities().filter(p => p.board_id === boardId)
        }
    })),
    withEventHandlers((
        store,
        events = inject(Events),
        api = inject(TaskApi)
    ) => ({
        actionGetTasks$: events
            .on(taskEvents.getTasks)
            .pipe(
                tap(() => console.log(taskEvents.getTasks.type)),
                switchMap(e =>
                    api.List(e.payload.board_id).pipe(
                        mapResponse({
                            next: tasks => taskEvents.setTasks(tasks),
                            error: error => taskEvents.failed(error)
                        })
                    )
                )
            ),

        actionCreate$: events
            .on(taskEvents.create)
            .pipe(
                tap(() => console.log(taskEvents.getTasks.type)),
                switchMap(e =>
                    api.Create(e.payload.board_id, e.payload.data).pipe(
                        mapResponse({
                            next: task => taskEvents.setTask(task),
                            error: error => taskEvents.failed(error)
                        })
                    )
                )
            ),

        actionChangeOrder$: events
            .on(taskEvents.changeOrder)
            .pipe(
                tap(() => console.log(taskEvents.changeOrder.type)),
                map(data => {
                    // same list
                    if (data.payload.listId && data.payload.tasks) {
                        return patchState(store, updateEntities({
                            ids: data.payload.tasks.map(t => t.id),
                            changes: (task) => {
                                const tasks = data.payload.tasks
                                if (!tasks) return {}
                                const t = tasks.find(t => t.id === task.id)
                                if (!t) return {}
                                return {
                                    order: t.order
                                }
                            }
                        }))
                    } else {

                    }
                })
            ),


        _setTasks$: events
            .on(taskEvents.setTasks)
            .pipe(
                tap(e => patchState(store, addEntities(e.payload ? e.payload : [])))
            ),

        _setTask$: events
            .on(taskEvents.setTask)
            .pipe(
                tap(e => patchState(store, addEntity(e.payload)))
            )
    }))
)
