import { Component, Input } from '@angular/core';
import { ProjectModel } from '@entities/project';
import { RouterLink } from '@angular/router';
import { CdkMenu, CdkMenuTrigger } from '@angular/cdk/menu';

@Component({
    selector: 'app-project-card',
    imports: [RouterLink, CdkMenu, CdkMenuTrigger],
    template: `
        <div class="card">
            <div class="card-body">
                <div class="d-flex justify-content-between align-items-center">
                    <a [routerLink]="['/projects', project.id]" class=""
                        style="color: inherit"
                    >
                        <h4 class="mb-0">
                            <i class="fa-regular fa-calendar-days"></i>
                            <b>{{ project.title }}</b>
                        </h4>
                    </a>

                    <button [cdkMenuTriggerFor]="menu" class="btn btn-sm">
                        <i class="fa-solid fa-ellipsis"></i>
                    </button>

                    <ng-template #menu>
                        <div class="dropdown-menu d-block" cdkMenu>
                            <a [routerLink]="['/projects', project.id, 'edit']" class="dropdown-item">Edit</a>
                        </div>
                    </ng-template>
                </div>
            </div>
        </div>
    `,
})
export class ProjectCard {
    @Input() project: ProjectModel;
}
