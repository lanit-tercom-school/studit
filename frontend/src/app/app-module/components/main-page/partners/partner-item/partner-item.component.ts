import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-partner-item',
  templateUrl: './partner-item.component.html',
  styleUrls: ['./partner-item.component.css']
})
export class PartnerItemComponent implements OnInit {

  @Input() PartnerItem;

  constructor() { }

  ngOnInit() {
  }

}
