import { createFeature } from "@ngrx/store";
import { listReducer } from "./model";

export const listFeature = createFeature({
    name: 'list',
    reducer: listReducer,
});
