import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs/Observable';

import { NewsItem } from 'models/news-item';
import { DataService } from 'services/data.service';
import { FileService } from 'services/file.service';
import { TeacherService } from 'services/teacher.service';

@Component({
  selector: 'app-main-news-page',
  templateUrl: './main-news-page.component.html',
  styleUrls: ['./main-news-page.component.css']
})
export class MainNewsPageComponent implements OnInit {
  public CreatedNews = new NewsItem();
  public IsCreated = false;
  constructor(
    private data: DataService,
    private fileService: FileService,
    private teacherService: TeacherService,
  ) { }

  ngOnInit() {
    window.scrollTo(0, 0);
    this.CreatedNews.Image = './assets/no_image.png';
  }
  load(event) {
  this.fileService.uploadFiles(event.target.files).subscribe(res => {
    this.CreatedNews.Image = res;
  });
  }

  makeNews() {
    this.teacherService.postNewNews(JSON.parse(localStorage.getItem('current_user')).Token, this.CreatedNews.Title,this.CreatedNews.Description,this.CreatedNews.Image)
      .subscribe(() => {
        console.log('Project was added');
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
