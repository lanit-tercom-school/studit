import { Component, OnInit, Input } from '@angular/core';
import { MaterialsItem } from './materials-item/materials-item';

@Component({
  selector: 'app-materials',
  templateUrl: './materials.component.html',
  styleUrls: ['./materials.component.css']
})
export class MaterialsComponent implements OnInit {

  @Input() public MaterialsList: MaterialsItem[]; 

  constructor() { }

  ngOnInit() {
  }

}
