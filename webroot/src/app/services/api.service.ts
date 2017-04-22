import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import { environment } from '../../environments/environment';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

@Injectable()
export class ApiService {

  constructor(private http: Http) {
  }

  validate(key: string) {
    return this.http.get(environment.apiUrl + '/v1/auth/register/?pass=' + key)
      .catch((error: any) => { return Observable.throw(error) });
  }

  register(user) {
    var headers = new Headers();

    headers.append('Content-Type', 'application/json');

    return this.http.post(environment.apiUrl + '/v1/auth/register', JSON.stringify(user), { headers: headers })
      .map((res: Response) => {
        if (res.json().code)
          localStorage.setItem('validation_code', res.json().code);
        else
          return Observable.throw('no code');
      })
      .catch((error: any) => { return Observable.throw(error) });
  }

  getPublicStudentInfoById(student_id: number) {
    return this.http.get(environment.apiUrl + '/v1/user/id/' + student_id)
  }

  getPublicAuthorInfoById(author_id: number) {
    return {
      "id": 1,
      "firstName": "Anton",
      "lastName": "Antonov",
      "company": "Lanit-Tercom",
      "rating": 45,
      "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin iaculis purus at congue aliquam.  Aliquam justo magna, ornare sed condimentum eget, rutrum rutrum leo. Donec condimentum, odio id mollis iaculis, libero tellus tempus purus, eu pellentesque leo tellus a felis. Donec vestibulum tincidunt ante eget gravida. Suspendisse aliquam sagittis ex a congue. Aliquam vitae erat nisl. Proin commodo turpis in molestie consectetur. Praesent gravida nulla quis elit euismod lobortis. Maecenas non tempus lorem. Curabitur luctus dolor ante, sit amet blandit elit pulvinar vel. Pellentesque egestas dolor ornare est vestibulum scelerisque.",
      "projects": [
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
        }
      ],
      "courses": [
        {
          "Id": 5,
          "Name": "Name here",
          "Description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin iaculis purus at congue aliquam.",
          "Picture": "project.jpg"
        }
      ],
    }
  }

  getMainPageProjects() {
    return this.http.get(environment.apiUrl + '/v1/land/projects/')
      .map((response: Response) => {
        var res = response.json();
        res.forEach(element => {
          element.Logo = environment.apiUrl + element.Logo;
        });
        return res;
      });

  }

  getProjectItems() {
    return this.http.get(environment.apiUrl + '/v1/project/').map((response: Response) => response.json());
  }
  getProjectItemsByUserId(userId: string) {
    return this.http.get(environment.apiUrl + '/v1/project/').map((response: Response) => response.json());
  }

  getProjectById(id: number) {
    return this.http.get(environment.apiUrl + '/v1/project/' + id);
  }


  getMaterialsItems(id: number) {
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

  getProjectNewsItem(id: number) {
    return [
      {
        "Description": "News 1",
        "Links": "#",
        "Main": "Topic 1",
        "Data": "20.07.16 22:10"

      },
      {
        "Description": "News 2",
        "Links": "#",
        "Main": "Topic 2",
        "Data": "19.07.16 16:02"

      }
    ];
  }


  getTaskItem(id: number) {
    return [
      {
        "Task": "Complete this exercise...",
        "Open": "More details",
        "Data": "20.03.17",
        "Number": "1"
      },
      {
        "Task": "Change this sentence...",
        "Open": "More details",
        "Data": "28.03.17",
        "Number": "2"

      }

    ];
  }

  getProjectAllTaskItem(id: number) {
    return [
      {
        "Number": "645",
        "TaskName": "Complete this exercise...",
        "Data": "20.03.17",
        "Author": "Roman",
        "Addressee": "User1",
        "Tags": ["tag1", "tag2"],
        "Rating": "3"

      },
      {
        "Number": "645",
        "TaskName": "Name of Task",
        "Data": "20.03.17",
        "Author": "Konstantin",
        "Addressee": "User2",
        "Tags": ["tag1", "tag2"],
        "Rating": "3"
      },

      {
        "Number": "645",
        "TaskName": "Name of Task",
        "Data": "20.03.17",
        "Author": "Sheldon",
        "Addressee": "User3",
        "Tags": ["some tag", "some tag 2"],
        "Rating": "4"
      }

    ];
  }

  getProjectStudentTaskItem(id: number) {
    return [
      {
        "Number": "645",
        "TaskName": "This is my task",
        "Data": "20.03.17",
        "Author": "Roman",
        "Addressee": "Me",
        "Tags": ["tag1", "tag2"],
        "Rating": "3"

      },
      {
        "Number": "645",
        "TaskName": "This is my task too",
        "Data": "20.03.17",
        "Author": "Konstantin",
        "Addressee": "Me",
        "Tags": ["tag1", "tag2"],
        "Rating": "3"
      }

    ];
  }

  getNewsPage() {
    return this.http.get(environment.apiUrl + '/v1/news/').map((response: Response) => response.json());
  }

  getNewsById(id_: number) {
    return this.http.get(environment.apiUrl + '/v1/news/' + id_)
      .map((response: Response) => response.json());
  }

  private jwt() {
    // create authorization-page header with jwt token
    let currentUser = JSON.parse(localStorage.getItem('current_user'));
    if (currentUser && currentUser.token) {
      let headers = new Headers({ 'Authorization': 'Bearer ' + currentUser.token });
      return new RequestOptions({ headers: headers });
    }
  }
  postProject(project, token: string) {
    let headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json')
    headers.append('Bearer-token', token);
    return this.http.post(environment.apiUrl + '/v1/project', JSON.stringify(project), { headers: headers });
  }
  deleteProject(id: string, token: string) {
    let headers = new Headers();
    headers.append('Accept', 'application/json')
    headers.append('Bearer-token', token);
    return this.http.delete(environment.apiUrl + '/v1/project/' + id, { headers: headers });
  }
  enrollToProject(id: number, token: string) {
    let headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'text/html')
    headers.append('Bearer-token', token);
    return this.http.post(environment.apiUrl + '/v1/project_sign_up' + id, '', { headers: headers });
  }
}
