import { Component, inject } from '@angular/core'
import { ActivatedRoute, RouterLink } from '@angular/router'

@Component({
    selector: `lu-project-overview-page`,
    imports: [RouterLink],
    template: ` <a class="underline" [routerLink]="['board']" [relativeTo]="route">Board</a> `,
})
export class ProjectOverviewComponent {
    route = inject(ActivatedRoute)
}
