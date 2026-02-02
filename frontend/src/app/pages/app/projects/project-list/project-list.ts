import { Component, inject } from '@angular/core';
import { ProjectsListFeature } from '@features/index';
import { UiService } from '@shared/ui';

@Component({
    selector: 'app-project-list',
    imports: [ProjectsListFeature],
    template: `<app-projects-list-feature />`,
})
export class ProjectList {
    private ui = inject(UiService);

    constructor() {
        this.ui.setPageTitle('Projects');
    }
}
