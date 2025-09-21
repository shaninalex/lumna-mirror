import {Component} from '@angular/core';
import {PrimaryLayout} from '@client/shared/layouts';
import {RouterOutlet} from '@angular/router';

@Component({
    selector: 'fr-root',
    imports: [
        PrimaryLayout,
        RouterOutlet,
    ],
    template: `
        <fr-primary-layout>
            <router-outlet/>
        </fr-primary-layout>
    `
})
export class PrimaryRoot {}
