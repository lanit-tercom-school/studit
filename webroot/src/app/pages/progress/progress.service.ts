import { Injectable } from '@angular/core';
import {TaskItem} from "./task-item";
import {ProjectCardItem} from "./project-card-item";
import {NotificationItem} from "./notification-item";

@Injectable()
export class ProgressService {

  constructor() { }

  getTasks(): TaskItem[] {
    return [
      {
        "Id": 1,
        "Name": "Name of task 1"
      },
      {
        "Id": 2,
        "Name": "Name of task 2"
      }
    ];
  }

  getNotifications(): NotificationItem[] {
    return [
      {
        "Id": 1,
        "Name": "Name of notification 1"
      },
      {
        "Id": 2,
        "Name": "Name of notification 2"
      }
    ];
  }

  getProjectCards(): ProjectCardItem[] {
    return [
      {
        "Id": 1,
        "Name": "Name of project 1",
        "Organization": "Lanit 1",
        "Description": "Description of project 1"
      },
      {
        "Id": 2,
        "Name": "Name of project 2",
        "Organization": "Lanit 2",
        "Description": "Description of project 2"
      }
    ];
  }

}
