import { Component, OnInit } from '@angular/core';
import {ProgressService} from "../progress.service";
import {NotificationItem} from "../notification";

@Component({
  selector: 'app-notification-list',
  templateUrl: './notification-list.component.html',
  styleUrls: ['./notification-list.component.css']
})
export class NotificationListComponent implements OnInit {

  constructor(private progressService: ProgressService) { }

  notificationItems: NotificationItem[];

  ngOnInit() {
    this.notificationItems = this.progressService.getNotifications();
  }

}
