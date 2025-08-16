import { Component, Input } from "@angular/core";
import { UiText } from "@ory/kratos-client";

@Component({
    selector: "jr-ui-kratos-form-messages",
    standalone: true,
    template: `
@if (messages) {
    @for (message of messages; track $index) {
        <div class="mt-2 text-sm text-white rounded-lg p-4 mb-4"
            [class.bg-red-400]="message.type === 'error'"
            [class.bg-blue-400]="message.type === 'info'"
            [class.bg-green-400]="message.type === 'success'"
            role="alert">
            {{ message.text }}
        </div>
    }
}
`
})
export class FormMessagesComponent {
    @Input() messages?: UiText[];
}
