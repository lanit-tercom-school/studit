import { Injectable } from '@angular/core';
import {User} from "./user";
import {Activity} from "./activity";

@Injectable()
export class ProfileService {

  constructor() {
  }

  getInformation(): User[] {
    return [
      {
        "Information": "I am here to show u my capabilities",
        "Name": "King",
        "Lessons": 10,
        "Rank": 100000,
        "Score": 10,
        "Status": "online",
        "Avatar": "myphoto"
      },
    ]
  }

  getActivity(): Activity[] {
    return [
      {
        "Days": "10",
        "Hometask": 100,
        "Comments": 10000,
        "Missed_lessons": 2,
        "Projects": 2
      },
    ]
  }

}
