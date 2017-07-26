import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { TasksItem } from "models/tasks-item";
import { environment } from '../../environments/environment';

@Injectable()
export class TaskService {

    constructor(private http: Http) {
    }

    getTaskItemsFromGitHub(gitHubUrl: string): Observable<TasksItem[]> {
        gitHubUrl = gitHubUrl.slice(18, gitHubUrl.length)
        return this.http.get("https://api.github.com/repos" + gitHubUrl + "/issues").map(res => {
            let taskList: TasksItem[] = new Array<TasksItem>();
            res.json().forEach(element => {
                let taskItem: TasksItem = new TasksItem();
                taskItem.Title = element.title;
                taskItem.Body=element.body;
                taskItem.Url=element.html_url;
                taskItem.NicnameOfUser=element.user.login;
                taskItem.DataOfCreation=element.created_at;
                taskList.push(taskItem);
            });
            return taskList;
        });
    }
}