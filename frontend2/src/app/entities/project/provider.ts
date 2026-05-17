import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { ProjectEffects, projectReducer } from "./model";

export function provideProjectFeature() {
    console.log("[PROVIDE] project");
    return [
        provideState("project", projectReducer),
        provideEffects(ProjectEffects)
    ];
}
