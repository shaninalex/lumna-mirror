import {Component, inject} from '@angular/core';
import {UiService} from '@client/shared/ui/ui.service';
import { ToastMessageComponent } from './toast-message.component';
import {ToastMessage} from '@client/shared/models';
import {AsyncPipe, JsonPipe} from '@angular/common';
import {Observable} from 'rxjs';

@Component({
    selector: "lu-toast-manager",
    imports: [
        ToastMessageComponent,
        AsyncPipe
    ],
    template: `
        @if(messages$ | async; as messages ) {
            <div class="fixed bottom-4 right-4 flex flex-col gap-2">
                @for (message of messages; track message.id) {
                    <lu-toast-message [toast]="message" (onClose)="handleOnClose($event)" />
                }
            </div>
        }
    `
})
export class ToastManagerComponent {
    private ui: UiService = inject(UiService);
    messages$: Observable<ToastMessage[]> = this.ui.messages.list()

    handleOnClose(id: string): void {
        this.ui.messages.remove(id);
    }
}
