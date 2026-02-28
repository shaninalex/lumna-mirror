import { ChangeDetectionStrategy, Component, effect, inject, input, signal } from '@angular/core';
import {
    actionColumnFailed,
    actionColumnPatch,
    actionColumnUpsert,
    ColumnState,
    ColumnModel,
    ColumnPayloadModel,
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
            <form (submit)="submit($event)" class="is-flexgap-2" (clickOutside)="formOpen.set(false)">
                <div>
                    <input
                        class="input"
                        placeholder="Project name"
                        [formField]="listForm.title"
                        autocomplete="off"
                    />
                    @if (listForm.title().touched() && listForm.title().errors().length) {
                        <ul class="error-list">
                            @for (error of listForm.title().errors(); track error) {
                                <li class="text-red-500 text-sm">{{ error.message }}</li>
                            }
                        </ul>
                    }
                </div>

                <div>
                    <button
                        type="submit"
                        class="button is-primary is-small"
                        [disabled]="listForm().invalid()"
                    >
                        @if (loading()) {
                            <i class="fa-solid fa-spinner spin"></i>
                        } @else {
                            Save
                        }
                    </button>
                </div>
            </form>
        } @else {
            <div class="card-header-title" (click)="formOpen.set(true)">
                {{ column().title }}
            </div>
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
