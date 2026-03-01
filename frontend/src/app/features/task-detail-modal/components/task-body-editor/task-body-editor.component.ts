import {Component, EventEmitter, Input, OnDestroy, OnInit, Output} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {Editor, NgxEditorComponent, NgxEditorMenuComponent, Toolbar} from 'ngx-editor';
import { toHTML } from 'ngx-editor';

@Component({
    selector: 'app-task-body-editor',
    template: `
        <div class="NgxEditor__Wrapper">
            <ngx-editor-menu [editor]="editor" [toolbar]="toolbar"></ngx-editor-menu>
            <ngx-editor
                [editor]="editor"
                [ngModel]="body"
                [disabled]="false"
                [placeholder]="'Type here...'"
            ></ngx-editor>
        </div>
    `,
    imports: [NgxEditorComponent, NgxEditorMenuComponent, FormsModule],
})
export class TaskBodyEditorComponent implements OnInit, OnDestroy {
    @Input() body: string;
    @Output() onChange: EventEmitter<string> = new EventEmitter();

    editor: Editor;
    toolbar: Toolbar = [
        // default value
        ['bold', 'italic'],
        ['underline', 'strike'],
        ['code', 'blockquote'],
        [{ heading: ['h1', 'h2', 'h3', 'h4', 'h5', 'h6'] }],
        ['link'],
        // or, set options for link:
        //[{ link: { showOpenInNewTab: false } }, 'image'],
        ['text_color'],
        ['ordered_list', 'bullet_list'],
        ['align_left', 'align_center', 'align_right', 'align_justify'],
    ];

    ngOnInit(): void {
        this.editor = new Editor();
        this.editor.valueChanges.subscribe(data => this.onChange.emit(toHTML(data)));
    }

    ngOnDestroy(): void {
        this.editor.destroy();
    }
}
