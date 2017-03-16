import { Injectable } from '@angular/core';

import { ProjectItem } from './../components/shared/project-list/project-item/project-item';
import { MaterialsItem } from './../components/pages/s-project-page/materials/materials-item/materials-item';
import {ProjectNewsItem} from "../components/pages/s-project-page/proj-news/proj-news-item/proj-news-item";
import {TaskItem} from "../components/pages/progress/task";
import {TasksItem} from "../components/pages/s-project-page/tasks/tasks-item/tasks-item";

@Injectable()
export class ApiService {

  constructor() {
  }

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

  getProjectById(id: number): ProjectItem
  {
    return this.getProjectItems().find(project => project.Id === id);
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

  getProjectNewsItem(id: number): ProjectNewsItem[] {
    return [
      {
        "Description": "News 1",
        "Links": "#",
        "Main":"Topic 1",
        "Data": "20.07.16 22:10"

      },
      {
        "Description": "News 2",
        "Links": "#",
        "Main":"Topic 2",
        "Data": "19.07.16 16:02"

      }
    ];
  }

  getTaskItem(id: number): TasksItem[] {
    return [
      {
        "Task": "Complete this exercise...",
        "Open": "More details",
        "Data": "20.03.17",
        "Number":"1"
      },
      {
        "Task": "Change this sentence...",
        "Open": "More details",
        "Data": "28.03.17",
        "Number": "2"

      }

    ];
  }

}
