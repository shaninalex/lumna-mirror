import { Component } from '@angular/core';
import { StaticLayout } from "@core/layout";

@Component({
    selector: 'lu-page-404-page',
    imports: [StaticLayout],
    template: `
        <lu-static-layout>
            <h1>Page Not Found</h1>
        </lu-static-layout>
    `,
})
export class Page404 {}
