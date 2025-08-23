import {Component} from '@angular/core';
import {PrimaryLayoutComponent} from '../../app/layouts';
import {RouterOutlet} from '@angular/router';

@Component({
    selector: "jr-page-home",
    imports: [
        PrimaryLayoutComponent,
        RouterOutlet
    ],
    template: `
        <jr-primary-layout>
            <router-outlet />
        </jr-primary-layout>`
})
export class HomeWrapperComponent {}
