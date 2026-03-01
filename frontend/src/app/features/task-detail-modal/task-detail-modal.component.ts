import {Component, EventEmitter, inject, Input, OnInit, Output} from '@angular/core';
import {Store} from '@ngrx/store';
import {filter, Observable} from 'rxjs';
import {actionTaskChange, TaskModel} from '@entities/task';
import {AsyncPipe} from '@angular/common';

import {TaskBodyEditorComponent, TaskChangeStatusComponent, TaskTitleEditorComponent} from './components';
import {TaskActivityComponent} from '@features/task-detail-modal/components/task-activity/task-activity.component';


@Component({
    selector: 'task-detail-modal-feature',
    imports: [
        AsyncPipe,
        TaskBodyEditorComponent,
        TaskTitleEditorComponent,
        TaskChangeStatusComponent,
        TaskActivityComponent,
    ],
    templateUrl: './task-detail-modal.component.html'
})
export class TaskDetailModalFeature implements  OnInit {
    private store = inject(Store);
    @Input() task$: Observable<TaskModel>
    @Output() onClose: EventEmitter<any> = new EventEmitter();

    originalTask!: TaskModel;
    editedTask!: TaskModel;
    isChanged = false;

    ngOnInit() {
        this.task$
            .pipe(filter(Boolean))
            .subscribe(task => {
                this.originalTask = task;
                this.editedTask = structuredClone(task); // deep clone
                // this.checkChanges();
            });
    }

    save(): void {
        if (!this.isChanged) return;

        this.store.dispatch(
            actionTaskChange({
                task_id: this.originalTask.id,
                data: this.editedTask
            })
        );

        this.originalTask = structuredClone(this.editedTask);
        this.isChanged = false;
    }

    close(): void {
        this.onClose.emit();
    }

    onBackdropClick(event: MouseEvent): void {
        if (event.target === event.currentTarget) {
            this.onClose.emit();
        }
    }

    handleOnChangeBody(body: string): void {
        this.editedTask.body = body;
        this.checkChanges();
    }

    handleOnChangeTitle(title: string): void {
        this.editedTask.title = title;
        this.checkChanges();
    }

    private checkChanges(): void {
        this.isChanged = JSON.stringify(this.originalTask) !== JSON.stringify(this.editedTask);
        console.log(JSON.stringify(this.originalTask))
        console.log(JSON.stringify(this.editedTask))
    }
}
