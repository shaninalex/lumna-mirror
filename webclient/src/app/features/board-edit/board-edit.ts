import { Component, effect, inject, input, signal } from '@angular/core';
import { BoardModel, BoardPayloadModel } from '@entities/board';
import { Field, form, required } from '@angular/forms/signals';
import { Dispatcher, Events } from '@ngrx/signals/events';
import { projectEvents } from '@entities/project';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { boardEvents } from '@entities/board/model/board.events';

@Component({
    selector: 'app-board-edit-feature',
    imports: [Field],
    template: `
        <form (submit)="submit($event)">
            <div class="mb-2">
                <input class="input" placeholder="Project name" [field]="boardForm.name" />
                @if (boardForm.name().touched()) {
                    <ul class="error-list">
                        @for (error of boardForm.name().errors(); track error) {
                            <li class="text-red-500 text-sm">{{ error.message }}</li>
                        }
                    </ul>
                }
            </div>

            @if (errors()) {
                <ul class="mb-2">
                    @for (error of errors(); track error) {
                        <li class="text-red-500 text-sm">{{ error }}</li>
                    }
                </ul>
            }

            <div>
                <button type="submit" class="btn btn-primary" [disabled]="boardForm().invalid()">
                    @if (loading()) {
                        Processing...
                    } @else {
                        Save
                    }
                </button>
            </div>
        </form>
    `,
})
export class BoardEditFeature {
    board = input<BoardModel>();
    boardId: string;
    private readonly dispatcher = inject(Dispatcher);
    private readonly events = inject(Events);
    boardFormModel = signal<BoardPayloadModel>({
        project_id: '',
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
                    project_id: b.project_id,
                    title: b.title,
                });
            }
        });

        this.events
            .on(boardEvents.failed)
            .pipe(takeUntilDestroyed())
            .subscribe((data) => {
                this.errors.set([data.payload.toString()]);
                this.loading.set(false);
            });

        this.events
            .on(boardEvents.set)
            .pipe(takeUntilDestroyed())
            .subscribe(() => this.loading.set(false));
    }

    submit(event: Event): void {
        event.preventDefault();
        this.loading.set(true);
        const formData = this.boardFormModel();
        this.dispatcher.dispatch(boardEvents.patch({ boardId: this.boardId, data: formData }));
    }
}
