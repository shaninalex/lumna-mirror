import { Component, inject, OnInit } from '@angular/core';
import { UiService } from '@shared/ui';

@Component({
    selector: 'app-breadcrumbs',
    template: `
        <ul class="flex items-center gap-2">
            @for (item of links; track $index) {
                <li>{{ item }}</li>
            }
        </ul>
    `,
})
export class BreadCrumbs implements OnInit {
    private ui = inject(UiService);
    links: Array<string> = [];

    ngOnInit(): void {
        this.ui.getUrl().subscribe((url) => this.processUrl(url));
    }

    private processUrl(url: string) {}
}
