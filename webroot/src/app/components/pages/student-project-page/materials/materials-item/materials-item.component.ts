import { Component, OnInit, Input } from '@angular/core';
import { MaterialsItem } from './materials-item';

@Component({
  selector: 'app-materials-item',
  templateUrl: './materials-item.component.html',
  styleUrls: ['./materials-item.component.css']
})
export class MaterialsItemComponent implements OnInit {

  @Input() public MaterialsItem: MaterialsItem;

  constructor() { }

  ngOnInit() {
  }

}
