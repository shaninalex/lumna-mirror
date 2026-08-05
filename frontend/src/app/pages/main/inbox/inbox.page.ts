import { Component } from '@angular/core';
import { MainLayout } from '@core/layout';

@Component({
    selector: 'lu-inbox',
    imports: [MainLayout],
    template: `<lu-main-layout> <p>inbox works!</p> </lu-main-layout>`,
})
export class InboxPage {}
