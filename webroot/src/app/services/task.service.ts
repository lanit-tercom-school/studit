import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { environment } from '../../environments/environment';

@Injectable()
export class TaskService {

    constructor(private http: Http) {
    }

    gettaskItem(id: number) {
        return [
            {
                "task": "Complete this exercise...",
                "open": "More details",
                "data": "20.03.17",
                "number": "1"
            },
            {
                "task": "Change this sentence...",
                "open": "More details",
                "data": "28.03.17",
                "number": "2"

            }

        ];
    }
}