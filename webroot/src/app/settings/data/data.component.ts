import { Component, OnInit } from '@angular/core';
import {SettingsService} from "../settings.service";
import {Data} from "../data";

@Component({
  selector: 'app-data',
  templateUrl: './data.component.html',
  styleUrls: ['./data.component.css'],
  providers: [SettingsService]
})
export class DataComponent implements OnInit {

  constructor(private dataService: SettingsService) { }
  data: Data[];
  ngOnInit() {
    this.data=this.dataService.getInformation();
  }

}
