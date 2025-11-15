import { Component, inject } from '@angular/core'
import { ActivatedRoute, RouterLink } from '@angular/router'

@Component({
    selector: `lu-project-overview-page`,
    imports: [RouterLink],
    template: `
	<div>
		<div class="card flex flex-col gap-4">
			<div class="card-title">Project Lumna</div>
			<div class="flex flex-col gap-2">
				<a class="hover-space inline-block" [routerLink]="['board']" [relativeTo]="route">
					<span class="inline-flex gap-2 items-center">
						<i class="text-lg i-board"></i>
						Default Task board ( 4 active tasks )
					</span>
				</a>
				<div>
					<button class="btn btn-primary">Add board</button>
				</div>
			</div>
			<div>
				<div class="card-title">Calendar</div>
				<button class="btn btn-primary">Add calendar</button>
			</div>
			<div>
				<div class="card-title">Docs</div>
				<button class="btn btn-primary">Add document</button>
			</div>
		</div>
	</div>`,
})
export class ProjectOverviewComponent {
    route = inject(ActivatedRoute)
}
