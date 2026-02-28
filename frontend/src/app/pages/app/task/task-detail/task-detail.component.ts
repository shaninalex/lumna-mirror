import { Component, inject } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { filter, map, Observable, tap } from 'rxjs';
import { TaskModel } from '@entities/task';
import { UiService } from '@shared/ui';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-task-detail-page',
    template: `
        <!-- Modal -->
        <div class="modal fade show d-block" tabindex="-1" role="dialog" (click)="onBackdropClick($event)">
            <div class="modal-dialog modal-lg modal-dialog-centered modal-dialog-scrollable">
                <div class="modal-content">

                    <!-- Header -->
                    <div class="modal-header">
                        @if (task$ | async; as task) {
                            <h5 class="modal-title">Task: {{ task.title }}</h5>
                        }
                        <button
                                type="button"
                                class="btn-close"
                                aria-label="Close"
                                (click)="close()">
                        </button>
                    </div>

                    <!-- Body -->
                    <div class="modal-body">
                        <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus pretium ac nisi et ornare.
                            Curabitur luctus nisl eu luctus pellentesque. Proin dignissim nunc a lectus interdum congue.
                            Aliquam vel viverra nisl, vel tempor diam. Phasellus eu ex eu arcu finibus fermentum non id
                            lacus. In lacinia ipsum quis risus luctus mollis. In dictum, erat quis mattis feugiat, nunc
                            sem rhoncus magna, laoreet vulputate est tortor quis lectus.</p>
                    </div>

                    <!-- Footer -->
                    <div class="modal-footer">
                        <button
                                type="button"
                                class="btn btn-secondary"
                                (click)="close()">
                            Close
                        </button>
                    </div>

                </div>
            </div>
        </div>

        <!-- Backdrop -->
        <div class="modal-backdrop fade show" (click)="onBackdropClick($event)"></div>
    `,
    imports: [AsyncPipe],
})
export class TaskDetailComponent {
    private route = inject(ActivatedRoute);
    private router = inject(Router);
    private ui = inject(UiService);

    task$: Observable<TaskModel> = this.route.data.pipe(
        filter((data) => !!data['task']),
        map((data) => data['task'] as TaskModel),
        tap((task) => this.ui.setPageTitle(`Task: ${task.title}`)),
    );

    close(): void {
        this.router.navigate(['../..'], { relativeTo: this.route });
    }

    onBackdropClick(event: MouseEvent): void {
        if (event.target === event.currentTarget) {
            this.close();
        }
    }
}
