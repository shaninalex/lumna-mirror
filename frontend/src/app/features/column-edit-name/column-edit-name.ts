import { ChangeDetectionStrategy, Component, effect, inject, input, signal } from '@angular/core';
import {
    actionColumnFailed,
    actionColumnPatch,
    actionColumnUpsert,
    ColumnState,
    ColumnModel,
} from '@entities/column';
import { FormField, form, required } from '@angular/forms/signals';
import { ClickOutsideDirective } from '@shared/directives';
import { Actions, ofType } from '@ngrx/effects';
import { tap } from 'rxjs';
import { Store } from '@ngrx/store';

@Component({
    selector: 'app-column-edit-name-feature',
    imports: [FormField, ClickOutsideDirective],
    template: `
        @if (formOpen()) {
            <form (submit)="submit($event)" class="w-100" (clickOutside)="formOpen.set(false)">
                <div class="input-group">
                    <input type="text" class="form-control form-control-sm fw-bold" placeholder="Column name" autocomplete="off" [formField]="listForm.title">
                    <button class="btn btn-sm btn-outline-primary" type="submit">
                        @if (loading()) {
                            <i class="fa-solid fa-spinner spin"></i>
                        } @else {
                            Save
                        }
                    </button>
                </div>
                @if (listForm.title().touched() && listForm.title().errors().length) {
                    @for (error of listForm.title().errors(); track error) {
                        <div class="text-danger small">{{ error.message }}</div>
                    }
                }
            </form>
        } @else {
            <div class="fw-bold small" (click)="formOpen.set(true)">{{ column().title }}</div>
        }
    `,
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ColumnEditNameFeature {
    private actions$ = inject(Actions);
    private store = inject(Store<ColumnState>);

    column = input.required<ColumnModel>();
    formOpen = signal(false);
    loading = signal(false);
    columnFormModel = signal<{ title: string }>({ title: '' });

    listForm = form(this.columnFormModel, (schemaPath) => {
        required(schemaPath.title, { message: 'Name is required' });
    });

    constructor() {
        effect(() =>
            this.columnFormModel.set({
                title: this.column().title,
            }),
        );

        this.actions$
            .pipe(
                ofType(actionColumnFailed),
                tap(() => this.loading.set(false)),
            )
            .subscribe();

        this.actions$
            .pipe(
                ofType(actionColumnUpsert),
                tap(() => {
                    this.loading.set(false);
                    this.formOpen.set(false);
                }),
            )
            .subscribe();
    }

    submit(event: Event): void {
        event.preventDefault();
        this.loading.set(true);
        const data = this.columnFormModel();

        this.store.dispatch(
            actionColumnPatch({
                columnId: this.column().id,
                data: {
                    title: data.title,
                    order: this.column().order,
                    board_id: this.column().board_id,
                },
            }),
        );
    }
}
