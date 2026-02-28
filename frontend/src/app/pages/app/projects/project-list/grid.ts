import {Component} from '@angular/core';

@Component({
    selector: "app-project-list--grid",
    template: `
        <div class="fixed-grid has-3-cols">
            <div class="grid">
                <ng-content />
            </div>
        </div>
    `
})
export class projectListGrid {}
