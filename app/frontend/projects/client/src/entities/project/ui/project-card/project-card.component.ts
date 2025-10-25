import { Component, Input } from "@angular/core"
import { Project } from "@client/entities/project"
import { RouterLink } from "@angular/router"

@Component({
	selector: "lu-project-card",
	imports: [RouterLink],
	template: `
		<div class="card hover-space">
			<div class="card-title">
				<a [routerLink]="[project.code]">
					{{ project.title }}
				</a>
			</div>
		</div>
	`,
})
export class ProjectCardComponent {
	@Input() project: Project
}
