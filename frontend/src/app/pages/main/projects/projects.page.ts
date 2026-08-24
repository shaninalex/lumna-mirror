import type { OnInit } from '@angular/core';
import { Component, inject } from '@angular/core';
import { GlobalLayout } from '@core/layout';
import { RouterLink } from "@angular/router";
import { UiService } from '@shared/ui';
import { ProjectCardComponent, selectProjects } from '@entities/project';
import { selectWorkspaces } from '@entities/workspace';
import { Store } from '@ngrx/store';
import { filter, switchMap, map } from 'rxjs';
import { AsyncPipe } from '@angular/common';
import { AppRoutes } from '@core';

@Component({
    selector: 'lu-projects-page',
    imports: [GlobalLayout, RouterLink, AsyncPipe, ProjectCardComponent],
    templateUrl: './projects.page.html',
})
export class ProjectsPage implements OnInit {
    private ui = inject(UiService);
    private store = inject(Store);
    readonly appRoutes = inject(AppRoutes);

    workspace$ = this.store.select(selectWorkspaces.currentWorkspace).pipe(
        filter(workspace => workspace !== null),
        switchMap((workspace) => 
            this.store.select(selectProjects.byWorkspaceId(workspace.id)).pipe(
                map((projects) => ({workspace, projects}))
            )
        ),
    );

    ngOnInit(): void {
        this.ui.setPageTitle("Projects")
    }
}
