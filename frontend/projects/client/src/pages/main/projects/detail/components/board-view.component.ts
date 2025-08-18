import {Component} from '@angular/core';
import {IssueCardComponent} from '@client/entities/issue';
import {
    CdkDragDrop,
    moveItemInArray,
    transferArrayItem,
    CdkDrag,
    CdkDropList,
} from '@angular/cdk/drag-drop';

@Component({
    selector: "ts-board-view",
    imports: [
        IssueCardComponent,
        CdkDropList,
        CdkDrag,
    ],
    styleUrl: 'exmple-board-view.component.scss',
    template: `
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
                    class="flex flex-col gap-2"
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
                    class="flex flex-col gap-2"
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
                    class="flex flex-col gap-2"
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
export class BoardViewComponent {
    todo = [
        '1',
        '2',
        '3',
    ];

    progress = [
        '7',
        '6',
        '5',
        '4',
    ];

    done = [
        '8',
        '9',
        '10',
    ];

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
