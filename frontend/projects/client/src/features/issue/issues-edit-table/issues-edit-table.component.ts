import {Component} from '@angular/core';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatTableModule} from '@angular/material/table';
import {Issue} from '@client/entities/issue';
import {DatePipe} from '@angular/common';


const ELEMENT_DATA: Issue[] = [
    {
        id: 1,
        type: "Task",
        title: 'Create basic layout structure ( static layout, without state and models )',
        status: "Backlog",
        assignee: "@me",
        createdAt: new Date(),
        updatedAt: new Date(),
    },{
        id: 2,
        type: "Task",
        title: 'Create basic layout structure ( static layout, without state and models )',
        status: "Backlog",
        assignee: "@me",
        createdAt: new Date(),
        updatedAt: new Date(),
    },{
        id: 3,
        type: "Task",
        title: 'Create basic layout structure ( static layout, without state and models )',
        status: "Backlog",
        assignee: "@me",
        createdAt: new Date(),
        updatedAt: new Date(),
    },{
        id: 4,
        type: "Task",
        title: 'Create basic layout structure ( static layout, without state and models )',
        status: "Backlog",
        assignee: "@me",
        createdAt: new Date(),
        updatedAt: new Date(),
    },{
        id: 5,
        type: "Task",
        title: 'Create basic layout structure ( static layout, without state and models )',
        status: "Backlog",
        assignee: "@me",
        createdAt: new Date(),
        updatedAt: new Date(),
    },
];

@Component({
    selector: "ts-issues-edit-table",
    styleUrl: 'issues-edit-table.component.scss',
    template: `
        <table class="w-full">
            @for (row of ELEMENT_DATA; track row.id) {
                <tr>
                    <td class="p-2">{{ row.id }}</td>
                    <td class="p-2">{{ row.title }}</td>
                    <td class="p-2">{{ row.status }}</td>
                    <td class="p-2">{{ row.assignee }}</td>
                    <td class="p-2">{{ row.createdAt | date }}</td>
                </tr>
            }
        </table>

    `,
    imports: [MatTableModule, MatCheckboxModule, DatePipe],
})
export class IssuesEditTableComponent {
    protected readonly ELEMENT_DATA = ELEMENT_DATA;
}
