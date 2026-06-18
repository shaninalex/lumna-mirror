import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { SessionApi, SessionEffects, sessionReducer } from "@core/store";
import { isDevMode } from "@angular/core";

export function provideSessionFeature() {
    if (isDevMode()) {
        console.info("[PROVIDE] session");
    }
    return [
        provideState("session", sessionReducer),
        provideEffects(SessionEffects),

        SessionApi
    ];
}
