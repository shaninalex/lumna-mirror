import {Component, Input} from '@angular/core';
import {CdkMenuModule} from '@angular/cdk/menu';
import {MatTabsModule} from '@angular/material/tabs';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {BacklogViewComponent} from '@client/entities/project/ui/detail/components';
import {BoardViewComponent} from '@client/features/project';

@Component({
    selector: "ts-project-detail",
    imports: [
        BacklogViewComponent,
        BoardViewComponent,
        CdkMenuModule,
        MatIconModule,
        MatTabsModule,
        MatButtonModule,
    ],
    template: `
        <div>
            <div class="flex items-center gap-2 mb-4">
                <img src="/assets/img/project.svg" class="w-6"/>
                <h3 class="font-bold text-xl">Taskiro</h3>
                <button matIconButton [cdkMenuTriggerFor]="project_detail_menu" class="example-standalone-trigger">
                    <mat-icon>more_vert</mat-icon>
                </button>
                <ng-template #project_detail_menu>
                    <div class="bg-white flex flex-col gap-2 border p-2 rounded" cdkMenu>
                        <button cdkMenuItem>Refresh</button>
                        <button cdkMenuItem>Settings</button>
                        <button cdkMenuItem>Help</button>
                        <button cdkMenuItem>Sign out</button>
                    </div>
                </ng-template>
            </div>

            <!--
                selectIndex should come's from localStorage.
                Save latest open tab.
             -->
            <mat-tab-group mat-stretch-tabs="false" mat-align-tabs="start" selectedIndex="2">
                <mat-tab label="Summary">
                    <div class="py-4">
                        content Summary
                    </div>
                </mat-tab>
                <mat-tab label="Backlog">
                    <div class="py-4">
                        <ts-backlog-view/>
                    </div>
                </mat-tab>
                <mat-tab label="Board">
                    <div class="py-4 w-full">
                        <ts-board-view [projectKey]="projectKey"/>
                    </div>
                </mat-tab>
                <mat-tab label="Calendar">
                    <div class="py-4">
                        Calendar
                    </div>
                </mat-tab>
                <mat-tab label="Pages">
                    <div class="py-4">
                        content Pages
                    </div>
                </mat-tab>
            </mat-tab-group>
        </div>
    `
})
export class ProjectDetailComponent {
    @Input() projectKey: string;
}
