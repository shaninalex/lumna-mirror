import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { SprintEffects } from "@entities/sprint/model/sprint.effects";
import { SprintApi } from "@entities/sprint/api";
import { sprintReducer } from "@entities/sprint/model";
import { isDevMode } from "@angular/core";

export function provideSprintFeature() {
    if (isDevMode()) {
        console.info("[PROVIDE] sprints");
    }
    return [
        provideState("sprint", sprintReducer),
        provideEffects(SprintEffects),

        SprintApi
    ];
}
