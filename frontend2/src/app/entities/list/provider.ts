import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { ListEffects, listReducer } from "./model";

export function provideListFeature() {
    console.log("[PROVIDE] list");
    return [provideState("list", listReducer), provideEffects(ListEffects)];
}
