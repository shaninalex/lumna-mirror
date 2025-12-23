import {Component, inject} from '@angular/core';
import {RouterOutlet} from '@angular/router';
import {AppBootstrapService} from '@core/bootstrap.service';

@Component({
    selector: 'app-root',
    imports: [RouterOutlet],
    templateUrl: './app.html',
    styleUrl: './app.css',
})
export class App {
    readonly application = inject(AppBootstrapService)
}
