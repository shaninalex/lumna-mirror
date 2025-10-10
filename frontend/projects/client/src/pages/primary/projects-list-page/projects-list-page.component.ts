import {Component} from '@angular/core';
import {ProjectListFeatureComponent} from '@client/features/project';

@Component({
    selector: 'fr-projects-list',
    imports: [
        ProjectListFeatureComponent
    ],
    template: `
        <fr-project-list-feature/>
    `
})
export class ProjectsListPageComponent {
}
