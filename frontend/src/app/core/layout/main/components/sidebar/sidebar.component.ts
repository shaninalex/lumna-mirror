import { Component, DestroyRef, inject } from '@angular/core';
import { Actions, ofType } from '@ngrx/effects';
import { actionToggleSidebar } from '@core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { AsyncPipe, NgClass } from '@angular/common';
import { RouterLink } from "@angular/router";
import { Store } from '@ngrx/store';
import { selectCurrentWorkspaceId } from '@entities/workspace';
import { filter, map, switchMap } from 'rxjs';
import { selectCurrentProject, selectProjectsByWorkspaceID } from '@entities/project';

@Component({
    selector: 'lu-sidebar',
    imports: [NgClass, RouterLink, AsyncPipe],
    styleUrl: './sidebar.component.css',
    templateUrl: './sidebar.component.html',
})
export class SidebarComponent {
    private actions$ = inject(Actions);
    private ref = inject(DestroyRef);
    private store = inject(Store);

    currentWorkspaceId = this.store.selectSignal(selectCurrentWorkspaceId);
    hideSidebar = false;

    currentProject = this.store.selectSignal(selectCurrentProject);

    workspace$ = this.store.select(selectCurrentWorkspaceId).pipe(
        filter(workspaceId => workspaceId !== null),
        switchMap((workspaceId) => 
            this.store.select(selectProjectsByWorkspaceID(workspaceId)).pipe(
                map((projects) => ({workspaceId, projects}))
            )
        ),
    );

    constructor() {
        this.actions$
            .pipe(
                ofType(actionToggleSidebar),
                takeUntilDestroyed(this.ref),
            ).subscribe(() => this.hideSidebar = !this.hideSidebar);
    }
}

