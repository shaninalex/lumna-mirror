import {Component, inject} from '@angular/core';
import {PageTitleSetter} from '@client/shared/ui';
import {ProjectDetailComponent} from '@client/entities/project';
import {ActivatedRoute} from '@angular/router';

@Component({
    selector: "ts-project-detail-page",
    template: `
        <ts-project-detail [projectKey]="projectKey" />
    `,
    imports: [
        ProjectDetailComponent,
    ]
})
export class ProjectDetailPageComponent extends PageTitleSetter {
    pageTitle = "Taskiro project";
    route = inject(ActivatedRoute)

    projectKey: string = this.route.snapshot.params["projectKey"];
}
