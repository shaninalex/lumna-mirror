import {Component, inject} from '@angular/core';
import {UiService} from '@shared/ui';
import {toSignal} from '@angular/core/rxjs-interop';

@Component({
    selector: 'app-header',
    imports: [],
    templateUrl: './header.html',
    styleUrl: './header.css',
})
export class Header {
    private readonly ui = inject(UiService);
    readonly title = toSignal(this.ui.pageTitle);
}
