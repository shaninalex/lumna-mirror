import { createFeature } from "@ngrx/store";
import { boardReducer } from "./model";

export const boardFeature = createFeature({
    name: 'board',
    reducer: boardReducer,
});
