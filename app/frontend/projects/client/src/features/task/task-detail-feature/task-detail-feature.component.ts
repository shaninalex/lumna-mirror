import { Component, inject, Input, OnDestroy, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { AppState } from '@client/shared/store';
import { Task, TaskDeleteAction, TaskDetailInput, TaskPatchAction } from '@client/entities/task';
import { FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { DatePipe } from '@angular/common';
import { TaskStatusDropdownComponent } from './task-status-dropdown';
import { NgxEditorComponent, NgxEditorMenuComponent, Editor } from 'ngx-editor';
import { UiService } from '@client/shared/ui/ui.service';
import { ActivatedRoute, Router } from '@angular/router';
import {CdkMenuModule} from '@angular/cdk/menu';
import {OverlayModule} from '@angular/cdk/overlay';

@Component({
    selector: "lu-task-detail-feature",
    imports: [
        DatePipe,
        FormsModule,
        ReactiveFormsModule,
        TaskStatusDropdownComponent,
        NgxEditorComponent,
        NgxEditorMenuComponent,
        CdkMenuModule,
        OverlayModule
    ],
    styleUrl: './task-detail-feature.component.scss',
    template: `
        <div class="fixed cursor-pointer inset-0 bg-black/20" (click)="close()"></div>
        <div class="task-detail-feature">
            <div class="card">
                <div class="flex items-center space-between p-4 border-b border-gray-300">
                    <lu-task-status-dropdown [task]="task"/>
                    <span class="flex-grow"></span>
                    <div class="flex items-center gap-2">
                        <button class="cursor-pointer"
                                cdkOverlayOrigin
                                #trigger="cdkOverlayOrigin"
                                (click)="isOpen = !isOpen">
                            <i class="i-dots-menu"></i>
                        </button>
                        <ng-template
                            cdkConnectedOverlay
                            [cdkConnectedOverlayOrigin]="trigger"
                            [cdkConnectedOverlayOpen]="isOpen"
                            [cdkConnectedOverlayHasBackdrop]="true"
                            (backdropClick)="isOpen = false"
                        >
                            <div class="card">
                                <button (click)="onDelete()" class="text-red-500 cursor-pointer" type="button">Delete</button>
                            </div>
                        </ng-template>
                        <button (click)="close()" class="cursor-pointer text-2xl"><i class="i-close-circle"></i>
                        </button>
                    </div>
                </div>

                <div class="grid grid-cols-3">
                    <div class="col-span-2 p-4">
                        <form [formGroup]="form" (ngSubmit)="onSubmit()">
                            <div class="mb-4 card-title flex-grow">
                                <input class="input w-full" formControlName="title"/>
                            </div>
                            <div class="mb-4">
                                <div class="font-bold text-sm">Description</div>
                                <div class="NgxEditor__Wrapper">
                                    <ngx-editor-menu [editor]="editor"></ngx-editor-menu>
                                    <ngx-editor [editor]="editor" formControlName="description"></ngx-editor>
                                </div>
                            </div>
                            <div class="flex justify-between mb-4">
                                <button [disabled]="!form.valid" class="btn btn-primary" type="submit">Save</button>
                            </div>
                        </form>
                    </div>
                    <div class="bg-gray-100 p-4 rounded-br-xl">
                        <div class="text-xs text-gray-500">Created: {{ task.created_at | date }}</div>
                        <div class="font-bold text-lg text-gray-600">Activity:</div>
                    </div>
                </div>
            </div>
        </div>`
})
export class TaskDetailFeatureComponent implements OnInit, OnDestroy {
    @Input() task: Task

    private router = inject(Router);
    private route = inject(ActivatedRoute);

    private store = inject(Store<AppState>);
    private ui: UiService = inject(UiService);

    theme: string;
    editor: Editor;
    isOpen = false;
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
        this.store.dispatch(TaskPatchAction({ taskId: this.task.id, payload }))
        this.close() // NOTE: will be better to wait success result of operations
    }

    onDelete(): void {
        this.store.dispatch(TaskDeleteAction({ taskId: this.task.id }))
        this.close()
    }

    ngOnDestroy(): void {
        this.editor.destroy();
    }

    close() {
        this.router.navigate(['../'], { relativeTo: this.route });
    }
}
