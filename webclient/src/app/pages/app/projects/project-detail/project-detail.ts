import {Component, computed, inject} from '@angular/core';
import { ActivatedRoute, RouterLink } from "@angular/router";
import { ProjectStore } from '@entities/project';


@Component({
    selector: 'app-project-detail',
    imports: [RouterLink],
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

        <div class="bg-gray-200 card">
            <h2 class="font-medium text-lg mb-4">Boards</h2>
            <nav class="flex flex-wrap gap-4">
                <button class="btn btn-primary">Development</button>
                <button class="btn btn-primary">Design</button>
                <button class="btn btn-primary">Marketing/SEO</button>
            </nav>
        </div>

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
    private activatedRoute = inject(ActivatedRoute);

    private readonly store = inject(ProjectStore);
    project = computed(() => this.store.entities().find(p => p.id === this.projectId));

    constructor() {
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
