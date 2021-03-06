import { Component, OnInit, NgModule } from '@angular/core';
import { Router } from '@angular/router';

import { AuthService } from 'services/auth.service';
import { DataService } from 'services/data.service';
import { User } from 'models/user';


@Component({
  selector: 'app-authorization-page',
  templateUrl: './authorization-page.component.html',
  styleUrls: ['./authorization-page.component.css'],
  providers: [AuthService],
})
export class AuthorizationPageComponent implements OnInit {

  public LocalUser: User = { login: "", password: "" };
  private error: any;
  private ReturnUrl: string

  constructor(private auth: AuthService, private router: Router, private data: DataService) { }

  ngOnInit() {
    window.scrollTo(0,0);
    this.auth.unauthentificatenow();
    this.ReturnUrl = this.router.routerState.snapshot.root.queryParams['ReturnUrl'] || '/home';
  }

  login() {
    this.auth.authenticatenow(this.LocalUser).subscribe(
      data => {
        this.data.loadAll();
        /*if (this.ReturnUrl === '/registration') {
          this.ReturnUrl = '/home';
        }
        this.router.navigate([this.ReturnUrl]);*/
        this.router.navigate([this.returnPage()]);
      },
      error => {
        console.log(error);
        this.error = error;
      });
  }

  returnPage(): string{
    switch(this.ReturnUrl){
      case "/registration/validate":{
        return '/home';
      }
      case"/registration":{
        return '/home';
      }
      case "/error":{
        return '/home';
      }
      default:{
        return this.ReturnUrl;
      }
    }
  }
}
