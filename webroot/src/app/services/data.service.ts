import { Injectable } from '@angular/core';
import { ApiService } from './api.service'

@Injectable()
export class DataService {
  private Projects;
  private userProjects;
  private enrollingProjects;
  private userId: number;
  constructor(private api: ApiService) { }
  load_data() {
    this.userId = JSON.parse(localStorage.getItem('current_user')).id;
    console.log('Data loading...');
    this.api.getProjectItems().subscribe(res => {
      console.log('Projects:');
      this.Projects = res;
      console.log(this.Projects);
      for (let i = 0; i < this.Projects.length; i++) {
        
      }
    })
  }
  
}
