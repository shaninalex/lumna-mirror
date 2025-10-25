import { Component, EventEmitter, Input, Output } from "@angular/core"
import { Task } from "@client/entities/task"
import { DatePipe } from "@angular/common"

@Component({
	selector: "lu-task-card",
	imports: [DatePipe],
	template: `
		<div class="card">
			<div class="card-title mb-1">
				<button (click)="openTaskDetail.emit(task.code)" class="cursor-pointer">
					{{ task.title }}
					@if (task.completed) {
						<i class="i-check-circle text-lg text-green-500"></i>
					}
				</button>
			</div>
			<div>
				<div class="flex items-center gap-2">
					<div class="text-sm">
						<div class="text-xs">{{ task.created_at | date: "EEE, MMM d, HH:mm:ss" }}</div>
					</div>
					<div class="ms-auto">
						<img src="/img/1.png" class="w-6 rounded-full" title="Username" />
					</div>
				</div>
			</div>
		</div>
	`,
})
export class TaskCardComponent {
	@Input() task: Task
	@Input() projectCode: string
	@Output() openTaskDetail: EventEmitter<string> = new EventEmitter()
}
