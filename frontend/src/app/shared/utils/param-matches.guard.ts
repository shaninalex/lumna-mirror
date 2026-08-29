import type { CanMatchFn } from '@angular/router';

export const paramMatchesDigitsOnly: RegExp = /^\d+$/

export function paramMatches(paramName: string, pattern: RegExp): CanMatchFn {
    return (route, segments) => {
        const paramIndex = route.path?.split('/').findIndex((part) => part === `:${paramName}`);

        if (paramIndex === undefined || paramIndex < 0) {
            return false;
        }

        return pattern.test(segments[paramIndex]?.path ?? '');
    };
}
