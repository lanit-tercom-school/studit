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
    let headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json')
    headers.append('Bearer-token', token);
    console.log(environment.apiUrl + '/v1/project/id/', JSON.stringify(project));
    return this.http.post(environment.apiUrl + '/v1/project/id/',
      JSON.stringify(project), { headers: headers });
  }

  deleteProject(id: string, token: string) {
    let headers = new Headers();
    headers.append('Accept', 'application/json')
    headers.append('Bearer-token', token);
    return this.http.delete(environment.apiUrl + '/v1/project/id/' + id, { headers: headers });
  }
}