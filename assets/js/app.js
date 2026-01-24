const EVENT_BOARD_COLUMN_MOVED = "board:column-moved";
const EVENT_BOARD_TASK_MOVED = "board:task-moved";
const EVENT_BOARD_CREATED = "board:created";
const EVENT_BOARD_DELETED = "board:deleted";

function getCsrfToken() {
    const meta = document.querySelector('meta[name="csrf-token"]');
    return meta ? meta.getAttribute("content") : "";
}

const columns = document.querySelectorAll("[data-column]");
columns.forEach(function (column) {
    new Sortable(column, {
        animation: 150,
        group: "shared",
        // TODO: on before drop and revert if error ?
        onEnd: function (e) {
            document.dispatchEvent(
                new CustomEvent(EVENT_BOARD_TASK_MOVED, {
                    detail: {
                        taskId: e.item.id,
                        toColumnId: e.to.dataset.column,
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
        // TODO: on before drop and revert if error ?
        onEnd: function (e) {
            document.dispatchEvent(
                new CustomEvent(EVENT_BOARD_COLUMN_MOVED, {
                    detail: {
                        columnId: e.item.id,
                        newIndex: e.newIndex,
                    },
                }),
            );
        },
    });
}

document.addEventListener(EVENT_BOARD_COLUMN_MOVED, async (e) => {
    const { columnId, newIndex } = e.detail;
    try {
        const response = await fetch(`/hx/projects/lists/${columnId}/reorder`, {
            method: "PATCH",
            headers: {
                "Content-Type": "application/json",
                "X-CSRF-TOKEN": getCsrfToken(),
            },
            body: JSON.stringify({ order: newIndex }),
        });
        if (!response.ok) {
            console.error("Failed to reorder column");
            alert(err); // show toast instead
        }
    } catch (err) {
        console.error("Failed to reorder column:", err);
        alert(err); // show toast instead
    }
});

document.addEventListener(EVENT_BOARD_TASK_MOVED, async (e) => {
    const { taskId, toColumnId, newIndex } = e.detail;
    try {
        const response = await fetch(`/hx/projects/tasks/${taskId}/reorder`, {
            method: "PATCH",
            headers: {
                "Content-Type": "application/json",
                "X-CSRF-TOKEN": getCsrfToken(),
            },
            body: JSON.stringify({
                boardListId: toColumnId,
                order: newIndex,
            }),
        });
        if (!response.ok) {
            console.error("Failed to reorder task");
            alert(err); // show toast instead
        }
    } catch (err) {
        console.error("Failed to reorder task:", err);
        alert(err); // show toast instead
    }
});

document.addEventListener(EVENT_BOARD_CREATED, async (e) => {
    const boardAddFormModal = document.querySelector("#boardAddFormModal");
    if (!boardAddFormModal) return;

    boardAddFormModal.querySelector("form").reset();
    const modal = bootstrap.Modal.getInstance(boardAddFormModal);
    if (modal) modal.hide();
});

document.addEventListener(EVENT_BOARD_DELETED, async (e) => {
    console.log(e);
});

function closeModal(id) {
    const modal = document.querySelector(id);
    if (!modal) return;
    modal.remove();
}
