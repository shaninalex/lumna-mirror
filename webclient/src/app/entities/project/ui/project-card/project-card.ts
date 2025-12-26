import {Component, EventEmitter, Input, Output} from '@angular/core';
import {ProjectModel} from '@entities/project';
import {RouterLink} from '@angular/router';
import {CdkMenu, CdkMenuItem, CdkMenuTrigger} from '@angular/cdk/menu';

@Component({
    selector: 'app-project-card',
    imports: [
        RouterLink,
        CdkMenu,
        CdkMenuItem,
        CdkMenuTrigger
    ],
    template: `
        <div class="bg-gray-100 rounded-xl p-4 block">
            <div class="flex justify-between">
                <a [routerLink]="['/projects', project.id]" class="flex justify-between items-center">
                    <i class="fa-regular fa-calendar-days"></i>
                    <div>{{ project.name }}</div>
                </a>

                <button [cdkMenuTriggerFor]="menu">
                    <i class="fa-solid fa-ellipsis"></i>
                </button>

                <ng-template #menu>
                    <div class="bg-white border border-gray-200 rounded-xl p-4" cdkMenu>
                        <button class="block" cdkMenuItem>Edit</button>
                        <button class="block" cdkMenuItem (click)="deleteProject()">Delete</button>
                    </div>
                </ng-template>

            </div>
        </div>
    `,
})
export class ProjectCard {
    @Input() project: ProjectModel
    @Output() onDelete: EventEmitter<number> = new EventEmitter<number>();

    deleteProject(): void {
        this.onDelete.emit(this.project.id)
    }
}
