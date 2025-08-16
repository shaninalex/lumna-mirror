/* eslint-disable @typescript-eslint/no-explicit-any */

// Since we don't know what type of node attribute is
// I need to disable any type checking for this file
import { Component, Input } from "@angular/core";
import { UiNode, UiNodeAttributes } from "@ory/kratos-client";
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';

function nodeAttributes(attributes: UiNodeAttributes): any {
    return attributes as any
}

@Component({
    selector: 'jr-ui-kratos-input',
    templateUrl: './kratos-input.component.html',
    imports: [
        MatInputModule,
        MatButtonModule,
    ],
    standalone: true
})
export class KratosInputComponent {
    @Input() node: UiNode;

    attr(): any {
        return nodeAttributes(this.node.attributes)
    }
}

