import { Component, OnInit, DoCheck } from '@angular/core';
import { AuthService } from './../../services/auth.service';
import { DataService } from './../../services/data.service';
import { User } from './../pages/authorization-page/user';
import { Router } from '@angular/router';


@Component({
  selector: 'app-top-panel',
  templateUrl: './top-panel.component.html',
  styleUrls: ['./top-panel.component.css']
})
export class TopPanelComponent implements OnInit, DoCheck {

  private currentUser;
  url: string;
  constructor(private auth: AuthService, private router: Router, private data: DataService) { }

  ngOnInit() {
    if (localStorage.getItem('current_user')) {
      this.data.loadEnrollingAndUserProjects();
    }
  }

  ngDoCheck() {
    this.currentUser = JSON.parse(localStorage.getItem('current_user'));
    this.url = this.router.routerState.snapshot.url;

  }

  logout() {
    this.auth.unauthentificatenow();
  }

}
