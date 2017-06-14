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
  private currentUser;
  private url: string;
  constructor(private auth: AuthService, private router: Router) { }

  ngOnInit() {
  }

  ngDoCheck() {
    this.currentUser = JSON.parse(localStorage.getItem('current_user'));
    this.url = this.router.routerState.snapshot.url;
  }

  logout() {
    this.auth.unauthentificatenow();
    window.location.reload()
  }

}
