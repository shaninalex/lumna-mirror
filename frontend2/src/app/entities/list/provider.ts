import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { ListEffects, listReducer } from "./model";
import { isDevMode } from "@angular/core";

export function provideListFeature() {
    if (isDevMode()) {
        console.info("[PROVIDE] list");
    }

    return [provideState("list", listReducer), provideEffects(ListEffects)];
}
