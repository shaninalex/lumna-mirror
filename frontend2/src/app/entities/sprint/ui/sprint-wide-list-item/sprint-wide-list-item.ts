import { Component, Input } from "@angular/core";
import { SprintModel } from "@entities/sprint";

@Component({
    selector: "app-sprint-wide-list-item",
    imports: [],
    template: `
        <!-- Sprint -->
        <section
            class="mb-8 overflow-hidden rounded-md border border-zinc-200 bg-white"
        >
            <div
                class="flex items-center justify-between border-b border-zinc-200 bg-zinc-50 px-5 py-4"
            >
                <div>
                    <h2 class="text-lg font-semibold">{{ sprint.title }}</h2>
                    <p class="text-sm text-zinc-500">
                        14 issues • 32 story points
                    </p>
                </div>

                <button
                    class="rounded-md border border-zinc-300 px-3 py-2 text-sm hover:bg-white"
                >
                    Complete Sprint
                </button>
            </div>

            <div class="divide-y divide-zinc-200">
                <!-- Task -->
                <div class="flex items-start gap-2 px-5 py-4 hover:bg-zinc-50">
                    <div class="mt-1 h-3 w-3 rounded-full bg-green-500"></div>

                    <div class="flex-1">
                        <div class="mb-1 flex items-center gap-2">
                            <span
                                class="rounded bg-blue-100 px-2 py-0.5 text-xs font-medium text-blue-700"
                            >
                                LUM-142
                            </span>

                            <h3 class="font-medium">
                                Implement drag and drop for backlog items
                            </h3>
                        </div>

                        <p class="text-sm text-zinc-500">
                            Add sortable drag-n-drop support between backlog and
                            sprint sections.
                        </p>
                    </div>

                    <div class="flex items-center gap-2 text-xs">
                        <span
                            class="rounded bg-yellow-100 px-2 py-1 text-yellow-700"
                        >
                            In Progress
                        </span>

                        <span class="rounded bg-zinc-200 px-2 py-1">
                            5 SP
                        </span>
                    </div>
                </div>

                <!-- Task -->
                <div class="flex items-start gap-2 px-5 py-4 hover:bg-zinc-50">
                    <div class="mt-1 h-3 w-3 rounded-full bg-red-500"></div>

                    <div class="flex-1">
                        <div class="mb-1 flex items-center gap-2">
                            <span
                                class="rounded bg-blue-100 px-2 py-0.5 text-xs font-medium text-blue-700"
                            >
                                LUM-143
                            </span>

                            <h3 class="font-medium">
                                Add sprint statistics widget
                            </h3>
                        </div>

                        <p class="text-sm text-zinc-500">
                            Show completed tasks, velocity and remaining
                            workload.
                        </p>
                    </div>

                    <div class="flex items-center gap-2 text-xs">
                        <span
                            class="rounded bg-zinc-200 px-2 py-1 text-zinc-700"
                        >
                            Todo
                        </span>

                        <span class="rounded bg-zinc-200 px-2 py-1">
                            3 SP
                        </span>
                    </div>
                </div>

                <!-- Task -->
                <div class="flex items-start gap-2 px-5 py-4 hover:bg-zinc-50">
                    <div class="mt-1 h-3 w-3 rounded-full bg-purple-500"></div>

                    <div class="flex-1">
                        <div class="mb-1 flex items-center gap-2">
                            <span
                                class="rounded bg-blue-100 px-2 py-0.5 text-xs font-medium text-blue-700"
                            >
                                LUM-144
                            </span>

                            <h3 class="font-medium">
                                project role permissions
                            </h3>
                        </div>

                        <p class="text-sm text-zinc-500">
                            Create base RBAC structure for workspaces and
                            projects.
                        </p>
                    </div>

                    <div class="flex items-center gap-2 text-xs">
                        <span
                            class="rounded bg-green-100 px-2 py-1 text-green-700"
                        >
                            Done
                        </span>

                        <span class="rounded bg-zinc-200 px-2 py-1">
                            8 SP
                        </span>
                    </div>
                </div>
            </div>
        </section>
    `
})
export class SprintWideListItem {
    @Input() sprint: SprintModel;
}
