import { Component, OnInit } from '@angular/core';
import { AuthService } from './../../../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-authorization',
  templateUrl: './authorization.component.html',
  styleUrls: ['./authorization.component.css']
})
export class AuthorizationComponent implements OnInit {
  localUser = {
    username: '',
    password: ''
  }
  constructor(private auth: AuthService, private router: Router) { }

  ngOnInit() {
  }

  login() {
    let checknow = this.auth.authenticatenow(this.localUser);
    checknow.then((res) => {
      if (res) {
        this.router.navigate(['/project']);
      }
      else {
        console.log('Invalid user');
      }
    })
  }

}