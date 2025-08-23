export type SubtaskStatus = "To Do" | "In Progress" | "Done";

export interface Subtask {
    id: number;
    parentId: number;
    title: string;
    status: SubtaskStatus;
    assignee?: string;
}
