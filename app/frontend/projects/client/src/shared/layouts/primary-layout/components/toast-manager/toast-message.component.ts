import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {ToastMessage} from '@client/shared/models';

@Component({
    selector: "lu-toast-message",
    template: `
        <div class="rounded p-2 bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 text-sm">
            {{ toast.message }}
            <button (click)="onClose.emit(this.toast.id)" class="cursor-pointer"><i class="i-close-circle"></i></button>
        </div>
    `
})
export class ToastMessageComponent implements OnInit {
    @Input() toast: ToastMessage
    @Output() onClose: EventEmitter<string> = new EventEmitter();

    ngOnInit() {
        setTimeout(() => {
            this.onClose.emit(this.toast.id)
        }, 3000)
    }
}
