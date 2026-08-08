import { Component } from '@angular/core';
import { GlobalLayout } from '@core/layout';
import { WorkspaceCreateFeature } from "@features/workspace";

@Component({
    selector: 'lu-workspace-create-page',
    imports: [GlobalLayout, WorkspaceCreateFeature],
    templateUrl: './workspace-create.component.html',
})
export class WorkspaceCreateComponent {}
