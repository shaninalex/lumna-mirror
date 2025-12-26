import { Component, inject } from '@angular/core';
import { ActivatedRoute, RouterLink } from "@angular/router";
import { ProjectStore } from '@entities/project';

@Component({
    selector: 'app-project-detail',
    imports: [RouterLink],
    template: `
        <h1>Project Detail {{ projectId }}</h1>

        <a [routerLink]="['edit']" class="bg-teal-500 text-white rounded-lg px-3 py-1 cursor-pointer inline-block">Edit</a>
    `,
})
export class ProjectDetail {
    projectId: number
    private activatedRoute = inject(ActivatedRoute);

    private readonly projectStore = inject(ProjectStore);

    constructor() {
        // Access route parameters
        this.activatedRoute.params.subscribe((params) => {
            try {
                const id = parseInt(params['id'])
                this.projectId = id;
            } catch {
                alert("Invalid project id")
            }
        });
    }
}
