import { Component } from "@angular/core"
import { ProjectListFeatureComponent } from "@client/features/project"

@Component({
	selector: "lu-projects-list",
	imports: [ProjectListFeatureComponent],
	template: `<lu-project-list-feature /> `,
})
export class ProjectsListPageComponent {}
