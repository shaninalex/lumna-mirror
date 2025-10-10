import { Component } from '@angular/core';
import {RouterLink} from '@angular/router';

@Component({
    selector: 'fr-404-page',
    template: `
        <div class="text-center my-10">
            <h1 class="font-bold text-xl">Page not found</h1>
            <a [routerLink]="['/']" class="underline text-sky-600">Home</a>
        </div>
    `,
    imports: [
        RouterLink
    ]
})
export class Page404 {

}
