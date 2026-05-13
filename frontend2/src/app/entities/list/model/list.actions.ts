import { createAction, props } from "@ngrx/store";
import { ListModel } from "./list.model";

export const actionListGetList = createAction("[List] get list");

export const actionListSetList = createAction(
    "[List] set list",
    props<{ list: ListModel[] }>()
);
