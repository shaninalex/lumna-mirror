import {Component, inject, input, signal} from '@angular/core';
import {BoardModel} from '@entities/board';
import {Field, form, required} from '@angular/forms/signals';
import {Dispatcher, Events} from '@ngrx/signals/events';
import {ListPayloadModel} from '@entities/list';
import {listEvents} from '@entities/list/model/list.events';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';

@Component({
    selector: 'app-new-column-form',
    imports: [Field],
    template: `
        @if (openedForm()) {
            <form (submit)="submit($event)" class="w-[280px]">
                <div class="mb-2">
                    <input class="input w-full" placeholder="Column name" [field]="listForm.name" />
                    @if (listForm.name().touched() && listForm.name().errors().length) {
                        <ul class="error-list">
                            @for (error of listForm.name().errors(); track error) {
                                <li class="text-red-500 text-sm">{{ error.message }}</li>
                            }
                        </ul>
                    }
                </div>

                <div class="flex gap-2">
                    <button class="btn btn-sm btn-primary" (click)="submit($event)" [disabled]="listForm().invalid()">
                        @if (loading()) {
                            Processing...
                        } @else {
                            Create
                        }
                    </button>
                    <button class="btn btn-sm btn-secondary" (click)="this.openedForm.set(false)">cancel</button>
                </div>
            </form>
        } @else {
            <button class="btn btn-secondary" (click)="this.openedForm.set(true)">Create new column</button>
        }
    `,
    host: {class: 'flex-shrink-0'}
})
export class NewColumnForm {
    readonly dispatcher = inject(Dispatcher);
    readonly events = inject(Events);
    board = input<BoardModel>()

    openedForm = signal<boolean>(false)
    loading = signal(false)
    errors = signal<string[]>([]);
    listFormModel = signal<ListPayloadModel>({
        name: '',
        order: 0,
    });

    listForm = form(this.listFormModel, (schemaPath) => {
        required(schemaPath.name, { message: 'Name is required' });
    });

    constructor() {
        this.events.on(listEvents.setList)
            .pipe(takeUntilDestroyed())
            .subscribe(() => {
                this.loading.set(false)
                this.openedForm.set(false)
            })

        this.events.on(listEvents.failed)
            .pipe(takeUntilDestroyed())
            .subscribe(data => {
                this.errors.set([data.payload.toString()])
                this.loading.set(false)
            })
    }

    submit(event: Event) {
        event.preventDefault()
        const formData = this.listFormModel()
        const b = this.board()

        if (!formData.name) return
        if (!b) return

        this.dispatcher.dispatch(listEvents.create({
            boardId: b.id,
            data: formData
        }))
    }
}
