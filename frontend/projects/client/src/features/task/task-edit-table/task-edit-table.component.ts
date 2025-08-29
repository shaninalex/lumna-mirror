import {Component} from '@angular/core';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatTableModule} from '@angular/material/table';
import {DatePipe} from '@angular/common';
import {Task} from '@client/entities/task';


const ELEMENT_DATA: Task[] = [
    {
        id: "cb3af766-28b7-4431-8061-bb1870a0d4fc",
        creator_id: "ffc61ef8-9ea7-4bf9-842b-38201a6b49ae",
        epic_id: "a13938d8-82a3-4b65-a0fe-0e14babde166",
        sprint_id: "86ff9873-fdf2-4314-9aed-e059717638fc",
        project_id: "49d7689c-a434-4d8d-8b2f-50f89d9dab80",
        assignee: "alex",
        completed: false,
        title: "Implement authentication (login/register)",
        list_idx: 0,
        code: "task-385934",
        description: "Use Ory Kratos for identity management",
        status: "ef64b8a3-d8a1-478e-a8b4-a70ecf0ca3f2",
        created_at: new Date("2025-08-24T15:39:09.654941+03:00"),
        updated_at: new Date("2025-08-24T15:39:09.654941+03:00"),
    }
];

@Component({
    selector: "ts-task-edit-table",
    styleUrl: 'task-edit-table.component.scss',
    template: `
        <table class="w-full">
            @for (row of ELEMENT_DATA; track row.id) {
                <tr>
                    <td class="p-2">{{ row.code }}</td>
                    <td class="p-2">{{ row.title }}</td>
                    <td class="p-2">{{ row.status }}</td>
                    <td class="p-2">{{ row.assignee }}</td>
                    <td class="p-2">{{ row.created_at | date }}</td>
                </tr>
            }
        </table>

    `,
    imports: [MatTableModule, MatCheckboxModule, DatePipe],
})
export class TaskEditTableComponent {
    protected readonly ELEMENT_DATA = ELEMENT_DATA;
}
