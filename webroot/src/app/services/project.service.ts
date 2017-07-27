import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';

import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { ProjectItem } from "models/project-item";
import { MaterialsItem } from "models/materials-item";
import { ProjectNewsItem } from "models/proj-news-item";
import { ProjectTaskItem } from "models/project-task-item";
import { environment } from '../../environments/environment';

import { AlertService } from 'services/alert.service';

@Injectable()
export class ProjectService {

  constructor(private http: Http,
  private alert: AlertService) {
  }

  getMainPageProjects(): Observable<ProjectItem[]> {
    var query = `{ 
      ProjectList(Limit: "3" Offset: "0")
      {
        Description
        DateOfCreation
        Logo
        Tags
        Id
        GitHubUrl
        Name
      }
    }`;
    return this.http.get(environment.apiUrl + '/graphql?query=' + query)
      .map((response: Response) => {
        this.alert.checkGraphQLResponse(response);
        return response.json().data.ProjectList;
      })
      .catch((error: any) => {
          return Observable.throw(error);
      });
  }

  // получить все проекты
  getProjectItems(limit_: number, offset_: number): Observable<ProjectItem[]> {
    var variables = { limit: limit_, offset: offset_ }
    var query = `query($limit:String, $offset: String)
   {
     ProjectList(Offset: $offset Limit: $limit)
    {
      Description
      DateOfCreation
      Logo
      Tags
      Id
      Name
      GitHubUrl
    }
  }&variables=`+ JSON.stringify(variables);
    return this.http.get(environment.apiUrl + '/graphql?query=' + query)
      .map((response: Response) => {
        this.alert.checkGraphQLResponse(response);
        let res = response.json().data.ProjectList;
        return res;
      })
      .catch((error: any) => {
        return Observable.throw(error);
       });
  }

  getProjectById(id_: number): Observable<ProjectItem> {
    var variable = { id: id_ };
    var query = `query($id:ID)
   {
    Project(Id: $id)
    {
      Logo
      Tags
      Status
      Id
      Name
      Description
      DateOfCreation
      GitHubUrl
    }
  }&variables=`+ JSON.stringify(variable);
    return this.http.get(environment.apiUrl + '/graphql?query=' + query)
    .map(response => {
      this.alert.checkGraphQLResponse(response);
      return response.json().data.Project;
    })
    .catch((error: any) => {
         return Observable.throw(error);
    });
  }
  
    getMaterialsItems(id: number) {
    return [
      {
        "description": "Resource one",
        "link": "#"
      },
      {
        "description": "Resource two",
        "link": "#"
      },
      {
        "description": "Resource three",
        "link": "#"
      }
    ];
  }

  getProjectNewsItem(id: number): ProjectNewsItem[] {
    return [
      {
        "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean porttitor dapibus magna",
        "links": "#",
        "main": "Lorem ipsum dolor sit amet",
        "data": "20.07.16 22:10"
      },
      {
        "description": "Nullam cursus ornare quam, vitae tincidunt neque ullamcorper in.",
        "links": "#",
        "main": "Fusce odio lorem",
        "data": "19.07.16 16:02"
      }
    ];
  }

  getProjectAllTaskItem(id: number): ProjectTaskItem[] {
    return [
      {
        "number": "645",
        "taskname": "Complete this exercise...",
        "data": "20.03.17",
        "author": "Roman",
        "addressee": "User1",
        "tags": ["tag1", "tag2"],
        "rating": "3"
      },
      {
        "number": "645",
        "taskname": "Name of task",
        "data": "20.03.17",
        "author": "Konstantin",
        "addressee": "User2",
        "tags": ["tag1", "tag2"],
        "rating": "3"
      },
      {
        "number": "645",
        "taskname": "Name of task",
        "data": "20.03.17",
        "author": "Sheldon",
        "addressee": "User3",
        "tags": ["some tag", "some tag 2"],
        "rating": "4"
      }
    ];
  }

}