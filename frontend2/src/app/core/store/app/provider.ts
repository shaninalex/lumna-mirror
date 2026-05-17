import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { ApplicationEffects } from "./app.effects";
import { appReducer } from "./app.store";

export function provideApplicationFeature() {
    console.log("[GLOBAL PROVIDE] application");
    return [
        provideState("application", appReducer),
        provideEffects(ApplicationEffects)
    ];
}
