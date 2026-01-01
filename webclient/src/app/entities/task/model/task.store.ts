import { inject } from '@angular/core'
import { signalStore } from "@ngrx/signals";
import { withEntities } from "@ngrx/signals/entities";
import { TaskModel } from "./task.model";
import { Events, withEventHandlers } from "@ngrx/signals/events";
import { taskEvents } from './task.events';

export const TaskStore = signalStore(
    { providedIn: 'root' },
    withEntities<TaskModel>(),
    withEventHandlers((
        store,
        events = inject(Events),
    ) => ({
        actionGetTasks$: events
            .on(taskEvents.getTasks)
    }))
)
