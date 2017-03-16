import { Component, OnInit, DoCheck } from '@angular/core';
import { AuthService } from './../../services/auth.service';

@Component({
  selector: 'app-top-panel',
  templateUrl: './top-panel.component.html',
  styleUrls: ['./top-panel.component.css']
})
export class TopPanelComponent implements OnInit, DoCheck {

  private isAuthentificated = false;

  constructor(private auth: AuthService) { }

  ngOnInit() {

  }

  ngDoCheck() {
    this.isAuthentificated = window.localStorage.getItem("auth_key") != null;
  }

  logout() {
    window.localStorage.removeItem("auth_key");
    this.auth.unauthentificatenow();
  }

}
