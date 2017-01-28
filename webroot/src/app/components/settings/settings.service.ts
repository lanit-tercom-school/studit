import { Injectable } from '@angular/core';
import {Data} from "./data";
import {SocialProfile} from "./social-profile";

@Injectable()
export class SettingsService {

  constructor() {}

  getInformation(): Data[] {
    return [
      {
        "Address": "dhisd@mail.com",
        "Number": 89382378427,
        "LastPassword": "ede",
        "NewPassword": "eepek",
        "RepeatPassword": "eepek",
        "Interests": "love",
        "Photo": "photo"
      }
    ];
  }
  getProfile(): SocialProfile[]{
    return [
      {
        "Profile":"profile"
      }
      ]
  }

}
