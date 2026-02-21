import { ActivatedRouteSnapshot, ResolveFn } from '@angular/router';
import { Store } from '@ngrx/store';
import { inject } from '@angular/core';
import { filter, tap } from 'rxjs';
import { actionTaskGetTaskById, selectTaskById, TaskModel, TaskState } from '@entities/task';

export const taskResolver: ResolveFn<TaskModel | null> = (route: ActivatedRouteSnapshot) => {
    const store = inject(Store<TaskState>);
    const taskId = Number(route.paramMap.get('id'));

    return store.select(selectTaskById(taskId)).pipe(
        tap((task: TaskModel | null) => {
            if (!task)
                store.dispatch(actionTaskGetTaskById({ task_id: taskId }));
        }),
        filter((task) => !!task)
    );
};
