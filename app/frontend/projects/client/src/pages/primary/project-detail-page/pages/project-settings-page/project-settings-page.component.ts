import {Component, inject, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {Project, ProjectPatchAction} from '@client/entities/project';
import {AppState} from '@client/shared/store';
import {Store} from '@ngrx/store';
import {StatusPatchSortAction, selectProjectStatusList, Status} from '@client/entities/status';
import {filter, map, take, tap} from 'rxjs';
import {CdkDragDrop, CdkDropList, CdkDrag, moveItemInArray} from '@angular/cdk/drag-drop';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {NgClass} from '@angular/common';


@Component({
    selector: 'lu-project-settings-page',
    imports: [
        CdkDropList,
        CdkDrag,
        ReactiveFormsModule,
        NgClass,
    ],
    styleUrl: './project-settings-page.component.scss',
    template: `
        <div class="card mb-4">
            <div class="card-title">Project name:</div>
            <form [formGroup]="form" (ngSubmit)="onSubmitTitleForm()">
                <div class="mb-4">
                    <input class="input" type="text" formControlName="title">
                </div>
                <div class="flex gap-2">
                    <button [disabled]="!form.valid" type="submit" class="btn btn-primary">Save</button>
                    <button (click)="form.reset()" class="btn btn-secondary">Reset</button>
                </div>
            </form>
        </div>
        <div class="card">
            <div class="card-title">Status order:</div>
            <div cdkDropList class="status-sort-list mb-4" (cdkDropListDropped)="drop($event)">
                @for (status of statusList; track status.id) {
                    <div cdkDrag class="status-sort-item" [ngClass]="{'accent': status.complete}">
                        {{ status.title }}
                    </div>
                }
            </div>
            <div class="flex gap-2">
                <button (click)="saveOrder()" class="btn btn-primary">Save</button>
            </div>
        </div>
    `
})
export class ProjectSettingsPageComponent implements OnInit {
    private store = inject(Store<AppState>)
    private route = inject(ActivatedRoute);
    project: Project
    statusList: Status[]
    form: FormGroup = new FormGroup({
        title: new FormControl('', Validators.required),
    })

    ngOnInit() {
        this.route.data.pipe(
            take(1),
            map(data => data['project'] as Project),
            filter(project => !!project),
            tap(project => {
                this.project = project
                this.store.select(selectProjectStatusList(project.id)).subscribe(data => {
                    this.statusList = data
                })
                this.form.setValue({title: project.title})
            }),
        ).subscribe()
    }

    drop(event: CdkDragDrop<string[]>) {
        moveItemInArray(this.statusList, event.previousIndex, event.currentIndex);
    }

    saveOrder(): void {
        const order: Record<number, number> = {}
        for (let i = 0; i < this.statusList.length; i++) {
            order[i+1] = this.statusList[i].id
        }
        this.store.dispatch(StatusPatchSortAction({ projectId: this.project.id, payload: order }))
    }

    onSubmitTitleForm(): void {
        if (!this.form.valid) return
        this.store.dispatch(ProjectPatchAction({
            projectId: this.project.id,
            payload: {
                title: this.form.value['title'],
            }
        }))
    }
}
