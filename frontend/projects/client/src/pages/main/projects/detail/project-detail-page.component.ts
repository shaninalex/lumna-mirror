import {Component} from '@angular/core';
import {PageTitleSetter} from '@client/shared/ui';
import {MatMenuModule} from '@angular/material/menu';
import {MatIconButton} from '@angular/material/button';
import {MatIconModule} from '@angular/material/icon';
import {MatTabsModule} from '@angular/material/tabs';
import {BacklogViewComponent, BoardViewComponent} from '@client/pages/main/projects/detail/components';
import {CdkMenuModule} from '@angular/cdk/menu';

@Component({
    selector: "ts-project-detail-page",
    template: `
        <div>
            <div class="flex items-center gap-2 mb-4">
                <img src="/assets/img/project.svg" class="w-6"/>
                <h3 class="font-bold text-xl">{{ pageTitle }}</h3>
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
                    <div class="py-4">
                        <ts-board-view/>
                    </div>
                </mat-tab>
                <mat-tab label="Timeline">
                    <div class="py-4">
                        content Timeline
                    </div>
                </mat-tab>
                <mat-tab label="Pages">
                    <div class="py-4">
                        content Pages
                    </div>
                </mat-tab>
            </mat-tab-group>
        </div>`,
    imports: [
        MatMenuModule,
        MatIconModule,
        MatIconButton,
        MatTabsModule,
        BacklogViewComponent,
        CdkMenuModule,
        BoardViewComponent,
    ]
})
export class ProjectDetailPageComponent extends PageTitleSetter {
    pageTitle = "Taskiro";
}
