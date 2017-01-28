import { Component, OnInit } from '@angular/core';
import {ProfileService} from "../profile.service";
import {Activity} from "../activity";

@Component({
  selector: 'app-activity',
  templateUrl: './activity.component.html',
  styleUrls: ['./activity.component.css'],
  providers: [ProfileService]
})
export class ActivityComponent implements OnInit {
  act: Activity[];
  constructor(private actService:ProfileService) { }

  ngOnInit() {
    this.act=this.actService.getActivity();
  }

}
