import { Component } from '@angular/core';
import { RouterLink } from "@angular/router";
import { GlobalLayout } from '@core/layout';

@Component({
    selector: 'lu-workspaces-page',
    imports: [RouterLink, GlobalLayout],
    templateUrl: 'workspaces.page.html',
})
export class WorkspacesPage {}
