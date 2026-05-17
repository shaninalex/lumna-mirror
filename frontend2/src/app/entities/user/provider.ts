import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { UserApi } from "./api";
import { UserEffects, userReducer } from "./model";

export function provideUserFeature() {
    console.log("[GLOBAL PROVIDE] user");
    return [
        provideState("user", userReducer),
        provideEffects(UserEffects),

        UserApi
    ];
}
