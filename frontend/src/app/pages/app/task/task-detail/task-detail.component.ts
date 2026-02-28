import { Component, inject } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { filter, map, Observable, tap } from 'rxjs';
import { TaskModel } from '@entities/task';
import { UiService } from '@shared/ui';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-task-detail-page',
    template: `
        <div class="fixed inset-0 z-50 flex items-start justify-center pt-16 px-4" (click)="onBackdropClick($event)">
            <div class="absolute inset-0 bg-black/50"></div>
            <div class="relative bg-white dark:bg-gray-800 rounded-lg shadow-xl w-full max-w-2xl max-h-[80vh] overflow-y-auto p-6">
                <button class="absolute top-4 right-4 text-gray-400 hover:text-gray-600 cursor-pointer" (click)="close()">
                    <i class="fa-solid fa-xmark text-xl"></i>
                </button>
                @if(task$ | async; as task) {
                    <h1 class="text-xl font-semibold">{{ task.title }}</h1>
                }
            </div>
        </div>
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
