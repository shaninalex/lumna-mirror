import { Component } from '@angular/core';
import { ProjectsListFeature } from "@features/index";

@Component({
    selector: 'app-project-list',
    imports: [ProjectsListFeature],
    template: `<app-projects-list-feature />`,
})
export class ProjectList {

}
