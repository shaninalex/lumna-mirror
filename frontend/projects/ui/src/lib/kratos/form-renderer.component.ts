import {Component, Input} from '@angular/core';
import {Observable} from 'rxjs';
import {AsyncPipe} from '@angular/common';
import {KratosInputComponent} from './input.component';
import {TFlowUI} from './interfaces';
import {FormMessagesComponent} from '@dev/ui/kratos/form-message.component';

@Component({
    selector: "ui-form-renderer",
    imports: [
        AsyncPipe,
        KratosInputComponent,
        FormMessagesComponent
    ],
    template: `
        @if (flow$ | async; as flow) {
            @if (flow.ui.messages) {
                <ui-kratos-form-messages [messages]="flow.ui.messages"/>
            }
            <form [action]="flow.ui.action" [method]="flow.ui.method" class="d-flex flex-column">
                @for (node of flow.ui.nodes; track node) {
                    <ui-kratos-input [node]="node"/>
                }
            </form>
        }
    `
})
export class KratosFormRenderer {
    @Input() flow$: Observable<TFlowUI>
}
