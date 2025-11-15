import { Component, Input } from "@angular/core"
import { Project } from "@client/entities/project"
import { RouterLink } from "@angular/router"

@Component({
	selector: "lu-project-card",
	imports: [RouterLink],
	template: `
		<a [routerLink]="[project.code]" class="card hover-space block">
			<div class="card-title">
				{{ project.title }}
			</div>
		</a>
	`,
})
export class ProjectCardComponent {
	@Input() project: Project
}
