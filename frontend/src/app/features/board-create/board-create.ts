import { Component, inject, Input, signal } from '@angular/core';
import { Dialog, DIALOG_DATA, DialogRef } from '@angular/cdk/dialog';
import { BoardPayloadModel } from '@entities/board';
import { Dispatcher } from '@ngrx/signals/events';
import { Field, form, required, validate } from '@angular/forms/signals';
import { boardEvents } from '@entities/board/model/board.events';

@Component({
    selector: 'app-board-create-feature',
    imports: [],
    template: ` <button class="btn btn-secondary" (click)="newBoard()">Create board</button> `,
})
export class BoardCreateFeature {
    @Input() projectId: string;

    dialog = inject(Dialog);
    dispatcher = inject(Dispatcher);

    newBoard(): void {
        const dialogRef = this.dialog.open<BoardPayloadModel>(BoardForm, {
            width: '250px',
            data: {
                project_id: this.projectId,
            },
        });

        dialogRef.closed.subscribe((result) => {
            if (result) this.dispatcher.dispatch(boardEvents.create(result));
        });
    }
}

@Component({
    selector: 'app-board-create-feature-form',
    imports: [Field],
    templateUrl: './board-create.html',
    host: { class: 'modal' },
})
export class BoardForm {
    dialogRef = inject<DialogRef<BoardPayloadModel>>(DialogRef<BoardPayloadModel>);
    data = inject(DIALOG_DATA);

    boardFormModel = signal<BoardPayloadModel>({
        project_id: this.data.project_id,
        title: '',
    });

    boardForm = form(this.boardFormModel, (schemaPath) => {
        required(schemaPath.title, { message: 'Title is required' });
        validate(schemaPath.title, ({ value }) => {
            if (value().trim().length === 0) {
                return {
                    kind: 'string',
                    message: 'Title should not be empty string',
                };
            }
            return null;
        });
    });

    submit(event: Event): void {
        event.preventDefault();
        this.dialogRef.close(this.boardFormModel());
    }

    cancel(): void {
        this.dialogRef.close();
    }
}
