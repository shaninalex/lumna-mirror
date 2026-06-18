import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { TaskEffects, taskReducer } from "./model";
import { TaskApi } from "@entities/task/api";
import { isDevMode } from "@angular/core";

export function provideTaskFeature() {
    if (isDevMode()) {
        console.info("[PROVIDE] task");
    }

    return [
        provideState("task", taskReducer),
        provideEffects(TaskEffects),

        TaskApi
    ];
}
