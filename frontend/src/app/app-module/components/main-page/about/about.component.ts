import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { UserRegister } from 'models/user-register';
import { AuthService } from 'services/auth.service';

@Component({
  selector: 'app-about',
  templateUrl: './about.component.html',
  styleUrls: ['./about.component.css']
})
export class AboutComponent implements OnInit {
  public User: UserRegister = { login: "", nickname: "", password: "" };
  private error: string;
  private currentUser;
  constructor(private auth: AuthService, private router: Router) {
    this.currentUser = JSON.parse(localStorage.getItem('current_user'));
  }

  ngOnInit() {
  }

  register() {
    this.auth.register(this.User).subscribe(
      data => {
        this.router.navigate(['/registration/validate']);
      },
      error => {
        console.log(error);
        this.error = error;
      });
  }

}
