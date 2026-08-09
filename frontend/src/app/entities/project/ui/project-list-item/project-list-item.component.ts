import { Component, Input } from '@angular/core';
import { ProjectModel } from '@entities/project/model';

@Component({
    selector: 'lu-project-list-item',
    imports: [],
    template: `
        <button cdkMenuItem type="button" class="list-group-item d-flex justify-content-between align-items-center list-group-item-action gap-2">
            <div class="me-auto">
                {{ project.title }}
            </div>
            <span class="badge text-bg-primary rounded-pill">14</span>
        </button>
    `,
})
export class ProjectListItemComponent {
    @Input() project: ProjectModel;

    // select current open ( active ) tasks for this project
}
