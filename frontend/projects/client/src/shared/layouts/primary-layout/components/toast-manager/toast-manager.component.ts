import {Component, inject} from '@angular/core';
import {Observable} from 'rxjs';
import {UiService} from '@client/shared/ui/ui.service';
import {AsyncPipe} from '@angular/common';
import { ToastMessageComponent } from './toast-message.component';
import {ToastMessage} from '@client/shared/models';

@Component({
    selector: "lu-toast-manager",
    imports: [
        AsyncPipe,
        ToastMessageComponent
    ],
    template: `
        @if (messages$ | async; as messages) {
            <div class="fixed bottom-4 right-4 flex flex-col gap-2">
                @for (message of messages; track $index) {
                    <lu-toast-message [toast]="message" (onClose)="handleOnClose($event)" />
                }
            </div>
        }
    `
})
export class ToastManagerComponent {
    private ui: UiService = inject(UiService);
    messages$: Observable<ToastMessage[]> = this.ui.messages$()

    handleOnClose(id: string): void {
        this.ui.removeMessage(id);
    }
}
