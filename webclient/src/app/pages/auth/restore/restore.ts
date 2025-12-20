import { Component } from '@angular/core';
import {AuthLayout} from "@shared/ui";
import {RouterLink} from "@angular/router";

@Component({
  selector: 'app-restore',
    imports: [
        AuthLayout,
        RouterLink
    ],
  templateUrl: './restore.html',
  styleUrl: './restore.css',
})
export class Restore {

}
