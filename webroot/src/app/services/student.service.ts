import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { environment } from '../../environments/environment';

@Injectable()
export class StudentService {

  constructor(private http: Http) {
  }
    //Отправить заявку на участие в проекте
    enrollToProject(id: number, token: string, message: string) {
    let headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');
    headers.append('Bearer-token', token);
    return this.http.post(environment.apiUrl + '/v1/project/enroll/' + id + '?message=' + message, JSON.stringify({}), { headers: headers });
  }

  //Отменить заявку на участие в проекте
  unenrollToProject(id: number, token: string) {
    let headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');
    headers.append('Bearer-token', token);
    return this.http.delete(environment.apiUrl + '/v1/project/enroll/' + id, { headers: headers });
  }

  // показать заявки пользователя
  getEnrolledUsersProject(id: number, token: string) {
    let headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');
    headers.append('Bearer-token', token);
    return this.http.get(environment.apiUrl + '/v1/user/id/' + id).map(res => {
      return res.json().enrolled_on;
    });
  }

// получить задачи студента по id проекта
    getProjectStudentTaskItem(id: number) {
    return [
      {
        "number": "645",
        "taskname": "This is my task",
        "data": "20.03.17",
        "author": "Roman",
        "addressee": "Me",
        "tags": ["tag1", "tag2"],
        "rating": "3"

      },
      {
        "number": "645",
        "taskname": "This is my task too",
        "data": "20.03.17",
        "author": "Konstantin",
        "addressee": "Me",
        "tags": ["tag1", "tag2"],
        "rating": "3"
      }

    ];
  }
}
