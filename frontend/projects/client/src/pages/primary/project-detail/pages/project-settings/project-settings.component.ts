import {Component, inject, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {Project} from '@client/entities/project';
import {AppState} from '@client/shared/store';
import {Store} from '@ngrx/store';
import {selectProjectStatusList, Status} from '@client/entities/status';
import {filter, map, take, tap} from 'rxjs';
import {CdkDragDrop, CdkDropList, CdkDrag, moveItemInArray} from '@angular/cdk/drag-drop';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';


@Component({
    selector: 'fr-project-settings-page',
    imports: [
        CdkDropList,
        CdkDrag,
        ReactiveFormsModule,
    ],
    styleUrl: './project-settings.component.scss',
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
                    <div cdkDrag class="status-sort-item">
                        {{ status.title }}
                    </div>
                }
            </div>
            <div class="flex gap-2">
                <button (click)="saveOrder()" class="btn btn-primary">Save</button>
                <button (click)="cancelOrder()" class="btn btn-secondary">Reset</button>
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
        for (let i = 0; i < this.statusList.length; i++) {
            console.log(this.statusList[i].title)
        }
    }

    cancelOrder(): void {

    }

    onSubmitTitleForm(): void {}

}
