// Source - https://stackoverflow.com/a/76676182
// Posted by BikyMandal
// Retrieved 2026-08-17, License - CC BY-SA 4.0

import type { PipeTransform } from '@angular/core';
import { Pipe } from '@angular/core';
import { formatDistance } from 'date-fns';

@Pipe({
    name: 'timeAgo',
})
export class TimeAgoPipe implements PipeTransform {
    transform(value: Date): unknown {
        return formatDistance(value, new Date(), { addSuffix: true });
    }
}
