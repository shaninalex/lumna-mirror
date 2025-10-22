import {Component, ElementRef, HostListener, inject, Input, OnInit} from '@angular/core';
import {Project} from '@client/entities/project';
import {StatusColumn} from '@client/features/project/board-view-feature/board.model';
import {FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {StatusPatchAction} from '@client/entities/status';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';

@Component({
    selector: 'lu-column-header',
    imports: [
        ReactiveFormsModule,
        FormsModule,
    ],
    template: `
        @if (!formsOpen) {
            <div class="text-slate-600 card-title mb-4 cursor-pointer" (click)="formsOpen = true">{{ column.title }}</div>
        } @else {
            <form [formGroup]="form" class="my-1">
                <input class="input" formControlName="title"/>
            </form>
        }
    `,
})
export class ColumnHeaderComponent implements OnInit {
    @Input() project: Project;
    @Input() column: StatusColumn;
    private store = inject(Store<AppState>);

    eRef = inject(ElementRef)
    formsOpen: boolean = false;
    form: FormGroup = new FormGroup({
        title: new FormControl('', Validators.required),
    })

    ngOnInit() {
        this.form.setValue({
            title: this.column.status.title,
        })
    }

    @HostListener('document:click', ['$event'])
    clickout(event: { target: any; }) {
        if (!this.eRef.nativeElement.contains(event.target) && this.formsOpen ) {
            this.onSubmit();
        }
    }

    onSubmit(): void {
        this.store.dispatch(StatusPatchAction({
            payload: {
                title: this.form.value['title'],
                complete: this.form.value['complete'],
            },
            projectId: this.column.status.project_id,
            statusId: this.column.status.id,
        }))
        this.formsOpen = false;
    }
}
