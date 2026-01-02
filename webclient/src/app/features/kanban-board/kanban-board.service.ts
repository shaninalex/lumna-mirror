import { computed, inject, Injectable, Signal } from "@angular/core";
import { KanbanModel } from './model/kanban.model';
import { ListStore } from '@entities/list';
import { TaskStore } from '@entities/task';
import { of } from "rxjs";

@Injectable()
export class KanbanBoardService {
    private readonly listStore = inject(ListStore);
    private readonly taskStore = inject(TaskStore);

    /*
     * BuildTree converts list of ListModel and list of TaskModel to a tree like struct where []ListModel
     * contain []TaskModel. It helps to display and nested drag&drop lists UI.
     * Main idea here that KanbanBoard component should operate only with that model.
     * */
    public BuildTree(boardId: number): any {
        const lists = this.listStore.boardLists(boardId);
        const tasks = this.taskStore.boardTasks(boardId);
        
        return lists
            .slice()
            .sort((a, b) => a.order - b.order)
            .map(list => ({
                id: list.id,
                title: list.name,
                order: list.order,
                tasks: tasks
                    .filter(t => t.list_id === list.id)
                    .sort((a, b) => a.order - b.order),
            })
        );
    }
}
