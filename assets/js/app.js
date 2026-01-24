const EVENT_BOARD_COLUMN_MOVED = "board:column-moved";
const EVENT_BOARD_TASK_MOVED = "board:task-moved";

const columns = document.querySelectorAll("[data-column]");
columns.forEach(function (column) {
    new Sortable(column, {
        animation: 150,
        group: "shared",
        onEnd: function (e) {
            console.log(e);
            document.dispatchEvent(
                new CustomEvent(EVENT_BOARD_TASK_MOVED, {
                    detail: {
                        taskId: e.item.id,
                        fromColumnId: e.from.dataset.column,
                        toColumnId: e.to.dataset.column,
                        oldIndex: e.oldIndex,
                        newIndex: e.newIndex,
                    },
                }),
            );
        },
    });
});

const board = document.querySelector("[data-board]");
if (board) {
    new Sortable(board, {
        animation: 150,
        onEnd: function (e) {
            document.dispatchEvent(
                new CustomEvent(EVENT_BOARD_COLUMN_MOVED, {
                    detail: {
                        columnId: e.item.id,
                        oldIndex: e.oldIndex,
                        newIndex: e.newIndex,
                    },
                }),
            );
        },
    });
}

document.addEventListener(EVENT_BOARD_COLUMN_MOVED, (e) =>
    console.log(e.detail),
);
document.addEventListener(EVENT_BOARD_TASK_MOVED, (e) => console.log(e.detail));
