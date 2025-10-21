import {Component} from '@angular/core';
import {PrimaryLayout} from '@client/shared/layouts';
import {RouterOutlet} from '@angular/router';

@Component({
    selector: 'lu-root',
    imports: [
        PrimaryLayout,
        RouterOutlet,
    ],
    template: `
        <lu-primary-layout>
            <router-outlet/>
        </lu-primary-layout>
    `
})
export class PrimaryRoot {}
