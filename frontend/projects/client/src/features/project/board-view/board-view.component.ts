import {Component, inject, Input, OnDestroy, OnInit} from '@angular/core';
import {Issue, IssueCardComponent, TaskService} from '@client/entities/issue';
import {CdkDrag, CdkDragDrop, CdkDropList, moveItemInArray, transferArrayItem,} from '@angular/cdk/drag-drop';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {Subscription, tap} from 'rxjs';
import {BoardViewApiService} from './board-view-api.service';
import {Status} from '@client/entities/project';
import {StatusColumn} from './board.model';

@Component({
    selector: "ts-board-view",
    imports: [
        IssueCardComponent,
        CdkDropList,
        CdkDrag,
        MatProgressSpinnerModule,
    ],
    providers: [BoardViewApiService],
    styleUrl: './board-view.component.scss',
    templateUrl: './board-view.component.html'
})
export class BoardViewComponent implements OnInit, OnDestroy {
    @Input() projectKey: string;
    private _taskApi = inject(TaskService)
    private _boardApi = inject(BoardViewApiService)
    private _sub = new Subscription()

    tasks: Issue[]
    statuses: Status[]
    columns: StatusColumn[]

    todo = ['1', '2', '3'];
    progress = ['7', '6', '5', '4'];
    done = ['8', '9', '10'];

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

    ngOnInit() {
        // TODO:
        //  - build columns
        //  - update template
        //  - update this.drop() function
        this._sub.add(
            this._taskApi.List(this.projectKey).pipe(
                tap(tasks => this.tasks = tasks),
            ).subscribe()
        )
        this._sub.add(
            this._boardApi.Statuses(this.projectKey).pipe(
                tap(statuses => this.statuses = statuses),
            ).subscribe(),
        )
    }

    ngOnDestroy() {
        this._sub.unsubscribe()
    }
}
