import { Component, OnInit } from '@angular/core';
import {SettingsService} from "../settings.service";
import {SocialProfile} from "../social-profile";

@Component({
  selector: 'app-social-profile',
  templateUrl: './social-profile.component.html',
  styleUrls: ['./social-profile.component.css'],
  providers: [SettingsService]

})
export class SocialProfileComponent implements OnInit {

  constructor(private profileService: SettingsService) { }
  profile: SocialProfile[];
  ngOnInit() {
    this.profile=this.profileService.getProfile();
  }

}
