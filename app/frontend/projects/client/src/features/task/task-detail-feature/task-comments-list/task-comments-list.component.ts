import { Component, Input } from "@angular/core"
import { Task } from "@client/entities/task"
import { DatePipe } from "@angular/common"

@Component({
	selector: 'lu-task-comments-list',
	imports: [
		DatePipe
	],
	template: `
	@if (task.comments) {
		@for (comment of task.comments; track $index) {
			<div class="flex items-start gap-2">
				<div>
					<img src="img/1.png" alt="" class="rounded-full w-8">
				</div>

				<div>
					<div class="flex items-center gap-2 mb-2">
						<div>John Doe</div>
						<div class="text-xs">{{ comment.created_at | date: "MMM d, HH:mm" }}</div>
					</div>
					<div>
						{{ comment.content }}
					</div>
				</div>
			</div>
		}
	}
`
})
export class CommentsListComponent {
	@Input() task: Task
}
