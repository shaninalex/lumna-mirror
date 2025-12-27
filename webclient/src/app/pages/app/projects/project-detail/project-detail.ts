import {Component, computed, effect, inject} from '@angular/core';
import { ActivatedRoute, RouterLink } from "@angular/router";
import { ProjectStore } from '@entities/project';
import { Title } from '@angular/platform-browser';
import {UiService} from '@shared/ui';


@Component({
    selector: 'app-project-detail',
    imports: [RouterLink],
    template: `
        <div class="bg-lime-200 p-4 rounded-xl mb-4">
            <nav class="flex flex-wrap gap-4">
                <a href="#" class="hover:underline">Invite members</a>
                <a href="#" class="hover:underline">Roles</a>
                <a href="#" class="hover:underline">Webhooks</a>
                <a href="#" class="hover:underline">Notifications</a>
                <a href="#" class="hover:underline">Emails</a>
                <a href="#" class="hover:underline">Integrations</a>
            </nav>
        </div>

        <div class="bg-gray-200 p-4 rounded-xl mb-4">
            <h2 class="font-medium text-lg mb-4">Boards</h2>
            <nav class="flex flex-wrap gap-4">
                <button class="bg-teal-500 text-white rounded-lg px-4 py-2 cursor-pointer">Development</button>
                <button class="bg-teal-500 text-white rounded-lg px-4 py-2 cursor-pointer">Design</button>
                <button class="bg-teal-500 text-white rounded-lg px-4 py-2 cursor-pointer">Marketing/SEO</button>
            </nav>
        </div>

        <div class="bg-amber-100 p-4 rounded-xl mb-4">
            <a class="font-medium text-lg mb-4 hover:underline">Calendar</a>
        </div>

        <div class="bg-lime-100 p-4 rounded-xl mb-4">
            <a class="font-medium text-lg mb-4 hover:underline">Documents/Notes</a>
        </div>

        <a [routerLink]="['edit']"
           class="bg-teal-500 text-white rounded-lg px-3 py-1 cursor-pointer inline-block">Edit</a>
    `,
})
export class ProjectDetail {
    projectId: number
    private activatedRoute = inject(ActivatedRoute);
    private titleService = inject(Title);

    private readonly store = inject(ProjectStore);
    project = computed(() => this.store.entities().find(p => p.id === this.projectId));

    constructor() {
        effect(() => {
            const p = this.project()
            if (p) {
                this.titleService.setTitle(`Project: ${p.name}`)
            }
        })

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
