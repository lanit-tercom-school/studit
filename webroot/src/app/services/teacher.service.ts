import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { environment } from '../../environments/environment';
import { EnrollItem } from 'models/enroll-item';

@Injectable()
export class TeacherService {

  constructor(private http: Http) {
  }

  getEnrollsForTeacher(token: string, id: number) {
    var variable = { id: id };
    var query = `query($id:ID){
  User(Id:$id){
            ProjectOn{
                  Enrolls{
                            Id
                            DateOfCreation
                            Message
                            Project
                            {
                              Id
                              Name
                            }
                            User{
                              Nickname
                              Id 
                            }
                        }
                    }
               }

} &variables=`+ JSON.stringify(variable);
    let headers = new Headers();
    headers.append('Authorization', 'Bearer ' + token);
    return this.http.get(environment.apiUrl + '/graphql?query=' + query, { headers: headers })
      .map((response: Response) => {
        let Enrolls: EnrollItem[] = new Array<EnrollItem>();
        response.json().data.User.ProjectOn.forEach(element => {
          if (element.Enrolls.length != 0) {
            element.Enrolls.forEach(e => {
              Enrolls.push(e)
            });
          }
        });
        return Enrolls
      });
  }

  postUserToProject(user_id: number, project_id: number, token: string) {//Добавить пользователя в проект
    console.log(user_id);
    var variable = { user_id: user_id, project_id: project_id };
    var query = `mutation($user_id:Int! $project_id:Int!)
    {
      PostProjectOn(User:$user_id Project:$project_id){
        Project{
          Id
        }
      }
    } &variables=`+ JSON.stringify(variable);
    let headers = new Headers();
    headers.append('Authorization', 'Bearer ' + token);
    return this.http.get(environment.apiUrl + '/graphql?query=' + query, { headers: headers })
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