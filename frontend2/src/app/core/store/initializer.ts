import { inject } from "@angular/core";
import { Store } from "@ngrx/store";
import { actionApplicationInit } from "./app";

export const ApplicationInit = () => {
    inject(Store).dispatch(actionApplicationInit());
};
