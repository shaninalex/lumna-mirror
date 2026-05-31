import { Component, Input } from "@angular/core";
import { ProjectModel } from "@entities/project";
import { TruncPipe } from "@shared/pipes";

@Component({
    selector: "app-project-icon",
    imports: [TruncPipe],
    template: `
        <div
            class="flex h-12 w-12 items-center justify-center rounded-md bg-blue-600 text-sm font-bold text-white"
        >
            {{ project.title | trunc }}
        </div>
    `
})
export class ProjectIcon {
    @Input() project: ProjectModel;

    // TODO: get project color from project.meta
}
