import { Component, OnInit, DoCheck } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { Router } from '@angular/router';
import { AuthService } from 'services/auth.service';
import { NewsItem } from 'models/news-item';
import { DataService } from 'services/data.service';
import { FileService } from 'services/file.service';
import { TeacherService } from 'services/teacher.service';



@Component({
  selector: 'app-main-news-page',
  templateUrl: './main-news-page.component.html',
  styleUrls: ['./main-news-page.component.css']
})

export class MainNewsPageComponent implements OnInit, DoCheck {
  public CreatedNews = new NewsItem();
  public IsCreated = false;
  public CurrentUser;
  private url: string;
  constructor(
    private data: DataService,
    private fileService: FileService,
    private teacherService: TeacherService,
    private auth: AuthService,
    private router: Router
  ) { }

  ngOnInit() {
    window.scrollTo(0, 0);
    this.CreatedNews.Image = './assets/no_image.png';
    this.CurrentUser = JSON.parse(localStorage.getItem('current_user'));


  }
  ngDoCheck() {
    this.url = this.router.routerState.snapshot.url;
  }

  logout() {
    this.auth.unauthentificatenow();
    this.router.navigateByUrl('/auth');
  }

  load(event) {
  this.fileService.uploadFiles(event.target.files).subscribe(res => {
    this.CreatedNews.Image = res;
  });
  }

  makeNews() {
    // tslint:disable-next-line:max-line-length
    this.teacherService.postNewNews(JSON.parse(localStorage.getItem('current_user')).Token, this.CreatedNews.Title, this.CreatedNews.Description,this.CreatedNews.Image)
      .subscribe(() => {
        console.log('News was added');
        this.IsCreated = true;
        // this.router.navigate(['/home']);
      });
  }

  addImage() {
    let promptValue = prompt('Укажите адрес картинки.', '');
    // tslint:disable-next-line:curly
    if (promptValue !== null && promptValue !== '')
      this.CreatedNews.Image = promptValue;
  }
}
