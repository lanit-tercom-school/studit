import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { AuthService } from 'services/auth.service';
import { AlertService } from 'services/alert.service';

import { UserRegister } from 'models/user-register';

@Component({
  selector: 'app-registration-page',
  templateUrl: './registration-page.component.html',
  styleUrls: ['./registration-page.component.css']
})
export class RegistrationPageComponent implements OnInit {

  private user: UserRegister = { login: "", nickname: "", password: "" };

  constructor(private auth: AuthService,
    private router: Router,
    private alert: AlertService) { }

  ngOnInit() {
  }

  register() {
    this.auth.register(this.user).subscribe(
      () => {
        this.router.navigate(['/registration/validate']);
      },
      error => {
        this.alert.alertError(error, 'register() -> auth.register()');
      });
  }

}
