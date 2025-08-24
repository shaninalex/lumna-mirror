import {Issue} from '@client/entities/issue';

export interface StatusColumn {
    id: string;        // same as Status.id
    title: string;     // Status title (TODO, IN PROGRESS, etc.)
    issues: Issue[];   // tasks in this column
}
