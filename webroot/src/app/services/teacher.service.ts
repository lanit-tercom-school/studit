import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { AlertService } from 'services/alert.service';

import { environment } from '../../environments/environment';
import { EnrollItem } from 'models/enroll-item';
import { ProjectItem } from "models/project-item";

@Injectable()
export class TeacherService {

  constructor(private http: Http,
  private alert: AlertService) { }

  getEnrollsForTeacher(token: string, id: number): Observable<EnrollItem[]> {
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
        this.alert.checkGraphQLResponse(response);
        let Enrolls: EnrollItem[] = new Array<EnrollItem>();
        response.json().data.User.ProjectOn.forEach(element => {
          if (element.Enrolls.length != 0) {
            element.Enrolls.forEach(e => {
              Enrolls.push(e)
            });
          }
        });
        return Enrolls;
      })
      .catch((error: any) => {
        return Observable.throw(error);
      });
  }

  postUserToProject(user_id: number, project_id: number, token: string): Observable<ProjectItem> {//Добавить пользователя в проект
    var variable = { user_id: user_id, project_id: project_id };
    var query = `mutation($user_id:Int! $project_id:Int!)
    {
      PostProjectOn(User:$user_id Project:$project_id){
        Project{
          Id
          Name
          Description
          DateOfCreation
        }
      }
    } &variables=`+ JSON.stringify(variable);
    let headers = new Headers();
    headers.append('Authorization', 'Bearer ' + token);
    return this.http.get(environment.apiUrl + '/graphql?query=' + query, { headers: headers })
      .map((response: Response) => {
        this.alert.checkGraphQLResponse(response);
        return response.json().data.Project;
      })
      .catch((error: any) => {
        return Observable.throw(error);
      });
  }

  //TODO: It not work!
  deleteProjectUser(project_id: number, user_id: number, token: string) {//Удалить пользователя проекта
    let headers = new Headers();
    headers.append('Accept', 'application/json');
    headers.append('Bearer-token', token);
    return this.http.delete(environment.apiUrl + '/v1/project/users/?user_id=' + user_id + '&project_id=' + project_id, { headers: headers });
  }


  postProject(project, token: string): Observable<ProjectItem> {
    var variables = {
      name: project.Name,
      description: project.Description,
      logo: project.Logo,
      tags: project.Tags,
      github: project.GitHubUrl,
    };
    var query = `mutation ($name: String!
     $description: String!
     $logo: String
     $github: String!
     $tags: String)
    {
      PostProject(Name: $name
       Description: $description
       GitHubUrl: $github
       Logo: $logo
       Tags: $tags)
      {
        Id
        Name
        Description
        DateOfCreation
        GitHubUrl
      }
    } &variables=`+ JSON.stringify(variables);
    let headers = new Headers();
    headers.append('Authorization', 'Bearer ' + token);
    return this.http.get(environment.apiUrl + '/graphql?query=' + query, { headers: headers })
    .map((response: Response) => {
      this.alert.checkGraphQLResponse(response);
      return response.json().data.PostProject;
    })
    .catch((error: any) => {
        return Observable.throw(error);
      });
  }

  //TODO: It not work!
  deleteProject(id: string, token: string) {
    let headers = new Headers();
    headers.append('Accept', 'application/json')
    headers.append('Bearer-token', token);
    return this.http.delete(environment.apiUrl + '/v1/project/id/' + id, { headers: headers });
  }
}