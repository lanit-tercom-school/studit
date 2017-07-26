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

  private localUser: User = { login: "", password: "" };
  private ReturnUrl: string

  constructor(private auth: AuthService, private router: Router, private data: DataService) { }

  ngOnInit() {
    this.auth.unauthentificatenow();
    this.ReturnUrl = this.router.routerState.snapshot.root.queryParams['ReturnUrl'] || '/home';
  }

  login() {
    this.auth.authenticatenow(this.localUser).subscribe(
      data => {
        this.data.loadAll();
        /*if (this.ReturnUrl === '/registration') {
          this.ReturnUrl = '/home';
        }
        this.router.navigate([this.ReturnUrl]);*/
         this.router.navigate(['home']);
      },
      error => {
        this.data.alertError(error, 'ERROR: loadProjects() -> getProjectItems()');
      });
  }
}
