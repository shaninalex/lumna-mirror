import { Component, Input } from '@angular/core';
import { ProjectModel } from '@entities/project';
import { RouterLink } from '@angular/router';
import { CdkMenu, CdkMenuTrigger } from '@angular/cdk/menu';

@Component({
    selector: 'app-project-card',
    imports: [RouterLink, CdkMenu, CdkMenuTrigger],
    template: `
        <div class="card">
            <div class="card-content">
                <div class="is-flex is-justify-content-space-between">
                    <a [routerLink]="['/projects', project.id]" class="is-flex is-justify-content-space-between is-align-items-center is-size-5"
                        style="color: inherit"
                    >
                        <i class="fa-regular fa-calendar-days"></i>
                        <b>{{ project.title }}</b>
                    </a>

                    <button [cdkMenuTriggerFor]="menu">
                        <i class="fa-solid fa-ellipsis"></i>
                    </button>

                    <ng-template #menu>
                        <div class="dropdown p-4" cdkMenu>
                            <a [routerLink]="['/projects', project.id, 'edit']">Edit</a>
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
