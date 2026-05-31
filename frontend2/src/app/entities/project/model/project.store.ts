import { createEntityAdapter, EntityState } from "@ngrx/entity";
import { createReducer, on } from "@ngrx/store";
import {
    actionProjectCreateSuccessful,
    actionProjectSetList
} from "./project.actions";
import { ProjectModel } from "./project.model";

export interface ProjectState extends EntityState<ProjectModel> {}
export const projectAdapter = createEntityAdapter<ProjectModel>();
const initialState = projectAdapter.getInitialState();

export const projectReducer = createReducer(
    initialState,
    on(actionProjectSetList, (state, { list }) =>
        projectAdapter.addMany(list, state)
    ),
    on(actionProjectCreateSuccessful, (state, { project }) =>
        projectAdapter.addOne(project, state)
    )
);
