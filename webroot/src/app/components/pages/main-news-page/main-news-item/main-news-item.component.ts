import { Component, OnInit } from '@angular/core';

export class MainNewsItem {
  constructor(
    public Image: string,
    public Date: string,
    public Titile: string,
    public Text: string) { }
}

/*import {MainNewsItem} from "/webroot/src/app/models/main-news-item";*/

@Component({
  selector: 'app-main-news-item',
  templateUrl: './main-news-item.component.html',
  styleUrls: ['./main-news-item.component.css']
})
export class MainNewsItemComponent implements OnInit {
	newsItems = [
  new MainNewsItem("https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_(2012).svg/2000px-Microsoft_logo_(2012).svg.png", "26 декабря 2016", "Новость НОМЕР", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur malesuada quam sit amet velit tempus gravida. Fusce nisl nunc, vulputate mollis sodales nec, tristique vel mi. Vestibulum ex lectus, finibus eget vulputate ut, molestie quis augue. Pellentesque neque purus, hendrerit eu commodo rutrum, pharetra sit amet nibh. Ut et urna at erat maximus pellentesque at ac elit. In faucibus gravida orci vitae suscipit. Aliquam venenatis lacus a dui bibendum pretium."),
  new MainNewsItem("https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_(2012).svg/2000px-Microsoft_logo_(2012).svg.png", "26 декабря 2016", "Новость НОМЕР", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur malesuada quam sit amet velit tempus gravida. Fusce nisl nunc, vulputate mollis sodales nec, tristique vel mi. Vestibulum ex lectus, finibus eget vulputate ut, molestie quis augue. Pellentesque neque purus, hendrerit eu commodo rutrum, pharetra sit amet nibh. Ut et urna at erat maximus pellentesque at ac elit. In faucibus gravida orci vitae suscipit. Aliquam venenatis lacus a dui bibendum pretium."),
  new MainNewsItem("https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_(2012).svg/2000px-Microsoft_logo_(2012).svg.png", "26 декабря 2016", "Новость НОМЕР", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur malesuada quam sit amet velit tempus gravida. Fusce nisl nunc, vulputate mollis sodales nec, tristique vel mi. Vestibulum ex lectus, finibus eget vulputate ut, molestie quis augue. Pellentesque neque purus, hendrerit eu commodo rutrum, pharetra sit amet nibh. Ut et urna at erat maximus pellentesque at ac elit. In faucibus gravida orci vitae suscipit. Aliquam venenatis lacus a dui bibendum pretium."),
  new MainNewsItem("https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_(2012).svg/2000px-Microsoft_logo_(2012).svg.png", "26 декабря 2016", "Новость НОМЕР", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur malesuada quam sit amet velit tempus gravida. Fusce nisl nunc, vulputate mollis sodales nec, tristique vel mi. Vestibulum ex lectus, finibus eget vulputate ut, molestie quis augue. Pellentesque neque purus, hendrerit eu commodo rutrum, pharetra sit amet nibh. Ut et urna at erat maximus pellentesque at ac elit. In faucibus gravida orci vitae suscipit. Aliquam venenatis lacus a dui bibendum pretium.")
];
myItems = this.newsItems[0];
  constructor() { }

  ngOnInit() {
  }

}
