import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { AlertService } from 'services/alert.service';

import { ProjectItem } from "models/project-item";
import { Message } from "models/message";
import { ProjectTaskItem } from "models/project-task-item";
import { EnrollItem } from "models/enroll-item";

import { environment } from '../../environments/environment';

@Injectable()
export class StudentService {

  constructor(private http: Http,
  private alert: AlertService) {
  }

  //Отправить заявку на участие в проекте
  enrollToProject(user_: number, project_: number, token: string, message_: string): Observable<ProjectItem> {
    var variables = { message: message_, user: user_, project: project_ };
    var query = `mutation ($message: String $user: Int! $project: Int!)
    {
      Enroll(Message: $message User: $user Project: $project)
      {
        Project
        {
          Name
          Description
          DateOfCreation
          Id
        }
      }
    } &variables=`+ JSON.stringify(variables);
    let headers = new Headers();
    headers.append('Authorization', 'Bearer ' + token);
    return this.http.get(environment.apiUrl + '/graphql?query=' + query, { headers: headers })
    .map((response: Response) => { this.alert.checkGraphQLResponse(response); })
    .catch((error: any) => {
         return Observable.throw(error);
       });
  }

  //Отменить заявку на участие в проекте
  unenrollToProject(id: number, token: string): Observable<Message> {
    var variable = { id: id };
    var query = `mutation($id:Int!){
    DeleteProjectEnroll(Id:$id)
    {
      Message
    }
} &variables=`+ JSON.stringify(variable);
    let headers = new Headers();
    headers.append('Authorization', 'Bearer ' + token);
    return this.http.get(environment.apiUrl + '/graphql?query=' + query, { headers: headers })
      .map((response: Response) => { 
        this.alert.checkGraphQLResponse(response);
        return response.json().data.DeleteProjectEnroll })
      .catch((error: any) => {
         return Observable.throw(error);
      });
  }


  // показать заявки пользователя
  getEnrolledUsersProject(id_: number, token: string): Observable<EnrollItem[]> {
    var variable = { id: id_ };
    var query = `query($id:ID){
   User(Id:$id)
   {
       Enrolls
        {
          Id
          DateOfCreation
          Message
          Project
          {
            Id
            Name
          }
        }
  }
} &variables=`+ JSON.stringify(variable);
    let headers = new Headers();
    headers.append('Authorization', 'Bearer ' + token);
    return this.http.get(environment.apiUrl + '/graphql?query=' + query, { headers: headers })
      .map((response: Response) => {
        this.alert.checkGraphQLResponse(response);
        return response.json().data.User.Enrolls
      })
      .catch((error: any) => {
         return Observable.throw(error);
      });
  }

  // получить задачи студента по id проекта
  getProjectStudentTaskItem(id: number): ProjectTaskItem[] {
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
