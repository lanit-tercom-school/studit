import { Component, OnInit, DoCheck } from '@angular/core';
import { Router } from '@angular/router';

import { AuthService } from 'services/auth.service';
import { User } from 'models/user';


@Component({
  selector: 'app-top-panel',
  templateUrl: './top-panel.component.html',
  styleUrls: ['./top-panel.component.css']
})
export class TopPanelComponent implements OnInit, DoCheck {
  public CurrentUser;
  private url: string;
  constructor(private auth: AuthService, private router: Router) { }

  ngOnInit() {
  }

  ngDoCheck() {
    this.CurrentUser = JSON.parse(localStorage.getItem('current_user'));
    this.url = this.router.routerState.snapshot.url;
  }

  logout() {
    this.auth.unauthentificatenow();
    this.router.navigateByUrl('/auth');
  }

}
