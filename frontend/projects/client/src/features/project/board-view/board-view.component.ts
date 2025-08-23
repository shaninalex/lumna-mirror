import {Component, inject, Input, OnInit} from '@angular/core';
import {Issue, IssueCardComponent, TaskService} from '@client/entities/issue';
import {CdkDrag, CdkDragDrop, CdkDropList, moveItemInArray, transferArrayItem,} from '@angular/cdk/drag-drop';
import {AsyncPipe} from '@angular/common';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {Observable, tap} from 'rxjs';
import {BoardViewApiService} from './board-view-api.service';

@Component({
    selector: "ts-board-view",
    imports: [
        IssueCardComponent,
        CdkDropList,
        CdkDrag,
        AsyncPipe,
        MatProgressSpinnerModule,
    ],
    providers: [BoardViewApiService],
    styleUrl: './board-view.component.scss',
    templateUrl: './board-view.component.html'
})
export class BoardViewComponent implements OnInit {
    @Input() projectKey: string;
    api = inject(TaskService)
    boardService = inject(BoardViewApiService)
    tasks: Observable<Issue[]>

    todo = ['1', '2', '3'];
    progress = ['7', '6', '5', '4'];
    done = ['8', '9', '10'];

    ngOnInit() {
        // TODO: rewrite with rxjs Observables
        this.tasks = this.api.List(this.projectKey)
        this.boardService.BoardView(this.projectKey).pipe(
            tap(data => console.log(data))
        ).subscribe()
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
