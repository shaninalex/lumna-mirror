import { Component, inject, Input, input, OnInit, signal } from '@angular/core';
import { FormField, form, required } from '@angular/forms/signals';
import {
    actionColumnCreate,
    actionColumnFailed,
    actionColumnSetList,
    ColumnPayloadModel,
    ColumnState,
} from '@entities/column';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { Actions, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { tap } from 'rxjs';

@Component({
    selector: 'app-new-column-form',
    imports: [FormField],
    template: `
        @if (openedForm()) {
            <form (submit)="submit($event)" class="w-[280px]">
                <div class="mb-2">
                    <input
                        class="input w-full"
                        placeholder="Column name"
                        [formField]="columnForm.title"
                    />
                    @if (columnForm.title().touched() && columnForm.title().errors().length) {
                        <ul class="error-list">
                            @for (error of columnForm.title().errors(); track error) {
                                <li class="text-red-500 text-sm">{{ error.message }}</li>
                            }
                        </ul>
                    }
                </div>

                <div class="flex gap-2">
                    <button
                        class="btn btn-sm btn-primary"
                        (click)="submit($event)"
                        [disabled]="columnForm().invalid()"
                    >
                        @if (loading()) {
                            Processing...
                        } @else {
                            Create
                        }
                    </button>
                    <button class="btn btn-sm btn-secondary" (click)="this.openedForm.set(false)">
                        cancel
                    </button>
                </div>
            </form>
        } @else {
            <button class="btn btn-secondary" (click)="this.openedForm.set(true)">
                Create new column
            </button>
        }
    `,
    host: { class: 'flex-shrink-0' },
})
export class NewColumnForm implements OnInit {
    private actions$ = inject(Actions);
    private store = inject(Store<ColumnState>);

    @Input() boardId: string;
    columns_length = input.required<number>();

    openedForm = signal<boolean>(false);
    loading = signal(false);
    errors = signal<string[]>([]);
    columnFormModel = signal<ColumnPayloadModel>({
        title: '',
        order: 0,
        board_id: '',
    });

    columnForm = form(this.columnFormModel, (schemaPath) => {
        required(schemaPath.title, { message: 'Name is required' });
    });

    ngOnInit() {
        this.actions$
            .pipe(
                ofType(actionColumnSetList),
                tap(() => {
                    this.loading.set(false);
                    this.openedForm.set(false);
                    this.columnForm().value.set({
                        title: '',
                        order: 0,
                        board_id: '',
                    });
                }),
            )
            .subscribe();

        this.actions$
            .pipe(
                ofType(actionColumnFailed),
                tap((data) => {
                    this.errors.set([data.error.toString()]);
                    this.loading.set(false);
                }),
            )
            .subscribe();
    }

    submit(event: Event) {
        event.preventDefault();
        const formData = this.columnFormModel();

        if (!formData.title) return;

        formData.order = this.columns_length() ? this.columns_length() : 0;
        formData.board_id = this.boardId;
        this.store.dispatch(actionColumnCreate({ data: formData }));
    }
}
