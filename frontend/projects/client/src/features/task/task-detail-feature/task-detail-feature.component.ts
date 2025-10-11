import {Component, inject, Input, OnDestroy, OnInit} from '@angular/core';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {Task, TaskDeleteAction, TaskDetailInput, TaskPatchAction} from '@client/entities/task';
import {FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {DatePipe} from '@angular/common';
import {MoveTaskComponent} from './move-task';
import { NgxEditorComponent, NgxEditorMenuComponent, Editor } from 'ngx-editor';
import {UiService} from '@client/shared/ui/ui.service';

@Component({
    selector: "lu-task-detail-feature",
    imports: [
        DatePipe,
        FormsModule,
        ReactiveFormsModule,
        MoveTaskComponent,
        NgxEditorComponent,
        NgxEditorMenuComponent,
    ],
    template: `
        <div class="mb-4">
            <lu-move-task [task]="task"/>
        </div>

        <form [formGroup]="form" (ngSubmit)="onSubmit()">
            <div class="flex justify-between mb-4">
                <button [disabled]="!form.valid" class="btn btn-primary" type="submit">Save</button>
                <button (click)="onDelete()" class="btn btn-danger" type="button">Delete</button>
            </div>

            <div class="mb-4 card-title">
                <input class="input w-full" formControlName="title"/>
                <div class="text-xs">Created: {{ task.created_at | date }}</div>
            </div>

            <div class="mb-4">
                <div class="font-bold text-sm">Description</div>
                <div class="NgxEditor__Wrapper">
                    <ngx-editor-menu [editor]="editor"> </ngx-editor-menu>
                    <ngx-editor [editor]="editor" formControlName="description"></ngx-editor>
                </div>
            </div>
        </form>
    `
})
export class TaskDetailFeatureComponent implements OnInit, OnDestroy {
    @Input() task: Task
    private store = inject(Store<AppState>);
    private ui: UiService = inject(UiService);

    theme: string;
    editor: Editor;

    form: FormGroup = new FormGroup({
        title: new FormControl('', Validators.required),
        description: new FormControl('', Validators.required),
    })

    ngOnInit() {
        this.theme = this.ui.theme.appTheme()
        this.editor = new Editor();
        this.form.setValue({
            title: this.task.title,
            description: this.task.description,
        })
    }

    onSubmit(): void {
        const payload: TaskDetailInput = {
            title: this.form.value["title"],
            completed: this.task.completed,
            description: this.form.value["description"],
            list_index: this.task.list_index,
            status_id: this.task.status_id,
        }
        this.store.dispatch(TaskPatchAction({taskId: this.task.id, payload}))
    }

    onDelete(): void {
        this.store.dispatch(TaskDeleteAction({taskId: this.task.id}))
    }

    ngOnDestroy(): void {
        this.editor.destroy();
    }
}
