import type { OnInit } from '@angular/core';
import { Component, inject } from '@angular/core';
import { MainLayout } from '@core/layout';
import { UiService } from '@shared/ui';
import { RouterLink } from "@angular/router";
import { AppRoutes } from '@core';
import { selectTasksByProject, TaskListItemComponent } from "@entities/task";
import { Store } from '@ngrx/store';
import { selectCurrentProjectId } from '@entities/project';
import { filter, map, switchMap } from 'rxjs';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'lu-backlog-page',
    imports: [MainLayout, RouterLink, TaskListItemComponent, AsyncPipe],
    templateUrl: './backlog.page.html',
})
export class BacklogPage implements OnInit {
    private ui = inject(UiService);
    private store = inject(Store);
    readonly appRoutes = inject(AppRoutes);

    data$ = this.store.select(selectCurrentProjectId).pipe(
        filter((projectId) => projectId !== null),
        switchMap((projectId) => 
            this.store.select(selectTasksByProject(projectId)).pipe(
                map((tasks) => ({
                    projectId, 
                    tasks
                }))
            )
        ),
    )
    
    ngOnInit(): void {
        this.ui.setPageTitle("Backlog");
    }

    returnTo(): string {
        return this.appRoutes.backlog().join('/')
    }
}
