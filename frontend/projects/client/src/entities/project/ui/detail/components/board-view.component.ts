import {Component, inject, Input, OnInit} from '@angular/core';
import {IssueCardComponent, TaskService} from '@client/entities/issue';
import {CdkDrag, CdkDragDrop, CdkDropList, moveItemInArray, transferArrayItem,} from '@angular/cdk/drag-drop';
import {AsyncPipe, JsonPipe} from '@angular/common';

@Component({
    selector: "ts-board-view",
    imports: [
        IssueCardComponent,
        CdkDropList,
        CdkDrag,
        AsyncPipe,
        JsonPipe,
    ],
    styleUrl: 'example-board-view.component.scss',
    template: `
        {{ tasks | async | json }}

        <div class="flex flex-row no-wrap gap-4">

            <div class="bg-slate-200 w-sm rounded p-4">
                <div class="mb-2">
                    <div class="font-bold text-slate-500">TODO</div>
                </div>
                <div
                    cdkDropList
                    #todoList="cdkDropList"
                    [cdkDropListData]="todo"
                    [cdkDropListConnectedTo]="[progressList, doneList]"
                    class="flex flex-col gap-2 min-h-20"
                    (cdkDropListDropped)="drop($event)">
                    @for (item of todo; track item) {
                        <ts-issue-card cdkDrag [id]="item"/>
                    }
                </div>
            </div>

            <div class="bg-slate-200 w-sm rounded p-4">
                <div class="mb-2">
                    <div class="font-bold text-slate-500">IN PROGRESS</div>
                </div>
                <div
                    cdkDropList
                    #progressList="cdkDropList"
                    [cdkDropListData]="progress"
                    [cdkDropListConnectedTo]="[todoList, doneList]"
                    class="flex flex-col gap-2 min-h-20"
                    (cdkDropListDropped)="drop($event)">
                    @for (item of progress; track item) {
                        <ts-issue-card cdkDrag [id]="item"/>
                    }
                </div>
            </div>


            <div class="bg-slate-200 w-sm rounded p-4">
                <div class="mb-2">
                    <div class="font-bold text-slate-500">DONE</div>
                </div>
                <div
                    cdkDropList
                    #doneList="cdkDropList"
                    [cdkDropListData]="done"
                    [cdkDropListConnectedTo]="[todoList, progressList]"
                    class="flex flex-col gap-2 min-h-20"
                    (cdkDropListDropped)="drop($event)">
                    @for (item of done; track item) {
                        <div class="example-box" cdkDrag>
                            <ts-issue-card [id]="item"/>
                        </div>
                    }
                </div>
            </div>
        </div>
    `
})
export class BoardViewComponent implements OnInit {
    @Input() projectKey: string;
    api = inject(TaskService)
    tasks: any

    todo = ['1', '2', '3'];
    progress = ['7', '6', '5', '4'];
    done = ['8', '9', '10'];

    ngOnInit() {
        // TODO: rewrite with rxjs Observables
        this.tasks = this.api.List(this.projectKey)
    }

    drop(event: CdkDragDrop<string[]>) {
        if (event.previousContainer === event.container) {
            moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
        } else {
            transferArrayItem(
                event.previousContainer.data,
                event.container.data,
                event.previousIndex,
                event.currentIndex,
            );
        }
        console.log("Drop completed", event)
    }
}
