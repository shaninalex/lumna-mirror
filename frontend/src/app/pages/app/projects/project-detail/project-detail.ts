import { Component, inject, OnInit } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { ProjectModel, ProjectState, selectProjectByID } from '@entities/project';
import { actionBoardGetList, BoardsList } from '@entities/board';
import { filter, Observable, switchMap, tap } from 'rxjs';
import { Store } from '@ngrx/store';
import { UiService } from '@shared/ui';
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
export class ProjectDetail implements OnInit {
    private ui = inject(UiService);
    private route = inject(ActivatedRoute);
    private store = inject(Store<ProjectState>);
    project$: Observable<ProjectModel>;

    ngOnInit() {
        this.project$ = this.route.params.pipe(
            switchMap((params) =>
                this.store.select(selectProjectByID(params['id'])).pipe(
                    filter((project) => !!project),
                    tap((project) => this.ui.setPageTitle(`Project: ${project.title}`)),
                ),
            ),
        );
    }
}
