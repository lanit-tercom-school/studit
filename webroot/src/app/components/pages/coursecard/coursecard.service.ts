import { Injectable } from '@angular/core';
import {Coursecard} from "./coursecard";

@Injectable()
export class CoursecardService {

  constructor() { }
  getCoursecards(): Coursecard[] {
    return [
      {
        "Id": 1,
        "Name": "Go",
        "Date": "5 мая 2017 г. — 30 июня 2017 г."
      },
      {
        "Id": 2,
        "Name": "AngularJS",
        "Date": "15 сентября 2017 г. — 19 декабря 2017 г."
      },
      {
        "Id": 3,
        "Name": "Team leading",
        "Date": "15 сентября 2017 г. — 19 декабря 2017 г."
      },
      {
        "Id": 4,
        "Name": "UX/UI",
        "Date": "15 сентября 2017 г. — 19 декабря 2017 г."
      }
    ];
  }
}
