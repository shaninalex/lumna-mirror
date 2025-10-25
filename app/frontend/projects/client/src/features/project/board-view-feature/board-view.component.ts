import { Component, ElementRef, HostListener, inject, Input, OnInit, Renderer2 } from "@angular/core"
import { selectTasksByProjectID, Task, TaskCardComponent, TaskChangeStatusAction } from "@client/entities/task"
import { CdkDrag, CdkDragDrop, CdkDropList, CdkDropListGroup, moveItemInArray, transferArrayItem } from "@angular/cdk/drag-drop"
import { BoardViewApiService } from "./api"
import { StatusColumn } from "./board.model"
import { Store } from "@ngrx/store"
import { AppState } from "@client/shared/store"
import { Project } from "@client/entities/project"
import { ColumnHeaderComponent } from "@client/features/project/board-view-feature/components"
import { CreateStatusFormComponent, selectProjectStatusList } from "@client/entities/status"
import { combineLatest, map, Observable } from "rxjs"
import { AsyncPipe } from "@angular/common"
import { TaskFormSmComponent } from "@client/features/project"
import { Router } from "@angular/router"

@Component({
	selector: "lu-board-view-feature",
	imports: [
		CdkDropList,
		CdkDrag,
		CdkDropListGroup,
		TaskCardComponent,
		ColumnHeaderComponent,
		AsyncPipe,
		CreateStatusFormComponent,
		TaskFormSmComponent,
	],
	providers: [BoardViewApiService],
	styleUrl: "./board-view.component.scss",
	template: `
		@if (columns$ | async; as columns) {
			<div cdkDropListGroup class="no-wrap flex h-full w-full items-start justify-start gap-4">
				@for (column of columns; track $index) {
					<div class="card board-column">
						<lu-column-header [project]="project" [column]="column" />

						<div
							class="my-4 flex h-full min-h-2 flex-col gap-2 overflow-x-hidden overflow-y-auto px-1 pb-3"
							cdkDropList
							[id]="column.id"
							[cdkDropListData]="column.tasks"
							(cdkDropListDropped)="drop($event)"
						>
							@for (task of column.tasks; track task.id) {
								<lu-task-card
									cdkDrag
									[projectCode]="project.code"
									[task]="task"
									[cdkDragData]="task"
									(openTaskDetail)="handleOpenTaskModal($event)"
								/>
							}
						</div>

						<lu-task-form-sm [project]="project" [column]="column" />
					</div>
				}
				<lu-create-status-form [projectId]="project.id" />
			</div>
		}
	`,
})
export class BoardViewComponent implements OnInit {
	@Input() project: Project
	private el: ElementRef = inject(ElementRef)
	private renderer: Renderer2 = inject(Renderer2)
	private store = inject(Store<AppState>)
	private router = inject(Router)

	columns$: Observable<StatusColumn[]>

	ngOnInit(): void {
		setTimeout(() => this.updateHeight(), 100)

		const status$ = this.store.select(selectProjectStatusList(this.project.id))
		const tasks$ = this.store.select(selectTasksByProjectID(this.project.id))

		this.columns$ = combineLatest([status$, tasks$]).pipe(
			map(([statusList, tasks]) => {
				return statusList.map(status => ({
					id: status.id.toString(),
					title: status.title,
					status: status,
					tasks: tasks.filter(t => t.status_id === status.id).sort((a, b) => a.list_index - b.list_index),
				}))
			})
		)
	}

	drop(event: CdkDragDrop<Task[]>) {
		const container = event.container.data
		const currentIdx = event.currentIndex

		if (event.previousContainer === event.container) {
			moveItemInArray(container, event.previousIndex, currentIdx)
		} else {
			transferArrayItem(event.previousContainer.data, container, event.previousIndex, currentIdx)
		}

		const prev = container[currentIdx - 1]
		const next = container[currentIdx + 1]

		let newIndex: number

		if (!prev && !next) {
			newIndex = 10000 // single item in column
		} else if (!prev) {
			newIndex = next.list_index / 2 // if no previous, it's a first item.
		} else if (!next) {
			newIndex = prev.list_index + 10000 // no next item, it's a last item
		} else {
			newIndex = (prev.list_index + next.list_index) / 2 // in the middle of between two items
		}

		this.store.dispatch(
			TaskChangeStatusAction({
				taskId: event.item.data.id,
				payload: {
					from_status: parseInt(event.previousContainer.id),
					to_status: parseInt(event.container.id),
					from_idx: event.previousIndex,
					to_idx: newIndex,
				},
			})
		)
	}

	handleOpenTaskModal(taskCode: string): void {
		this.router.navigate(["/projects", this.project.code, "board", taskCode])
	}

	@HostListener("window:resize")
	onResize() {
		this.updateHeight()
	}

	private updateHeight() {
		const elementTop = this.el.nativeElement.getBoundingClientRect().top
		const availableHeight = window.innerHeight - elementTop
		this.renderer.setStyle(this.el.nativeElement, "height", `${availableHeight - 16}px`)
	}
}
