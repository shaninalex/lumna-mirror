import {Component, Input, OnInit} from '@angular/core';
import {UiNode, UiNodeAttributes} from '@ory/kratos-client';

function nodeAttributes(attributes: UiNodeAttributes): any {
    return attributes as any
}

@Component({
    selector: 'ui-kratos-input',
    template: `
        <!-- a -->
        @if (node.type === 'a') {
            <div class="mb-4">
                <a class="rounded border-slate-300 bg-slate-200 px-2 py-1 cursor-pointer hover:bg-slate-300"
                   [attr.id]="attr().id" [attr.href]="attr().href">
                    {{ node.meta.label?.text }}
                </a>
            </div>
        }

        <!-- image -->
        @if (node.type === 'img') {
            <div class="mb-4">
                <p>{{ node.meta.label?.text }}</p>
                <img class="img-fluid" [attr.id]="attr().id" [attr.alt]="attr().name" [attr.width]="attr().width"
                     [attr.height]="attr().height" [attr.src]="attr().src"/>
            </div>
        }

        <!-- input -->
        @if (node.attributes.node_type === 'input') {
            @if (attr().type === 'submit') {
                <div class="mb-4">
                    <button
                        class="rounded border-slate-300 bg-slate-200 px-2 py-1 cursor-pointer hover:bg-slate-300"
                        [attr.disabled]="attr().disabled ? true : null" [attr.name]="attr().name"
                        [attr.type]="attr().type"
                        [attr.value]="attr().value">
                        {{ node.meta.label?.text }}
                    </button>
                </div>
            } @else {
                @if (attr().type === 'hidden') {
                    <input [attr.disabled]="attr().disabled ? true : null" [attr.name]="attr().name"
                           [attr.required]="attr().required"
                           [attr.type]="attr().type" [attr.placeholder]="node.meta.label?.text"
                           [attr.value]="attr().value? attr().value: null"/>
                } @else {
                    <div class="mb-4">
                        <label class="form-label" [attr.for]="node.attributes.name">{{ node.meta.label?.text }}</label>
                        <input class="border w-full py-1 px-2 rounded"
                               [attr.id]="node.attributes.name"
                               [attr.disabled]="attr().disabled ? true : null"
                               [attr.autocomplete]="attr().autocomplete"
                               [attr.name]="attr().name"
                               [attr.type]="attr().type"
                               [attr.placeholder]="node.meta.label?.text" [attr.value]="attr().value? attr().value: null"/>
                        @for (message of node.messages; track message) {
                            <div class="alert alert-danger" [class.text-danger]="message.type === 'error'"
                                 [class.text-info]="message.type === 'info'"
                                 [class.text-success]="message.type === 'success'">
                                {{ message.text }}
                            </div>
                        }
                    </div>
                }
            }
        }
        <!-- script (not implemented) -->

        <!-- text -->
        @if (node.type === 'text') {
            <div class="alert alert-info" role="alert">
                <h4 [id]="attr().id" class="alert-heading">{{ attr().text.text }}</h4>
                <p [id]="node.meta.label?.id">{{ node.meta.label?.text }}</p>
            </div>
        }
        `
})
export class KratosInputComponent {
    @Input() node: UiNode;
    attr(): any {
        return nodeAttributes(this.node.attributes)
    }
}
