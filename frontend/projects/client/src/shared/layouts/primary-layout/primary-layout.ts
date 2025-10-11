import {Component, inject} from '@angular/core';
import {HeaderComponent, SidebarComponent, ToastManagerComponent} from './components';
import {AsyncPipe, NgClass} from '@angular/common';
import {Observable} from 'rxjs';
import {UiService} from '@client/shared/ui/ui.service';

@Component({
    selector: 'fr-primary-layout',
    imports: [
        SidebarComponent,
        HeaderComponent,
        NgClass,
        AsyncPipe,
        ToastManagerComponent
    ],
    styleUrl: './primary-layout.scss',
    template: `
        <div class="h-screen overflow-hidden flex" [ngClass]="{'sidebar-closed': closeSidebar$ | async}">
            <div class="layout-sidebar">
                <fr-sidebar/>
            </div>
            <div class="flex flex-col flex-grow layout-content">
                <fr-header/>
                <div class="flex-grow p-4 overflow-y-auto bg-slate-50 dark:bg-gray-700">
                    <ng-content></ng-content>
                    <lu-toast-manager />
                </div>
            </div>
        </div>
    `
})
export class PrimaryLayout {
    private uiService = inject(UiService);
    closeSidebar$: Observable<boolean> = this.uiService.extendSidebar();
}
