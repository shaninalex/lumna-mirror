import {Component} from '@angular/core';

@Component({
    selector: 'fr-main-layout',
    imports: [],
    template: `
        <div>
            <ng-content></ng-content>
        </div>
    `
})
export class MainLayout {
}
