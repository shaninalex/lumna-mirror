import { ChangeDetectionStrategy, Component, effect, inject, input, signal } from '@angular/core';
import { ListModel, ListPayloadModel } from '@entities/list';
import { Field, form, required } from '@angular/forms/signals';
import { Dispatcher, Events } from '@ngrx/signals/events';
import { listEvents } from '@entities/list/model/list.events';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { ClickOutsideDirective } from '@shared/directives';

@Component({
    selector: 'app-list-edit-name-feature',
    imports: [Field, ClickOutsideDirective],
    template: `
        @if (formOpen()) {
            <form (submit)="submit($event)" class="flex gap-2" (clickOutside)="formOpen.set(false)">
                <div>
                    <input class="input" placeholder="Project name" [field]="listForm.title" />
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
                        class="btn btn-primary btn-sm"
                        [disabled]="listForm().invalid()"
                    >
                        @if (loading()) {
                            Processing...
                        } @else {
                            Save
                        }
                    </button>
                </div>
            </form>
        } @else {
            <div class="text-lg font-medium" (click)="formOpen.set(true)">
                {{ list().title }} [{{ list().id }}]
            </div>
        }
    `,
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ListEditNameFeature {
    private readonly dispatcher = inject(Dispatcher);
    private readonly events = inject(Events);

    list = input.required<ListModel>();
    formOpen = signal(false);
    loading = signal(false);
    listFormModel = signal<ListPayloadModel>({
        title: '',
        order: 0,
    });

    listForm = form(this.listFormModel, (schemaPath) => {
        required(schemaPath.title, { message: 'Name is required' });
    });

    constructor() {
        effect(() =>
            this.listFormModel.set({
                title: this.list().title,
                order: this.list().order,
            }),
        );

        this.events
            .on(listEvents.failed)
            .pipe(takeUntilDestroyed())
            .subscribe(() => {
                this.loading.set(false);
            });

        this.events
            .on(listEvents._patchSuccess)
            .pipe(takeUntilDestroyed())
            .subscribe(() => {
                this.loading.set(false);
                this.formOpen.set(false);
            });
    }

    submit(event: Event): void {
        event.preventDefault();
        this.loading.set(true);
        this.dispatcher.dispatch(
            listEvents.patch({
                listId: this.list().id,
                data: this.listFormModel(),
            }),
        );
    }
}
