import { Component, inject, Input, OnInit, signal } from '@angular/core';
import { ProjectModel, ProjectStore } from '@entities/project';

@Component({
    selector: 'app-project-edit-feature',
    imports: [],
    template: `
        <p>project-edit works! {{ project()?.name }}</p>
    `,
})
export class ProjectEditFeature implements OnInit {
    @Input() projectId: number;
    project = signal<ProjectModel|undefined>(undefined);
    private readonly store = inject(ProjectStore);

    ngOnInit(): void {
        const p = this.store.byId(this.projectId);
        this.project.set(p);
    }
}
