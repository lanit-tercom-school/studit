import { Component, OnInit } from '@angular/core';

import { DataService } from 'services/data.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  ngOnInit() {
    this.data.loadAll();
  }
  constructor(private data: DataService) { }
}
