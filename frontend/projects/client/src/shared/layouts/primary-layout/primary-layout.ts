import {Component} from '@angular/core';
import {HeaderComponent, SidebarComponent} from './components';

@Component({
    selector: 'fr-primary-layout',
    imports: [
        SidebarComponent,
        HeaderComponent
    ],
    template: `
        <div class="h-screen overflow-hidden flex">
            <fr-sidebar />
            <div class="flex flex-col flex-grow ">
                <fr-header />
                <div class="flex-grow p-4 overflow-y-auto bg-slate-100">
                    <ng-content></ng-content>
                </div>
            </div>
        </div>
    `
})
export class PrimaryLayout {
}
