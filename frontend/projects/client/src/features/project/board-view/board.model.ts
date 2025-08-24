import {Task} from '@client/entities/task';

export interface StatusColumn {
    id: string;        // same as Status.id
    title: string;     // Status title (TODO, IN PROGRESS, etc.)
    tasks: Task[];   // tasks in this column
}
