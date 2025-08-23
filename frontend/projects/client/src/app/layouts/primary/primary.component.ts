import {Component} from '@angular/core';
import {HeaderComponent, SidebarComponent} from './components';

@Component({
    selector: "jr-primary-layout",
    imports: [
        HeaderComponent,
        SidebarComponent
    ],
    template: `
        <div class="h-screen overflow-hidden flex flex-col">
            <ts-header />
            <div class="flex flex-grow">
                <ts-sidebar />
                <div class="flex-grow bg-slate-100 p-4">
                    <ng-content></ng-content>
                </div>
            </div>
        </div>`
})
export class PrimaryLayoutComponent {}
