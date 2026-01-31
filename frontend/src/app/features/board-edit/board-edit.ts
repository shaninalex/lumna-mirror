import { Component, effect, inject, input, signal } from '@angular/core';
import {
    actionBoardFailed,
    actionBoardPatch,
    actionBoardUpsert,
    BoardModel,
    BoardPayloadModel,
    BoardState,
} from '@entities/board';
import { FormField, form, required } from '@angular/forms/signals';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { Actions, ofType } from '@ngrx/effects';
import { tap } from 'rxjs';
import { Store } from '@ngrx/store';

@Component({
    selector: 'app-board-edit-feature',
    imports: [FormField],
    templateUrl: './board-edit.html',
})
export class BoardEditFeature {
    board = input<BoardModel>();
    boardId: string;
    private actions$ = inject(Actions);
    private store = inject(Store<BoardState>);

    boardFormModel = signal<BoardPayloadModel>({
        projectID: '',
        title: '',
    });
    boardForm = form(this.boardFormModel, (schemaPath) => {
        required(schemaPath.title, { message: 'Name is required' });
    });
    loading = signal(false);
    errors = signal<string[]>([]);

    constructor() {
        effect(() => {
            const b = this.board();
            if (b) {
                this.boardId = b.id;
                this.boardFormModel.set({
                    projectID: b.projectID,
                    title: b.title,
                });
            }
        });

        this.actions$
            .pipe(
                ofType(actionBoardFailed),
                takeUntilDestroyed(),
                tap((data) => {
                    this.errors.set([data.error.toString()]);
                    this.loading.set(false);
                }),
            )
            .subscribe();

        this.actions$
            .pipe(
                ofType(actionBoardUpsert),
                takeUntilDestroyed(),
                tap(() => this.loading.set(false)),
            )
            .subscribe();
    }

    submit(event: Event): void {
        event.preventDefault();
        this.loading.set(true);
        const formData = this.boardFormModel();
        this.store.dispatch(actionBoardPatch({ boardId: this.boardId, data: formData }));
    }
}
