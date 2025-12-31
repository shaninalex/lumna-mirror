import { Injectable } from "@angular/core";

@Injectable({
    providedIn: 'root'
})
export class KanbanBoardService {

    /*
     * BuildTree converts list of ListModel and list of TaskModel to a tree like struct where []ListModel
     * contain []TaskModel. It helps to display and nested drag&drop lists UI
     * */
    public BuildTree(): object {
        return {}
    }
}
