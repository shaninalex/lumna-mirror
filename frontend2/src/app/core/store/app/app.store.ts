import { createReducer, on } from "@ngrx/store";
import { actionApplicationSidebarToggle } from "./app.actions";

export type AppState = {
    sidebar: boolean;
};

const initialState: AppState = {
    sidebar: true
};

export const appReducer = createReducer(
    initialState,
    on(actionApplicationSidebarToggle, (state, action) => ({
        sidebar: !state.sidebar
    }))
);
