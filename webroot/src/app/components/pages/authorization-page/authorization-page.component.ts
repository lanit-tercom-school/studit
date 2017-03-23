import { Component, OnInit } from '@angular/core';
import { AuthService } from './../../../services/auth.service';
import { Router } from '@angular/router';
import { User } from './user';
import { NgModule } from '@angular/core';

@Component({
  selector: 'app-authorization-page',
  templateUrl: './authorization-page.component.html',
  styleUrls: ['./authorization-page.component.css'],
  providers: [AuthService],
})
export class AuthorizationPageComponent implements OnInit {

  private localUser: User = { Login: "", Password: "" };
  private error: any;

  constructor(private auth: AuthService, private router: Router) { }

  ngOnInit() {
    this.auth.unauthentificatenow();
  }

  login() {
    this.auth.authenticatenow(this.localUser).subscribe(
      data => {
        this.router.navigate(['/home']);
      },
      error => {
        console.log(error);
        this.error = error;
      });
  }
}