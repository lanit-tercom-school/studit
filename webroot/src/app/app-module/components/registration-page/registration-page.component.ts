import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { AuthService } from 'services/auth.service';
import { UserInfo } from 'models/user-info';

@Component({
  selector: 'app-registration-page',
  templateUrl: './registration-page.component.html',
  styleUrls: ['./registration-page.component.css']
})
export class RegistrationPageComponent implements OnInit {

  private user: UserInfo = { login: "", nickname: "", password: "" };
  private error: string;

  constructor(private auth: AuthService, private router: Router) { }

  ngOnInit() {
  }

  register() {
    this.auth.register(this.user).subscribe(
      data => {
        this.router.navigate(['/registration/validate']);
      },
      error => {
        console.log(error);
        this.error = error;
      });
  }

}
