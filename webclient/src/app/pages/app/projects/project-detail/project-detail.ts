import {Component, computed, effect, inject} from '@angular/core';
import { ActivatedRoute, RouterLink } from "@angular/router";
import { ProjectStore } from '@entities/project';
import { Title } from '@angular/platform-browser';


@Component({
    selector: 'app-project-detail',
    imports: [RouterLink],
    template: `
        <h1>Project Detail {{ project()?.name }}</h1>

        <a [routerLink]="['edit']" class="bg-teal-500 text-white rounded-lg px-3 py-1 cursor-pointer inline-block">Edit</a>
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
