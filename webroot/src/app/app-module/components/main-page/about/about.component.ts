import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { UserInfo } from 'models/user-info';
import { ApiService } from 'services/api.service';

@Component({
  selector: 'app-about',
  templateUrl: './about.component.html',
  styleUrls: ['./about.component.css']
})
export class AboutComponent implements OnInit {
  private user: UserInfo = { login: "", nickname: "", password: "" };
  private error: string;
  private currentUser;
  constructor(private api: ApiService, private router: Router) {
    this.currentUser = JSON.parse(localStorage.getItem('current_user'));
  }

  ngOnInit() {
  }

  register() {
    this.api.register(this.user).subscribe(
      data => {
        this.router.navigate(['/registration/validate']);
      },
      error => {
        console.log(error);
        this.error = error;
      });
  }

}
