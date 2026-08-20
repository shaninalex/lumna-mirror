import { AsyncPipe } from '@angular/common';
import { RouterLink } from '@angular/router';
import { Component, inject } from '@angular/core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';
import { Store } from '@ngrx/store';
import { filter, switchMap, map } from 'rxjs';

import { TrimPipe } from '@shared/utils';
import { selectCurrentWorkspace } from '@entities/workspace';
import { ProjectListItemComponent, selectCurrentProject, selectProjectsByWorkspaceID } from '@entities/project';

@Component({
    selector: 'lu-project-dropdown',
    imports: [
        CdkMenu,
        CdkMenuItem,
        CdkMenuTrigger,
        AsyncPipe,
        TrimPipe,
        ProjectListItemComponent,
        RouterLink,
    ],
    templateUrl: './project-dropdown.component.html',
})
export class ProjectDropdownComponent {
    private store = inject(Store);

    open = false;
    project$ = this.store.select(selectCurrentProject);
    workspace$ = this.store.select(selectCurrentWorkspace).pipe(
        filter((workspace) => workspace !== null),
        switchMap((workspace) =>
            this.store
                .select(selectProjectsByWorkspaceID(workspace.id))
                .pipe(map((projects) => ({ workspace, projects }))),
        ),
    );
}
