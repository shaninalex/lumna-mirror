import { createFeature } from "@ngrx/store";
import { taskReducer } from "./model";

export const taskFeature = createFeature({
    name: 'task',
    reducer: taskReducer,
});
