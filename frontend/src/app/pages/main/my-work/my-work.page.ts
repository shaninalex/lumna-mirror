import { Component } from '@angular/core';
import { MainLayout } from '@core/layout';

@Component({
    selector: 'lu-my-work',
    imports: [MainLayout],
    template: `<lu-main-layout> <p>my-work works!</p> </lu-main-layout>`,
})
export class MyWorkPage {}
