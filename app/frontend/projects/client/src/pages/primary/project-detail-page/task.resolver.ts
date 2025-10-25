import { ResolveFn } from '@angular/router'
import { inject } from '@angular/core'
import { Store } from '@ngrx/store'
import { AppState } from '@client/shared/store'
import { filter, switchMap, take, tap } from 'rxjs'
import { selectTask, Task, TaskService, TaskUpdateAction } from '@client/entities/task'

export const taskResolver: ResolveFn<Task | undefined> = route => {
    const store: Store<AppState> = inject(Store<AppState>)
    const api = inject(TaskService)

    return store.select(selectTask(route.params['taskCode'])).pipe(
        filter((task): task is Task => !!task),
        take(1),
        switchMap(task =>
            api.Detail(task.id).pipe(
                tap(task => {
                    store.dispatch(TaskUpdateAction({ payload: task }))
                })
            )
        )
    )
}
