import {Component} from '@angular/core';
import {RouterOutlet} from '@angular/router';
import {MatButtonModule} from '@angular/material/button';
import {CdkMenuModule} from '@angular/cdk/menu';

@Component({
    selector: 'fr-root',
    imports: [
        RouterOutlet,
        MatButtonModule,
        CdkMenuModule,
    ],
    template: `<router-outlet/>`,
})
export class App {}
