import { Component, inject, Input } from '@angular/core';
import { ProjectModel } from '@entities/project/model';
import { selectCurrentWorkspaceId } from '@entities/workspace';
import { Store } from '@ngrx/store';
import { RouterLink } from "@angular/router";

@Component({
    selector: 'lu-project-list-item',
    imports: [RouterLink],
    template: `
        <a cdkMenuItem 
            [routerLink]="['/app/w', currentWorkspaceId() || '', 'p', project.id]"
            class="list-group-item d-flex justify-content-between align-items-center list-group-item-action gap-2">
            <div class="me-auto">
                {{ project.title }}
            </div>
            <span class="badge text-bg-primary rounded-pill">14</span>
        </a>
    `,
})
export class ProjectListItemComponent {
    @Input() project: ProjectModel;

    private store = inject(Store);

    currentWorkspaceId = this.store.selectSignal(selectCurrentWorkspaceId);

    // select current open ( active ) tasks for this project
}
