import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { environment } from '../../environments/environment';

@Injectable()
export class TeacherService {

  constructor(private http: Http) {
  }

   getEnrollsForTeacher(token: string) {
    let headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');
    headers.append('Bearer-token', token);
    return this.http.get(environment.apiUrl + '/v1/project/enroll', { headers: headers }).map((response: Response) => response.json());
  }

  postUserToProject(user_id: number, project_id: number, token: string) {//Добавить пользователя в проект
    let headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');
    headers.append('Bearer-token', token);
    return this.http.post(environment.apiUrl + '/v1/project/user/?user_id=' + user_id + '&project_id=' + project_id, {}, { headers: headers });
  }
    deleteProjectUser(project_id: number, user_id: number, token: string) {//Удалить пользователя проекта
    let headers = new Headers();
    headers.append('Accept', 'application/json');
    headers.append('Bearer-token', token);
    return this.http.delete(environment.apiUrl + '/v1/project/users/?user_id=' + user_id + '&project_id=' + project_id, { headers: headers });
  }

    postProject(project, token: string) {
    var variables = { name: project.Name,
       description: project.Description,
       logo: project.Logo,
       tags: project.Tags
     };
    var query = `mutation ($name: String!
     $description: String!
     $logo: String
     $tags: String)
    {
      PostProject(Name: $name
       Description: $description
       Logo: $logo
       Tags: $tags)
      {
        Id
      }
    } &variables=`+ JSON.stringify(variables);
    let headers = new Headers();
    headers.append('Authorization', 'Bearer ' + token);
      return this.http.get(environment.apiUrl + '/graphql?query=' + query, { headers: headers })
  }

  deleteProject(id: string, token: string) {
    let headers = new Headers();
    headers.append('Accept', 'application/json')
    headers.append('Bearer-token', token);
    return this.http.delete(environment.apiUrl + '/v1/project/id/' + id, { headers: headers });
  }
}