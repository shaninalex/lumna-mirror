import {Component, inject, Input, OnDestroy, OnInit} from '@angular/core';
import {Task, IssueCardComponent, TaskService} from '@client/entities/task';
import {
    CdkDrag,
    CdkDragDrop,
    CdkDropList,
    CdkDropListGroup,
    moveItemInArray,
    transferArrayItem,
} from '@angular/cdk/drag-drop';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {Subscription, tap} from 'rxjs';
import {BoardViewApiService} from './board-view-api.service';
import {StatusColumn} from './board.model';
import {MatCardModule} from '@angular/material/card';

@Component({
    selector: "ts-board-view",
    imports: [
        CdkDropList,
        CdkDrag,
        MatProgressSpinnerModule,
        CdkDropListGroup,
        MatCardModule,
        IssueCardComponent,
    ],
    providers: [BoardViewApiService],
    styleUrl: './board-view.component.scss',
    templateUrl: './board-view.component.html'
})
export class BoardViewComponent implements OnInit, OnDestroy {
    @Input() projectKey: string;
    private _boardApi = inject(BoardViewApiService)
    private _sub = new Subscription()

    columns: StatusColumn[]

    drop(event: CdkDragDrop<Task[]>) {
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
        this._boardApi.ChangeStatus(event.item.data.id, event.previousContainer.id, event.container.id)
    }

    ngOnInit() {
        this._sub.add(
            this._boardApi.Statuses(this.projectKey).pipe(
                tap(statuses => {
                    this.columns = statuses.map(s => ({
                        id: s.id,
                        title: s.title,
                        tasks: s.tasks
                    }));
                })
            ).subscribe()
        )
    }

    ngOnDestroy() {
        this._sub.unsubscribe()
    }
}
