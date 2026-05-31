import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { ProjectEffects, projectReducer } from "./model";
import { ProjectApi } from "@entities/project/api";

export function provideProjectFeature() {
    console.log("[PROVIDE] project");
    return [
        provideState("project", projectReducer),
        provideEffects(ProjectEffects),

        ProjectApi
    ];
}
