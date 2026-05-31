import { Component } from "@angular/core";
import { ProjectCreateFeature } from "@features";

@Component({
    selector: "app-project-create-page",
    imports: [ProjectCreateFeature],
    template: `
        <div class="p-4">
            <app-project-create-feature />
        </div>
    `
})
export class ProjectCreatePage {}
