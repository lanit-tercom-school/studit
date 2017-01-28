import { Injectable } from '@angular/core';

import { ProjectItem } from './../components/shared/project-list/project-item/project-item';

@Injectable()
export class ApiService {

  constructor() { }

  getProjectItems(): ProjectItem[] {
    return [
      {
        "Id": 1,
        "Name": "StudIT",
        "Description": "Разработки сайта летней школы и студенческих проектов Ланит-Терком",
        "Picture": "studit.jpg"
      },
      {
        "Id": 2,
        "Name": "TFS Mobile",
        "Description": "Разработка кроссплатфроменного мобильного клиента для Team Foundation Server",
        "Picture": "tfsmobile.jpg"
      },
      {
        "Id": 3,
        "Name": "CrossCon",
        "Description": "Разработка мобильного клиента расписания конференций",
        "Picture": "go.jpg"
      }
    ];
  }  
}
