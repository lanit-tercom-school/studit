import { Injectable } from '@angular/core';

import { ProjectItem } from './../components/shared/project-list/project-item/project-item';
import { MaterialsItem } from './../components/pages/s-project-page/materials/materials-item/materials-item';

@Injectable()
export class ApiService {

  constructor() { }

  getProjectItems(): ProjectItem[] {
    return [
      {
        "Id": 1,
        "Name": "StudIT",
        "Description": "Разработки сайта летней школы и студенческих проектов Ланит-Терком",
        "Picture": "project.jpg"
      },
      {
        "Id": 2,
        "Name": "TFS Mobile",
        "Description": "Разработка кроссплатфроменного мобильного клиента для Team Foundation Server",
        "Picture": "project.jpg"
      },
      {
        "Id": 3,
        "Name": "CrossCon",
        "Description": "Разработка мобильного клиента расписания конференций",
        "Picture": "project.jpg"
      }
    ];
  }

  getMaterialsItems(id: number): MaterialsItem[] {
    return [
      {
        "Description": "Resource one",
        "Link": "#"
      },
      {
        "Description": "Resource two",
        "Link": "#"
      },
      {
        "Description": "Resource three",
        "Link": "#"
      }
    ];
  }
}
