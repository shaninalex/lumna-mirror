import { createFeature } from "@ngrx/store";
import { statusReducer } from "./model";

export const statusFeature = createFeature({
    name: 'status',
    reducer: statusReducer,
});
