import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { UserApi } from "./api";
import { UserEffects, userReducer } from "./model";
import { isDevMode } from "@angular/core";

export function provideUserFeature() {
    if (isDevMode()) {
        console.info("[GLOBAL PROVIDE] user");
    }

    return [
        provideState("user", userReducer),
        provideEffects(UserEffects),

        UserApi
    ];
}
