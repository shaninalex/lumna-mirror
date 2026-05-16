import { createFeatureSelector, createSelector } from "@ngrx/store";
import { AppState } from "./app.store";

const feature = createFeatureSelector<AppState>("application");

export const selectAppState = createSelector(feature, (state) => {
    return state;
});

export const selectSidebarState = createSelector(
    selectAppState,
    (state) => state.sidebar
);
