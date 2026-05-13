import { createFeatureSelector, createSelector } from "@ngrx/store";
import { listAdapter, ListState } from "./list.store";
import { ListModel } from "./list.model";

const feature = createFeatureSelector<ListState>("list");
const selectors = listAdapter.getSelectors();
export const selectListList = createSelector(feature, (state) =>
    selectors.selectAll(state)
);

export const selectList = (id: number) =>
    createSelector(selectListList, (list) =>
        list.filter((a: ListModel) => a.id === id)
    );
