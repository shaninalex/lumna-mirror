import { Component, inject, OnInit } from '@angular/core';
import { GlobalLayout } from '@core/layout';
import { RouterLink } from "@angular/router";
import { UiService } from '@shared/ui';

@Component({
    selector: 'lu-projects-page',
    imports: [GlobalLayout, RouterLink],
    templateUrl: './projects.page.html',
})
export class ProjectsPage implements OnInit {
    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.setPageTitle("Projects")
    }
}
