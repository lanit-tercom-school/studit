import { Component, OnInit } from '@angular/core';
import { ApiService } from './../../../services/api.service';
import { MaterialsItem } from './materials/materials-item/materials-item';

@Component({
  selector: 'app-s-project-page',
  templateUrl: './s-project-page.component.html',
  styleUrls: ['./s-project-page.component.css']
})
export class SProjectPageComponent implements OnInit {

  constructor(private apiService: ApiService) { }

  ngOnInit() {
  }

  getMaterialsItems(): MaterialsItem [] {
    return this.apiService.getMaterialsItems(1);
  }

}
