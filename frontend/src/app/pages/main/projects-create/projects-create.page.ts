import { Component } from '@angular/core';
import { GlobalLayout } from '@core/layout';
import { ProjectCreateFeature } from '@features';

@Component({
    selector: 'lu-projects-create-page',
    imports: [GlobalLayout, ProjectCreateFeature],
    templateUrl: './projects-create.page.html',
})
export class ProjectsCreatePage {}
