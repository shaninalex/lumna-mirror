import {Component, inject, Input, OnDestroy, OnInit} from '@angular/core';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {Task, TaskDeleteAction, TaskDetailInput, TaskPatchAction} from '@client/entities/task';
import {FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {DatePipe} from '@angular/common';
import {MoveTaskComponent} from './move-task';
import { NgxEditorComponent, NgxEditorMenuComponent, Editor } from 'ngx-editor';
import {UiService} from '@client/shared/ui/ui.service';
import {ActivatedRoute, Router} from '@angular/router';

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
        <div class="fixed cursor-pointer inset-0 bg-black/20" (click)="close()"></div>
        <div class="fixed -translate-x-2/4 top-0 z-10 left-2/4 w-3xl h-screen py-4 overflow-auto">
            <div class="card">
                <button (click)="close()" class="cursor-pointer"><i class="i-close-circle"></i></button>
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
            </div>
        </div>
    `
})
export class TaskDetailFeatureComponent implements OnInit, OnDestroy {
    @Input() task: Task

    private router = inject(Router);
    private route = inject(ActivatedRoute);

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
        this.close() // NOTE: will be better to wait success result of operations
    }

    onDelete(): void {
        this.store.dispatch(TaskDeleteAction({taskId: this.task.id}))
        this.close()
    }

    ngOnDestroy(): void {
        this.editor.destroy();
    }

    close() {
        this.router.navigate(['../'], { relativeTo: this.route });
    }
}
