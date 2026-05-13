export interface ListModel {
    id: number;
    title: string;
    statuses: Status[];
    project_id: number;
    workspace_id: number;
}

// Task status. By default each list have "todo", "in progress", "done" statuses
// User can create his own status.
// In Board view statuses become columns, and change statuses happening via drag-n-drop
// in Backlog view tasks can change statuses via dropdown
export interface Status {
    id: number;
    title: string;

    // tasks with this status automatically completed.
    final: boolean;
}
