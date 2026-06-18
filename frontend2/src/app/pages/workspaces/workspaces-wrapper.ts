import { Component, inject } from "@angular/core";
import { Store } from "@ngrx/store";
import { actionWorkspaceListRequested } from "@entities/workspace";
import { RouterOutlet } from "@angular/router";

@Component({
    selector: "app-workspaces-wrapper",
    imports: [RouterOutlet],
    template: `<router-outlet />`
})
export class WorkspacesWrapper {
    private store = inject(Store);

    constructor() {
        this.store.dispatch(
            actionWorkspaceListRequested({ initiator: "WorkspacesWrapper" })
        );
    }
}
