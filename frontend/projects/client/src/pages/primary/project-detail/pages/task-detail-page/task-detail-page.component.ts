import {Component, inject, OnDestroy, OnInit} from '@angular/core';
import {filter, map, Observable, tap} from 'rxjs';
import { Task } from "@client/entities/task";
import {ActivatedRoute} from '@angular/router';
import {AsyncPipe, DatePipe} from '@angular/common';
import {Editor, NgxEditorComponent, NgxEditorMenuComponent} from 'ngx-editor';
import {FormsModule} from '@angular/forms';

@Component({
    selector: "fr-task-detail-modal",
    imports: [
        AsyncPipe,
        DatePipe,
        NgxEditorMenuComponent,
        NgxEditorComponent,
        FormsModule
    ],
    template: `
        <div class="card">
            @if (task$ | async; as task) {
                <div class="mb-4">
                    <div class="card-title">{{ task.title }}</div>
                    <div class="text-xs">Created: {{ task.created_at | date }}</div>
                </div>

                <div>
                    <div class="font-bold text-sm">Description</div>
                    <div class="NgxEditor__Wrapper">
                        <ngx-editor-menu [editor]="editor"> </ngx-editor-menu>
                        <ngx-editor
                            [editor]="editor"
                            [ngModel]="html"
                            [disabled]="false"
                            [placeholder]="'Type here...'"
                        ></ngx-editor>
                    </div>
                </div>
            }
        </div>
    `
})
export class TaskDetailPageComponent implements OnInit, OnDestroy  {
    editor: Editor;
    private route = inject(ActivatedRoute)
    html: string = '';
    task$: Observable<Task> = this.route.data.pipe(
        map(data => data["task"]),
        filter(task => !!task),
        tap(task => {
            this.html = task.description;
        })
    )

    ngOnInit(): void {
        this.editor = new Editor();
    }

    ngOnDestroy(): void {
        this.editor.destroy();
    }
}
