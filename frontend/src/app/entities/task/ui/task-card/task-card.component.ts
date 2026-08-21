import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';

@Component({
    selector: 'lu-task-card',
    imports: [RouterLink],
    template: `
        <a routerLink="/app/w/1/task" class="card text-decoration-none text-body">
            <div class="card-body">
                <div class="d-flex justify-content-between mb-2">
                    <span class="badge text-bg-primary"> Feature </span>

                    <small class="text-muted"> #AUTH-42 </small>
                </div>

                <h6 class="mb-2">Implement OAuth login</h6>

                <p class="small text-muted mb-3">Add Google and GitHub authentication.</p>

                <div class="d-flex justify-content-between align-items-center">
                    <small class="text-muted"> Alex </small>

                    <i class="fa-regular fa-message"></i>
                </div>
            </div>
        </a>
    `,
})
export class TaskCardComponent {}
