import {Component, inject, Input, signal} from '@angular/core';
import {Dialog, DIALOG_DATA, DialogRef} from '@angular/cdk/dialog';
import {BoardPayloadModel} from '@entities/board';
import {Dispatcher} from '@ngrx/signals/events';
import {Field, form, required, validate} from '@angular/forms/signals';
import {boardEvents} from '@entities/board/model/board.events';

@Component({
    selector: 'app-board-create-feature',
    imports: [],
    template: `
        <button class="btn btn-secondary" (click)="newBoard()">Create board</button>
    `,
})
export class BoardCreateFeature {
    @Input() projectId: number;

    dialog = inject(Dialog);
    dispatcher = inject(Dispatcher);

    newBoard(): void {
        const dialogRef = this.dialog.open<BoardPayloadModel>(BoardForm, {
            width: '250px',
            data: {
                project_id: this.projectId,
            }
        });

        dialogRef.closed.subscribe(result => {
            if (result) this.dispatcher.dispatch(boardEvents.create(result))
        });
    }
}

@Component({
    selector: 'app-board-create-feature-form',
    imports: [Field],
    template: `
        <h3 class="font-medium">Create new board</h3>
        <form (submit)="submit($event)" class="flex flex-col gap-2 ">
            <div>
                <input class="border block w-full rounded-lg px-2 py-1 border-gray-400"
                       type="text" placeholder="Board name" [field]="boardForm.name" />
                @if (boardForm.name().touched() && boardForm.name().errors()) {
                    <ul class="error-list">
                        @for (error of boardForm.name().errors(); track error) {
                            <li class="text-red-500 text-sm">{{ error.message }}</li>
                        }
                    </ul>
                }
            </div>

            <div class="flex items-center gap-4">
                <button
                    type="submit"
                    class="bg-teal-500 text-white rounded-lg px-2 py-1 cursor-pointer text-sm disabled:opacity-50"
                    [disabled]="boardForm().invalid()"
                >
                    Create
                </button>

                <button type="button" (pointerdown)="cancel()"
                        class="bg-gray-200 text-white rounded-lg px-2 py-1 cursor-pointer text-sm">
                    Cancel
                </button>
            </div>
        </form>

    `,
    host: { class: 'modal'}
})
export class BoardForm {
    dialogRef = inject<DialogRef<BoardPayloadModel>>(DialogRef<BoardPayloadModel>);
    data = inject(DIALOG_DATA);

    boardFormModel = signal<BoardPayloadModel>({
        project_id: this.data.project_id,
        name: '',
    });

    boardForm = form(this.boardFormModel, (schemaPath) => {
        required(schemaPath.name, {message: 'Name is required'});
        validate(schemaPath.name, ({value}) => {
            if (value().trim().length === 0) {
                return {
                    kind: 'string',
                    message: 'Name should not be empty string'
                }
            }
            return null
        })
    });

    submit(event: Event): void {
        event.preventDefault()
        this.dialogRef.close(this.boardFormModel())
    }

    cancel(): void {
        this.dialogRef.close()
    }
}
