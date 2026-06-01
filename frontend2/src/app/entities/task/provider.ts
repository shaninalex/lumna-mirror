import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { TaskEffects, taskReducer } from "./model";
import { TaskApi } from "@entities/task/api";

export function provideTaskFeature() {
    console.log("[PROVIDE] task");
    return [
        provideState("task", taskReducer),
        provideEffects(TaskEffects),

        TaskApi
    ];
}
