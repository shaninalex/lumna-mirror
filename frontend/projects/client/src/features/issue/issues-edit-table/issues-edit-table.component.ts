import {Component} from '@angular/core';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatTableModule} from '@angular/material/table';
import {Issue} from '@client/entities/issue';
import {DatePipe} from '@angular/common';


const ELEMENT_DATA: Issue[] = [
    {
        id: "46fc6efa-9c68-48f9-ae3d-9fe7862e93a5",
        creator_id: "ffc61ef8-9ea7-4bf9-842b-38201a6b49ae",
        epic_id: "a13938d8-82a3-4b65-a0fe-0e14babde166",
        sprint_id: "86ff9873-fdf2-4314-9aed-e059717638fc",
        project_id: "49d7689c-a434-4d8d-8b2f-50f89d9dab80",
        assignee: "alex",
        completed: false,
        title: "Use only Material SDK",
        description: "Replace default material components with SDK and manual created ui elements",
        status: "1c176c71-4b6f-4d2a-80b3-ab4c624538fc",
        created_at: new Date("2025-08-24T15:39:09.654941+03:00"),
        updated_at: new Date("2025-08-24T15:39:09.654941+03:00"),
    },
    {
        id: "b53a0e01-22bc-41b5-8a24-d1971834a591",
        creator_id: "ffc61ef8-9ea7-4bf9-842b-38201a6b49ae",
        epic_id: "a13938d8-82a3-4b65-a0fe-0e14babde166",
        sprint_id: "86ff9873-fdf2-4314-9aed-e059717638fc",
        project_id: "49d7689c-a434-4d8d-8b2f-50f89d9dab80",
        assignee: "alex",
        completed: false,
        title: "Add user profile \u0026 settings page",
        description: "Allow users to update their information",
        status: "1c176c71-4b6f-4d2a-80b3-ab4c624538fc",
        created_at: new Date("2025-08-24T15:39:09.654941+03:00"),
        updated_at: new Date("2025-08-24T15:39:09.654941+03:00"),
    },
    {
        id: "be474dc9-be55-42c6-8331-3335484df38a",
        creator_id: "ffc61ef8-9ea7-4bf9-842b-38201a6b49ae",
        epic_id: "385656d7-fb44-45b2-90c5-99b10bb13618",
        sprint_id: "99cb782c-dfdc-4430-ad35-d8595c712592",
        project_id: "49d7689c-a434-4d8d-8b2f-50f89d9dab80",
        assignee: "alex",
        completed: false,
        title: "Polish dashboard UI",
        description: "Make Taskiro visually appealing with Tailwind \u0026 animations",
        status: "1c176c71-4b6f-4d2a-80b3-ab4c624538fc",
        created_at: new Date("2025-08-24T15:39:09.654941+03:00"),
        updated_at: new Date("2025-08-24T15:39:09.654941+03:00"),
    },
    {
        id: "9e9147af-f857-429b-8fbf-df3b115bac09",
        creator_id: "ffc61ef8-9ea7-4bf9-842b-38201a6b49ae",
        epic_id: "0712bf3b-6965-4e95-b6a2-f617436b305a",
        sprint_id: "99cb782c-dfdc-4430-ad35-d8595c712592",
        project_id: "49d7689c-a434-4d8d-8b2f-50f89d9dab80",
        assignee: "alex",
        completed: false,
        title: "Add issues \u0026 epics",
        description: "Core task tracking functionality",
        status: "d581023a-236c-4e54-8419-0c279e30e3c3",
        created_at: new Date("2025-08-24T15:39:09.654941+03:00"),
        updated_at: new Date("2025-08-24T15:39:09.654941+03:00"),
    },
    {
        id: "3ba19c38-085b-403d-a13b-1d577b489447",
        creator_id: "ffc61ef8-9ea7-4bf9-842b-38201a6b49ae",
        epic_id: "0712bf3b-6965-4e95-b6a2-f617436b305a",
        sprint_id: "99cb782c-dfdc-4430-ad35-d8595c712592",
        project_id: "49d7689c-a434-4d8d-8b2f-50f89d9dab80",
        assignee: "alex",
        completed: false,
        title: "Create projects \u0026 organizations",
        description: "Implement CRUD for projects/organizations",
        status: "88c06901-ce83-495d-b970-38121bb9a54d",
        created_at: new Date("2025-08-24T15:39:09.654941+03:00"),
        updated_at: new Date("2025-08-24T15:39:09.654941+03:00"),
    },
    {
        id: "cb3af766-28b7-4431-8061-bb1870a0d4fc",
        creator_id: "ffc61ef8-9ea7-4bf9-842b-38201a6b49ae",
        epic_id: "a13938d8-82a3-4b65-a0fe-0e14babde166",
        sprint_id: "86ff9873-fdf2-4314-9aed-e059717638fc",
        project_id: "49d7689c-a434-4d8d-8b2f-50f89d9dab80",
        assignee: "alex",
        completed: false,
        title: "Implement authentication (login/register)",
        description: "Use Ory Kratos for identity management",
        status: "ef64b8a3-d8a1-478e-a8b4-a70ecf0ca3f2",
        created_at: new Date("2025-08-24T15:39:09.654941+03:00"),
        updated_at: new Date("2025-08-24T15:39:09.654941+03:00"),
    }
];

@Component({
    selector: "ts-issues-edit-table",
    styleUrl: 'issues-edit-table.component.scss',
    template: `
        <table class="w-full">
            @for (row of ELEMENT_DATA; track row.id) {
                <tr>
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
export class IssuesEditTableComponent {
    protected readonly ELEMENT_DATA = ELEMENT_DATA;
}
