import { Component } from '@angular/core';
import {RouterLink} from '@angular/router';

@Component({
  selector: 'fr-login-required.component',
    imports: [
        RouterLink
    ],
  templateUrl: './login-required.component.html',
  styleUrl: './login-required.component.scss'
})
export class LoginRequiredComponent {

}
