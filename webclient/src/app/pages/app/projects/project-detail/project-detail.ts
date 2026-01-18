import {Component, computed, inject, signal} from '@angular/core';
import { ActivatedRoute, RouterLink } from "@angular/router";
import {ProjectModel, ProjectStore} from '@entities/project';
import {BoardsList} from '@entities/board/ui/boards-list/boards-list';


@Component({
    selector: 'app-project-detail',
    imports: [RouterLink, BoardsList],
    template: `
        <div class="bg-lime-200 card">
            <nav class="flex flex-wrap gap-4">
                <a href="#" class="hover:underline">Invite members</a>
                <a href="#" class="hover:underline">Roles</a>
                <a href="#" class="hover:underline">Webhooks</a>
                <a href="#" class="hover:underline">Notifications</a>
                <a href="#" class="hover:underline">Emails</a>
                <a href="#" class="hover:underline">Integrations</a>
            </nav>
        </div>

        <app-boards-list [projectId]="projectId" />

        <div class="bg-amber-100 card">
            <a class="font-medium text-lg mb-4 hover:underline">Calendar</a>
        </div>

        <div class="bg-lime-100 card">
            <a class="font-medium text-lg mb-4 hover:underline">Documents/Notes</a>
        </div>

        <a [routerLink]="['edit']" class="btn btn-sm btn-primary">Edit</a>
    `,
})
export class ProjectDetail {
    projectId: number
    private route = inject(ActivatedRoute);
    readonly project = signal<ProjectModel | null>(null);

    constructor() {
        this.route.data.subscribe(data => {
            this.project.set(data['project']);
        });

        this.route.params.subscribe((params) => {
            try {
                const id = parseInt(params['id'])
                this.projectId = id;
            } catch {
                alert("Invalid project id")
            }
        });
    }
}
