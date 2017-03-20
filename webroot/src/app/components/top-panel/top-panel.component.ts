import { Component, OnInit, DoCheck } from '@angular/core';
import { AuthService } from './../../services/auth.service';

@Component({
  selector: 'app-top-panel',
  templateUrl: './top-panel.component.html',
  styleUrls: ['./top-panel.component.css']
})
export class TopPanelComponent implements OnInit, DoCheck {

  private isAuthentificated = false;
  private currentUser = "";

  constructor(private auth: AuthService) { }

  ngOnInit() {
  }

  ngDoCheck() {
    this.isAuthentificated = window.localStorage.getItem("auth_key") != null;
    this.currentUser = window.localStorage.getItem("current_user");
  }

  logout() {
    this.auth.unauthentificatenow();
  }

}
