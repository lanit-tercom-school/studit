import { Component, OnInit } from '@angular/core';
import { AuthService } from './../../../services/auth.service';
import { DataService } from './../../../services/data.service';
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
  private ReturnUrl: string

  constructor(private auth: AuthService, private router: Router, private data: DataService) { }

  ngOnInit() {
    this.auth.unauthentificatenow();
    this.ReturnUrl = this.router.routerState.snapshot.root.queryParams['ReturnUrl'] || '/home';
    console.log("You will be redirected to", this.ReturnUrl);

  }

  login() {
    this.auth.authenticatenow(this.localUser).subscribe(
      data => {
        if (this.ReturnUrl == '/registration') {
          this.ReturnUrl = '/home';
        }
        this.data.load_data();
        this.router.navigate([this.ReturnUrl]);
      },
      error => {
        console.log(error);
        this.error = error;
      });
  }
}
