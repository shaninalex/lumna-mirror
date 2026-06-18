import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { ApplicationEffects } from "@core/store";
import { appReducer } from "./app.store";
import { isDevMode } from "@angular/core";

export function provideApplicationFeature() {
    if (isDevMode()) {
        console.info("[GLOBAL PROVIDE] application");
    }

    return [
        provideState("application", appReducer),
        provideEffects(ApplicationEffects)
    ];
}
