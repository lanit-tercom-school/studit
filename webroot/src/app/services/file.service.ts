import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';

import { DataService } from "services/data.service";
import { environment } from '../../environments/environment';

@Injectable()
export class FileService {

  constructor(private data: DataService, private http: Http) { }

  uploadFiles(files): Observable<string> {
    let fileList: FileList = files;
    if (fileList.length > 0) {
      let file: File = fileList[0];
      let formData: FormData = new FormData();
      formData.append('uploadFile', file, file.name);
      let headers = new Headers();
      headers.append('Accept', 'application/json');
      headers.append('Authorization', 'Bearer ' + this.data.UserToken);
      var query = `mutation
      {
        PostFile
        {
          Id 
          Path
        }
      }
      `
      return this.http.post(environment.apiUrl + '/graphql?query=' + query, formData, { headers: headers })
        .map(response => response.json().data.PostFile.Path);
    }
  }

}