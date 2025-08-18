import {Component} from '@angular/core';
import {PageTitleSetter} from '@client/shared/ui';
import {MatMenuModule} from '@angular/material/menu';
import {MatIconButton} from '@angular/material/button';
import {MatIconModule} from '@angular/material/icon';
import {MatTabsModule} from '@angular/material/tabs';
import {BacklogViewComponent} from '@client/pages/main/projects/detail/components';

@Component({
    selector: "ts-project-detail-page",
    template: `
        <div>
            <div class="flex items-center gap-2 mb-4">
                <img src="/assets/img/project.svg" class="w-6"/>
                <h3 class="font-bold text-xl">{{ pageTitle }}</h3>
                <button matIconButton [matMenuTriggerFor]="menu" aria-label="Example icon-button with a menu">
                    <mat-icon>more_vert</mat-icon>
                </button>
                <mat-menu #menu="matMenu">
                    <button mat-menu-item>
                        <mat-icon>dialpad</mat-icon>
                        <span>Redial</span>
                    </button>
                    <button mat-menu-item disabled>
                        <mat-icon>voicemail</mat-icon>
                        <span>Check voice mail</span>
                    </button>
                    <button mat-menu-item>
                        <mat-icon>notifications_off</mat-icon>
                        <span>Disable alerts</span>
                    </button>
                </mat-menu>
            </div>

            <!--
                selectIndex should come's from localStorage.
                Save latest open tab.
             -->
            <mat-tab-group mat-stretch-tabs="false" mat-align-tabs="start" selectedIndex="1">
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
                        content Board
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
    ]
})
export class ProjectDetailPageComponent extends PageTitleSetter {
    pageTitle = "Taskiro";
}
