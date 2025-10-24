import {Component, inject, Input, OnInit} from '@angular/core';
import {CdkDrag, CdkDragDrop, CdkDropList, moveItemInArray} from '@angular/cdk/drag-drop';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {Project, ProjectPatchAction} from '@client/entities/project';
import {selectProjectStatusList, Status, StatusDeleteAction, StatusPatchSortAction} from '@client/entities/status';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';

@Component({
    selector: "lu-project-settings-feature",
    imports: [
        CdkDrag,
        CdkDropList,
        ReactiveFormsModule,
    ],
    styleUrl: './project-settings-feature.component.scss',
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
                    <div cdkDrag class="status-sort-item flex items-center justify-between">
                        <span class="block">{{ status.title }}</span>
                        <button (click)="onStatusDelete(status.id)"
                                title="Delete column {{ status.title }}?"
                                class="text-xl text-red-500 cursor-pointer">
                            <i class="i-close-circle"></i>
                        </button>
                    </div>
                }
            </div>
            <div class="flex gap-2">
                <button (click)="saveOrder()" class="btn btn-primary">Save</button>
            </div>
        </div>
    `
})
export class ProjectSettingsFeatureComponent implements OnInit {
    private store = inject(Store<AppState>)

    @Input() project: Project
    statusList: Status[]

    form: FormGroup = new FormGroup({
        title: new FormControl('', Validators.required),
    })

    ngOnInit() {
        this.store.select(selectProjectStatusList(this.project.id)).subscribe(data => {
            this.statusList = data
        })
        this.form.setValue({title: this.project.title})
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

    onStatusDelete(id: number): void {
        this.store.dispatch(StatusDeleteAction({
            projectId: this.project.id,
            statusId: id,
        }))
    }
}
