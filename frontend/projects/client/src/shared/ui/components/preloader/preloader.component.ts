

import {Component} from '@angular/core';

@Component({
    selector: 'jr-preloader',
    template: `
        <div class="spinner-border" role="status">
            <span class="visually-hidden">Loading...</span>
        </div>
    `
})
export class PreloaderComponent {}
