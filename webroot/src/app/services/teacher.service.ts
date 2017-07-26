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
      })
      .catch((error: any) => {
         console.log('ERROR: TeacherService -> getEnrollsForTeacher()');
         return Observable.throw(error);
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
    .catch((error: any) => {
         console.log('ERROR: TeacherService -> postUserToProject()');
         return Observable.throw(error);
       });
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
      .catch((error: any) => {
         console.log('ERROR: TeacherService -> postProject()');
         return Observable.throw(error);
       });
  }

  deleteProject(id: string, token: string) {
    let headers = new Headers();
    headers.append('Accept', 'application/json')
    headers.append('Bearer-token', token);
    return this.http.delete(environment.apiUrl + '/v1/project/id/' + id, { headers: headers });
  }
}