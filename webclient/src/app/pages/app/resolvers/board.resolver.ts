import {ActivatedRouteSnapshot, ResolveFn, Router, RouterStateSnapshot, UrlTree} from '@angular/router';
import {inject} from '@angular/core';
import {toObservable} from '@angular/core/rxjs-interop';
import {filter, firstValueFrom, map} from 'rxjs';
import {BoardModel, BoardStore} from '@entities/board';

export const boardResolver: ResolveFn<BoardModel | UrlTree> = (
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot,
)=> {
    const store = inject(BoardStore);
    const boardId = parseInt(route.paramMap.get('boardId')!);
    return firstValueFrom(
        toObservable(store.entities).pipe(
            filter(projects => projects.some(p => p.id === boardId)),
            map(projects => projects.find(p => p.id === boardId)!),
        )
    );
}
