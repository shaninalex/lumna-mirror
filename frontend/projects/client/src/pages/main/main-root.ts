import {Component} from '@angular/core';
import {MainLayout} from '@client/app/layouts';
import {RouterOutlet} from '@angular/router';

@Component({
    selector: 'fr-root',
    imports: [
        MainLayout,
        RouterOutlet
    ],
    template: `
        <fr-main-layout>
            <router-outlet/>
        </fr-main-layout>
    `
})
export class MainRoot {
}
