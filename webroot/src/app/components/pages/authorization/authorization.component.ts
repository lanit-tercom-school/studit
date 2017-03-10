import { Component, OnInit } from '@angular/core';
import { AuthService } from './../../../services/auth.service';
import { Router } from '@angular/router';
import { User } from './user';
import { NgModule } from '@angular/core';

@Component({
  selector: 'app-authorization',
  templateUrl: './authorization.component.html',
  styleUrls: ['./authorization.component.css'],
  providers: [AuthService],
})
export class AuthorizationComponent implements OnInit {

  private localUser: User = { email: "", password: "" };

  constructor(private auth: AuthService, private router: Router) { }

  ngOnInit() {
  }

  login() {
    this.auth.authenticatenow(this.localUser).subscribe(
      data => {
        this.router.navigate(['/home']);
      },
      error => {
        console.log("invalid login or password");
      });
  }
}