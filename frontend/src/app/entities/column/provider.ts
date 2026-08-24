import { createFeature } from "@ngrx/store";
import { statusReducer } from "./model";

export const columnFeature = createFeature({
    name: 'column',
    reducer: statusReducer,
});
