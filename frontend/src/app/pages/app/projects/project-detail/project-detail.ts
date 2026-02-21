import { Component, inject } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { ProjectModel } from '@entities/project';
import { BoardsList } from '@entities/board';
import { filter, map, Observable } from 'rxjs';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-project-detail',
    imports: [RouterLink, BoardsList, AsyncPipe],
    template: `
        <div class="bg-lime-200 dark:bg-lime-800 card">
            <nav class="flex flex-wrap gap-4">
                <a class="hover:underline">Invite members</a>
                <a class="hover:underline">Roles</a>
                <a class="hover:underline">Webhooks</a>
                <a class="hover:underline">Notifications</a>
                <a class="hover:underline">Emails</a>
                <a class="hover:underline">Integrations</a>
            </nav>
        </div>

        @if (project$ | async; as project) {
            <app-boards-list [projectId]="project.id" />
        }

        <div class="bg-amber-100 dark:bg-amber-800 card">
            <a class="font-medium text-lg mb-4 hover:underline">Calendar</a>
        </div>

        <div class="bg-lime-100 dark:bg-lime-800 card">
            <a class="font-medium text-lg mb-4 hover:underline">Documents/Notes</a>
        </div>

        <a [routerLink]="['edit']" class="btn btn-sm btn-primary">Edit</a>
    `,
})
export class ProjectDetail {
    private route = inject(ActivatedRoute);
    project$: Observable<ProjectModel> = this.route.data.pipe(
        filter((data) => !!data['project']),
        map((data) => data['project'] as ProjectModel),
    );
}
