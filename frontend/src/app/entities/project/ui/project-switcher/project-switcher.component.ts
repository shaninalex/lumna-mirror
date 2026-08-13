import { Component, inject } from '@angular/core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';
import { RouterLink } from "@angular/router";
import { Store } from '@ngrx/store';
import { selectCurrentWorkspace } from '@entities/workspace/model';
import { AsyncPipe } from '@angular/common';
import { selectCurrentProject, selectProjectsByWorkspaceID } from '@entities/project/model';
import { filter, map, switchMap, tap } from 'rxjs';
import { ProjectListItemComponent } from "../project-list-item";


@Component({
    selector: 'lu-project-switcher',
    imports: [CdkMenu, CdkMenuItem, CdkMenuTrigger, RouterLink, AsyncPipe, ProjectListItemComponent],
    template: `
        @if(workspace$ | async; as workspace) {
            <button [cdkMenuTriggerFor]="menu" class="btn btn-sm btn-outline-secondary">
                @if (project$ | async; as project) {
                    @if (project) {
                        {{ project.title }}
                    }
                } @else {
                    Select project
                }
                <i class="fa-solid fa-chevron-down"></i>
            </button>
        
            <ng-template #menu>
                <div class="list-group" cdkMenu>
                    @if (workspace.projects.length) {
                        @for (item of workspace.projects; track $index) {
                            <lu-project-list-item [project]="item" />
                        }
                        
                        <a cdkMenuItem [routerLink]="['/app/w', workspace.workspace.id, 'projects']" class="list-group-item list-group-item-action text-decoration-underline">
                            See all
                        </a>
                    } @else {
                        <a [routerLink]="['/app/w', workspace.workspace.id, 'projects', 'create']" class="list-group-item list-group-item-action text-decoration-underline">Create project</a>
                    }
                </div>
            </ng-template>
        }
    `,
})
export class ProjectSwitcherComponent {
    private store = inject(Store);

    project$ = this.store.select(selectCurrentProject);
    workspace$ = this.store.select(selectCurrentWorkspace).pipe(
        filter(workspace => workspace !== null),
        switchMap((workspace) => 
            this.store.select(selectProjectsByWorkspaceID(workspace.id)).pipe(
                map((projects) => ({workspace, projects}))
            )
        ),
    );
}
