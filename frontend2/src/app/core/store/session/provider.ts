import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { sessionReducer } from "./session.store";
import { SessionEffects } from "./session.effects";
import { SessionApi } from "./session.api";

export function provideSessionFeature() {
    console.log("[PROVIDE] session");
    return [
        provideState("session", sessionReducer),
        provideEffects(SessionEffects),

        SessionApi
    ];
}
