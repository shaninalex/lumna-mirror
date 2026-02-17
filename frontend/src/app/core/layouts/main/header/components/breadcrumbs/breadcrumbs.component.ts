import { Component, inject, OnInit } from '@angular/core';
import { UiService } from '@shared/ui';

@Component({
    selector: 'app-breadcrumbs',
    template: ``,
})
export class BreadCrumbs implements OnInit {
    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.getUrl().subscribe();
    }
}
