import {Component} from '@angular/core';
import {PageTitleSetter} from '@client/shared/ui';
import {ProjectListComponent} from '@client/entities/project';

@Component({
    selector: "ts-project-list-page",
    template: `<ts-projects-list />`,
    imports: [
        ProjectListComponent,
    ]
})
export class ProjectListPageComponent extends PageTitleSetter {
    pageTitle = "Projects";
}
