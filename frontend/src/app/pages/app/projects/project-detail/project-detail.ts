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
        <nav class="navbar navbar-expand-sm bg-body-tertiary mb-4">
            <div class="container-fluid">
                <div class="collapse navbar-collapse show">
                    <ul class="navbar-nav me-auto flex-wrap">
                        <li class="nav-item"><a class="nav-link" href="#">Calendar</a></li>
                        <li class="nav-item"><a class="nav-link" href="#">Documents</a></li>
                        <li class="nav-item"><a class="nav-link" href="#">Invite members</a></li>
                        <li class="nav-item"><a class="nav-link" href="#">Roles</a></li>
                        <li class="nav-item"><a class="nav-link" href="#">Webhooks</a></li>
                        <li class="nav-item"><a class="nav-link" href="#">Notifications</a></li>
                        <li class="nav-item"><a class="nav-link" href="#">Emails</a></li>
                        <li class="nav-item"><a class="nav-link" href="#">Integrations</a></li>
                    </ul>
                    <ul class="navbar-nav">
                        <li class="nav-item"><a [routerLink]="['edit']" class="nav-link">Edit</a></li>
                    </ul>
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
