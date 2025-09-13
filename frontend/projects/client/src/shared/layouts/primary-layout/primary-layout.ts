import {Component} from '@angular/core';
import {Header} from '@client/shared/layouts/primary-layout/components/header/header';
import {Sidebar} from '@client/shared/layouts/primary-layout/components/sidebar/sidebar';

@Component({
    selector: 'fr-primary-layout',
    imports: [
        Header,
        Sidebar
    ],
    template: `
        <div class="h-screen overflow-hidden flex">
            <fr-sidebar />
            <div class="flex flex-col flex-grow ">
                <fr-header />
                <div class="flex-grow p-4 border-t border-l border-slate-300 rounded-lg overflow-y-auto">
                    <ng-content></ng-content>
                </div>
            </div>
        </div>
    `
})
export class PrimaryLayout {
}
