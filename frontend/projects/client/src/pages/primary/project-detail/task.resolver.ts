import {ResolveFn} from '@angular/router';
import {inject} from '@angular/core';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {filter, take} from 'rxjs';
import {selectTask, Task} from '@client/entities/task';

export const taskResolver: ResolveFn<Task | undefined> = (route) => {
    const store: Store<AppState> = inject(Store<AppState>)
    return store.select(selectTask(route.params["taskCode"])).pipe(
        filter((task): task is Task => !!task),
        take(1),
        // tap(task => {
        //     if (task) {
        //         // Do something here
        //     }
        // })
    );
};
