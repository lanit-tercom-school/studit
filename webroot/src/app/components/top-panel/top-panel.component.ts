import { Component, OnInit, DoCheck } from '@angular/core';
import { AuthService } from './../../services/auth.service';
import { User } from './../pages/authorization-page/user';

@Component({
  selector: 'app-top-panel',
  templateUrl: './top-panel.component.html',
  styleUrls: ['./top-panel.component.css']
})
export class TopPanelComponent implements OnInit, DoCheck {

  private currentUser;

  constructor(private auth: AuthService) {
   }

  ngOnInit() {
  }

  ngDoCheck() {
    this.currentUser = JSON.parse(localStorage.getItem('current_user'));
  }

  logout() {
    this.auth.unauthentificatenow();
  }

}
