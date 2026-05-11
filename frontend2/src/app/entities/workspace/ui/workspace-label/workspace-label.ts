import { Component, Input } from "@angular/core";
import { WorkspaceModel } from "@entities/workspace";

@Component({
    selector: "app-workspace-label",
    template: `
        <div class="flex items-center gap-2">
            @switch (size) {
                @case ("md") {
                    <div>
                        <img
                            src="/img/project.svg"
                            alt=""
                            class="w-10 h-10 rounded"
                        />
                    </div>
                    <div class="leading-none mt-2">
                        <div class="font-bold">{{ workspace.title }}</div>
                        <div class="text-sm text-slate-300">
                            Some small text
                        </div>
                    </div>
                }
                @case ("sm") {
                    <div>
                        <img
                            src="/img/project.svg"
                            alt=""
                            class="rounded w-5 h-5"
                        />
                    </div>
                    <div class="leading-none">
                        <div class="font-bold">{{ workspace.title }}</div>
                    </div>
                }
            }
        </div>
    `
})
export class WorkspaceLabel {
    @Input() workspace: WorkspaceModel;
    @Input() size: "sm" | "md" | "lg" = "md";
}
