import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { ProjectEffects, projectReducer } from "./model";
import { ProjectApi } from "@entities/project/api";
import { isDevMode } from "@angular/core";

export function provideProjectFeature() {
    if (isDevMode()) {
        console.info("[PROVIDE] project");
    }

    return [
        provideState("project", projectReducer),
        provideEffects(ProjectEffects),

        ProjectApi
    ];
}
