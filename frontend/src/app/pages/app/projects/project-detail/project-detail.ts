import { Component, inject } from '@angular/core';
import { ActivatedRoute, RouterLink, RouterOutlet } from '@angular/router';
import { ProjectModel } from '@entities/project';
import { BoardsList } from '@entities/board';
import { filter, map, Observable, tap } from 'rxjs';
import { AsyncPipe } from '@angular/common';
import { UiService } from '@shared/ui';

@Component({
    selector: 'app-project-detail',
    imports: [RouterLink, BoardsList, AsyncPipe],
    template: `
        <nav class="navbar mb-2">
            <div class="navbar-menu">
                <div class="navbar-start">
                    <a class="navbar-item">Calendar</a>
                    <a class="navbar-item">Documents</a>
                    <a class="navbar-item">Invite members</a>
                    <a class="navbar-item">Roles</a>
                    <a class="navbar-item">Webhooks</a>
                    <a class="navbar-item">Notifications</a>
                    <a class="navbar-item">Emails</a>
                    <a class="navbar-item">Integrations</a>
                </div>

                <div class="navbar-end">
                    <a [routerLink]="['edit']" class="navbar-item">Edit</a>
                </div>
            </div>
        </nav>

        @if (project$ | async; as project) {
            <app-boards-list [projectId]="project.id" />
        }
    `,
})
export class ProjectDetail {
    private route = inject(ActivatedRoute);
    private ui = inject(UiService);

    project$: Observable<ProjectModel> = this.route.data.pipe(
        filter((data) => !!data['project']),
        map((data) => data['project'] as ProjectModel),
        tap((project) => this.ui.setPageTitle(`Project: ${project.title}`)),
    );
}

@Component({
    selector: 'app-project-container',
    imports: [RouterOutlet],
    template: `<router-outlet />`,
})
export class ProjectContainer {}
