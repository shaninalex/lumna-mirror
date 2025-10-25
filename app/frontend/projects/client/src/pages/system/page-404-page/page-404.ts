import { Component } from "@angular/core"
import { RouterLink } from "@angular/router"

@Component({
	selector: "lu-404-page",
	template: `
		<div class="my-10 text-center">
			<h1 class="text-xl font-bold">Page not found</h1>
			<a [routerLink]="['/']" class="text-sky-600 underline">Home</a>
		</div>
	`,
	imports: [RouterLink],
})
export class Page404 {}
