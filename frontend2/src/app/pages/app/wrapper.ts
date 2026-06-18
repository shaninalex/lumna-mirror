import { Component, inject } from "@angular/core";
import { RouterOutlet } from "@angular/router";
import { AppLayout } from "@core/layout";
import { Store } from "@ngrx/store";
import { actionWorkspaceListRequested } from "@entities/workspace";
import { actionProjectGetList } from "@entities/project";

@Component({
    selector: "application-wrapper",
    imports: [RouterOutlet, AppLayout],
    template: `
        <app-layout>
            <router-outlet />
        </app-layout>
    `
})
export class ApplicationWrapper {
    private store = inject(Store);

    constructor() {
        this.store.dispatch(
            actionWorkspaceListRequested({ initiator: "Application wrapper" })
        );
        this.store.dispatch(actionProjectGetList());
    }
}
