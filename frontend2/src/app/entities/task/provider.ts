import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { TaskEffects, taskReducer } from "./model";

export function provideTaskFeature() {
    console.log("[PROVIDE] task");
    return [provideState("task", taskReducer), provideEffects(TaskEffects)];
}
